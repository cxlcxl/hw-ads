package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs"
	"bs.mobgi.cc/library/curl"
	"bs.mobgi.cc/library/hlog"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type AppQueryLogic struct {
	tokenChan    chan *queryParam
	pageSize     int64
	url          string
	workers      int64
	pageRequests int64
	runChan      chan *run
	doneChan     chan struct{}
	pageDoneChan chan struct{}
}

func NewAppQueryLogic() *AppQueryLogic {
	return &AppQueryLogic{
		tokenChan:    make(chan *queryParam),
		doneChan:     make(chan struct{}),
		pageDoneChan: make(chan struct{}),
		workers:      0,
		runChan:      make(chan *run),
		pageRequests: 0,
		pageSize:     maxPageSize,
		url:          vars.YmlConfig.GetString("MarketingApis.App.Query"),
	}
}

func (l *AppQueryLogic) AppQuery() (err error) {
	tokenList, err := model.NewToken(vars.DBMysql).ReportAccessTokens(vars.AccountTypeMarket)
	if err != nil {
		return err
	}
	if len(tokenList) == 0 {
		fmt.Println("没有可用的 Token")
		return nil
	}

	go l.setTokens(tokenList)

	for token := range l.tokenChan {
		l.pageRequests++
		go l.queryPages(token)
	}
	// 正式调度前，清楚历史关联
	if err = model.NewAppAct(vars.DBMysql).FlushInfo(); err != nil {
		return err
	}
	for {
		select {
		case runParam, ok := <-l.runChan:
			if ok {
				l.workers++
				go l.query(runParam.queryParam, runParam.page)
			}
		case <-l.doneChan:
			l.workers--
		case <-l.pageDoneChan:
			l.pageRequests--
		default:
		}
		if l.workers == 0 && l.pageRequests == 0 {
			break
		}
	}
	return
}

func (l *AppQueryLogic) setTokens(tokenList []*model.Token) {
	for _, tokens := range tokenList {
		if tokens.ExpiredAt.Before(time.Now()) {
			at, err := jobs.Refresh(tokens)
			if err != nil {
				hlog.NewLog(logrus.ErrorLevel, "jobs-app-refreshToken").Log(logrus.Fields{
					"account_id": tokens.AccountId,
				}, err)
				continue
			}
			l.tokenChan <- &queryParam{
				accountId:   tokens.AccountId,
				accessToken: fmt.Sprintf("Bearer %s", at),
			}
		} else {
			l.tokenChan <- &queryParam{
				accountId:   tokens.AccountId,
				accessToken: fmt.Sprintf("Bearer %s", tokens.AccessToken),
			}
		}
	}
	close(l.tokenChan)
}

func (l *AppQueryLogic) queryPages(param *queryParam) {
	// 防止某个账户的任务太慢，第一个还没有执行完，其他所有账户任务运行结束了
	defer func() {
		l.pageDoneChan <- struct{}{}
	}()
	data := AppRequest{
		Page:     1,
		PageSize: 10, // api 要求最小 10
		Filtering: ReqFiltering{
			ProductType: vars.ProductTypeAndroidApp,
			AgAppType:   agApp,
		},
	}
	var response AppResponse
	c, err := curl.New(l.url).Get().JsonData(data)
	if err != nil {
		hlog.NewLog(logrus.ErrorLevel, "jobs-app-page-request").Log(logrus.Fields{
			"account_id": param.accountId,
		}, "API 接口构建失败："+err.Error())
		return
	}
	if err = c.Request(&response, curl.Authorization(param.accessToken), curl.JsonHeader()); err != nil {
		hlog.NewLog(logrus.ErrorLevel, "jobs-app-page-request").Log(logrus.Fields{
			"account_id": param.accountId,
		}, "API 接口请求失败："+err.Error())
		return
	}
	if response.Code != "200" {
		hlog.NewLog(logrus.ErrorLevel, "jobs-app-page-request").Log(logrus.Fields{
			"account_id": param.accountId,
		}, "API 接口响应失败："+response.Message)
		return
	}

	totalPage := utils.CeilPages(response.Data.Total, l.pageSize)
	var i int64 = 1
	for i <= totalPage {
		if l.workers < maxGoroutine {
			l.runChan <- &run{
				queryParam: param,
				page:       i,
			}
			i++
		}
	}
	return
}

func (l *AppQueryLogic) query(param *queryParam, page int64) {
	defer func() {
		l.doneChan <- struct{}{}
	}()
	data := AppRequest{
		Page:     page,
		PageSize: l.pageSize,
		Filtering: ReqFiltering{
			ProductType: vars.ProductTypeAndroidApp,
			AgAppType:   agApp,
		},
	}
	var response AppResponse
	c, err := curl.New(l.url).Debug(false).Get().JsonData(data)
	if err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-app-request").Log(logrus.Fields{
			"account_id": param.accountId,
		}, "API 接口构建失败："+err.Error())

		l.setFailed(param, page)
		return
	}
	if err = c.Request(&response, curl.Authorization(param.accessToken), curl.JsonHeader()); err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-app-request").Log(logrus.Fields{
			"account_id": param.accountId,
		}, "API 接口请求失败："+err.Error())
		l.setFailed(param, page)
		return
	}
	if response.Code != "200" {
		hlog.NewLog(logrus.WarnLevel, "jobs-app-request").Log(logrus.Fields{
			"account_id": param.accountId,
		}, "API 接口响应失败："+response.Message)
		l.setFailed(param, page)
		return
	}
	if err = l.saveAppData(param.accountId, response.Data.Data); err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-app-request").Log(logrus.Fields{
			"account_id": param.accountId,
		}, "数据存储失败："+err.Error())
		l.setFailed(param, page)
		return
	}

	return
}

func (l *AppQueryLogic) setFailed(param *queryParam, page int64) {
	for {
		if l.workers < maxGoroutine {
			if param.failed < failRetryTimes {
				l.runChan <- &run{
					queryParam: &queryParam{
						accountId:   param.accountId,
						accessToken: param.accessToken,
						failed:      param.failed + 1,
					},
					page: page,
				}
			} else {
				// log
				hlog.NewLog(logrus.ErrorLevel, "jobs-app-request").Log(logrus.Fields{
					"account_id": param.accountId,
					"failed":     param.failed,
				}, "接口调用错误次数超出")
			}

			break
		}
	}
	return
}

func (l *AppQueryLogic) saveAppData(accountId int64, data []*ProductInfo) (err error) {
	if len(data) > 0 {
		m := make([]*model.App, len(data))
		appAccounts := make([]*model.AppAccount, 0)
		now := time.Now()
		for i, d := range data {
			app := d.ProductInfo.App
			m[i] = &model.App{
				AppId:     app.AppId,
				AppName:   app.ProductName,
				PkgName:   app.PackageName,
				Channel:   vars.AppChannelGallery,
				IconUrl:   app.IconUrl,
				ProductId: d.ProductId,
				Timestamp: model.Timestamp{CreatedAt: now, UpdatedAt: now},
			}

			appAccounts = append(appAccounts, &model.AppAccount{
				AccountId:   accountId,
				AppId:       app.AppId,
				AccountType: vars.AccountTypeMarket,
			})
		}

		return model.NewApp(vars.DBMysql).BatchInsert(m, appAccounts)
	}
	return nil
}
