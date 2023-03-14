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
	"sync"
	"time"
)

type CountryQueryLogic struct {
	tokenChan    chan *queryParam
	statDay      string
	appMap       map[string]*app
	pageSize     int64
	url          string
	workers      int64
	pageRequests int64
	runChan      chan *run
	doneChan     chan struct{}
	pageDoneChan chan struct{}
	pageLock     sync.RWMutex
	actId        int64
}

func NewCountryQueryLogic(day string) *CountryQueryLogic {
	return &CountryQueryLogic{
		tokenChan:    make(chan *queryParam),
		doneChan:     make(chan struct{}),
		pageDoneChan: make(chan struct{}),
		appMap:       map[string]*app{},
		statDay:      day,
		runChan:      make(chan *run),
		workers:      0,
		pageRequests: 0,
		pageSize:     vars.YmlConfig.GetInt64("MarketingApis.PageSize"),
		url:          vars.YmlConfig.GetString("MarketingApis.Reports.CountryQuery"),
		pageLock:     sync.RWMutex{},
	}
}

func (l *CountryQueryLogic) CountryQuery() (err error) {
	if err = l.getApps(); err != nil {
		return err
	}
	tokenList, err := model.NewToken(vars.DBMysql).ReportAccessTokens(vars.AccountTypeMarket, l.actId)
	if err != nil {
		return err
	}
	if len(tokenList) == 0 {
		hlog.NewLog(logrus.WarnLevel, "jobs-country").Log(logrus.Fields{}, "没有可用的 Token")
		return nil
	}

	go l.setTokens(tokenList)

	for token := range l.tokenChan {
		l.pageRequests++
		go l.queryPages(token)
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

func (l *CountryQueryLogic) setTokens(tokenList []*model.Token) {
	for _, tokens := range tokenList {
		if tokens.ExpiredAt.Before(time.Now()) {
			at, err := jobs.Refresh(tokens)
			if err != nil {
				hlog.NewLog(logrus.ErrorLevel, "jobs-country-refreshToken").Log(logrus.Fields{
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

func (l *CountryQueryLogic) query(param *queryParam, page int64) {
	defer func() {
		l.doneChan <- struct{}{}
	}()
	data := CountryRequest{
		TimeGranularity: StateTimeDaily,
		StartDate:       l.statDay,
		EndDate:         l.statDay,
		Page:            page,
		PageSize:        l.pageSize,
		IsAbroad:        true,
		OrderType:       OrderAsc,
		Filtering: Filtering{
			OtherFilterType: FilterTypeAdgroup,
		},
	}
	var response CountryResponse
	c, err := curl.New(l.url).Debug(false).Post().JsonData(data)
	if err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-country-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口构建失败："+err.Error())
		go l.setFailed(param, page)
		return
	}
	if err = c.Request(&response, curl.Authorization(param.accessToken), curl.JsonHeader()); err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-country-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口请求失败："+err.Error())
		go l.setFailed(param, page)
		return
	}
	if response.Code != "0" {
		hlog.NewLog(logrus.WarnLevel, "jobs-country-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口响应失败："+response.Message)
		go l.setFailed(param, page)
		return
	}
	if err = l.saveCountryData(param.accountId, response.Data.List); err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-country-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "数据存储失败："+err.Error())
		go l.setFailed(param, page)
		return
	}

	return
}

func (l *CountryQueryLogic) setFailed(param *queryParam, page int64) {
	for {
		if l.workers < maxGoroutine {
			if param.failed < failRetryTimes {
				l.runChan <- &run{
					queryParam: &queryParam{failed: param.failed + 1, accountId: param.accountId, accessToken: param.accessToken},
					page:       page,
				}
			} else {
				// log
				hlog.NewLog(logrus.ErrorLevel, "jobs-country-request").Log(logrus.Fields{
					"account_id": param.accountId,
					"failed":     param.failed,
					"stat_day":   l.statDay,
				}, "接口调用错误次数超出")
			}
			break
		}
	}
	return
}

func (l *CountryQueryLogic) setPageFailed(param *queryParam) {
	if param.failed < failRetryTimes {
		// 写锁?
		l.pageLock.Lock()
		l.pageRequests++
		l.pageLock.Unlock()
		go l.queryPages(&queryParam{failed: param.failed + 1, accountId: param.accountId, accessToken: param.accessToken})
	} else {
		// log
		hlog.NewLog(logrus.ErrorLevel, "jobs-country-page-request").Log(logrus.Fields{
			"account_id": param.accountId,
			"failed":     param.failed,
			"stat_day":   l.statDay,
		}, "接口调用错误次数超出")
	}
	return
}

func (l *CountryQueryLogic) queryPages(param *queryParam) {
	// 防止某个账户的任务太慢，第一个还没有执行完，其他所有账户任务运行结束了
	defer func() {
		l.pageDoneChan <- struct{}{}
	}()
	data := CountryRequest{
		TimeGranularity: StateTimeDaily,
		StartDate:       l.statDay,
		EndDate:         l.statDay,
		Page:            1,
		PageSize:        1,
		IsAbroad:        true,
		OrderType:       OrderAsc,
		Filtering: Filtering{
			OtherFilterType: FilterTypeAdgroup,
		},
	}
	var response CountryPagesResponse
	c, err := curl.New(l.url).Debug(false).Post().JsonData(data)
	if err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-country-page-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口构建失败："+err.Error())
		go l.setPageFailed(param)
		return
	}
	if err = c.Request(&response, curl.Authorization(param.accessToken), curl.JsonHeader()); err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-country-page-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口构建失败："+err.Error())
		go l.setPageFailed(param)
		return
	}
	if response.Code != "0" {
		hlog.NewLog(logrus.WarnLevel, "jobs-country-page-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口响应失败："+response.Message)
		go l.setPageFailed(param)
		return
	}

	totalPage := utils.CeilPages(response.Data.PageInfo.TotalNum, l.pageSize)
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

func (l *CountryQueryLogic) getApps() error {
	list, err := model.NewApp(vars.DBMysql).JobGetApps()
	if err != nil {
		return err
	}
	for _, _app := range list {
		l.appMap[_app.PkgName] = &app{
			appId:   _app.AppId,
			appName: _app.AppName,
		}
	}

	return nil
}

func (l *CountryQueryLogic) matchApp(pkgName string) *app {
	if rs, ok := l.appMap[pkgName]; ok {
		return rs
	}
	return &app{appId: "", appName: ""}
}

func (l *CountryQueryLogic) saveCountryData(accountId int64, data []*Country) (err error) {
	if len(data) > 0 {
		m := make([]*model.ReportMarketSource, len(data))
		now := time.Now()
		for i, d := range data {
			day, err := time.Parse("2006010215", d.StatDatetime)
			if err != nil {
				day, _ = time.Parse("2006-01-02 15", l.statDay+" 00")
			}
			_app := l.matchApp(d.PackageName)
			m[i] = &model.ReportMarketSource{
				StatDay:       day,
				StatHour:      uint8(day.Hour()),
				Country:       d.Country,
				AccountId:     accountId,
				AppId:         _app.appId,
				AppName:       _app.appName,
				PkgName:       d.PackageName,
				CampaignId:    d.CampaignId,
				CampaignName:  d.CampaignName,
				AdgroupId:     d.AdgroupId,
				AdgroupName:   d.AdgroupName,
				CreativeId:    d.CreativeId,
				CreativeName:  d.CreativeName,
				Cost:          utils.StringToFloat(d.Cost),
				ShowCount:     d.ShowCount,
				ClickCount:    d.ClickCount,
				DownloadCount: d.DownloadCount,
				InstallCount:  d.InstallCount,
				ActivateCount: d.ActiveCountNormalized,
				RetainCount:   d.RetainCountNormalized,
				Cpm:           utils.StringToFloat(d.Cpm),
				Cpc:           utils.StringToFloat(d.Cpc),
				Cpd:           utils.StringToFloat(d.Cpd),
				Cpi:           utils.StringToFloat(d.Cpi),
				Cpa:           utils.StringToFloat(d.Cpa),
				RetainCost:    utils.StringToFloat(d.RetainCostNormalized),
				Timestamp:     model.Timestamp{CreatedAt: now, UpdatedAt: now},
				//ClickThroughRate:     getRate(float64(d.ClickCount), d.ShowCount, 6),
				//ClickDownloadRate:    getRate(float64(d.DownloadCount), d.ClickCount, 6),
				//DownloadActivateRate: getRate(float64(d.ActiveCountNormalized), d.DownloadCount, 6),
			}
		}

		return model.NewRMS(vars.DBMysql).BatchInsert(m)
	}
	return nil
}
