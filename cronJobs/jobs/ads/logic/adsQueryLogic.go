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

type AdsQueryLogic struct {
	tokenChan    chan *queryParam
	statDay      string
	pageSize     int64
	url          string
	workers      int64
	pageRequests int64
	runChan      chan *run
	doneChan     chan struct{}
	pageDoneChan chan struct{}
	clientInfo   map[int64]*client
	mu           sync.Mutex
}

func NewAdsQueryLogic(day string) *AdsQueryLogic {
	return &AdsQueryLogic{
		tokenChan:    make(chan *queryParam),
		doneChan:     make(chan struct{}),
		pageDoneChan: make(chan struct{}),
		statDay:      day,
		workers:      0,
		runChan:      make(chan *run),
		pageRequests: 0,
		pageSize:     vars.YmlConfig.GetInt64("MarketingApis.PageSize"),
		url:          vars.YmlConfig.GetString("MarketingApis.Reports.AdsQuery"),
		mu:           sync.Mutex{},
	}
}

func (l *AdsQueryLogic) AdsQuery() (err error) {
	tokenList, err := model.NewToken(vars.DBMysql).ReportAccessTokens(vars.AccountTypeAds)
	if err != nil {
		return err
	}
	if len(tokenList) == 0 {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-request").Log(logrus.Fields{}, "没有可用的 Token")
		return nil
	}
	if err = l.setClientInfo(); err != nil {
		return err
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

func (l *AdsQueryLogic) setTokens(tokenList []*model.Token) {
	for _, tokens := range tokenList {
		if tokens.ExpiredAt.Before(time.Now()) {
			clientId, secret := l.getClientInfo(tokens.AccountId)
			at, err := jobs.RefreshV3(tokens, clientId, secret)
			if err != nil {
				hlog.NewLog(logrus.ErrorLevel, "jobs-ads-refreshToken").Log(logrus.Fields{
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

func (l *AdsQueryLogic) queryPages(param *queryParam) {
	// 防止某个账户的任务太慢，第一个还没有执行完，其他所有账户任务运行结束了
	defer func() {
		l.pageDoneChan <- struct{}{}
	}()
	data := AdsRequest{
		TimeGranularity: StateTimeDaily,
		StartDate:       l.statDay,
		EndDate:         l.statDay,
		Page:            1,
		PageSize:        1,
		OrderType:       OrderAsc,
		GroupBy:         RequestGroupBy,
		Filtering: Filtering{
			Currency: CurrencyUsd,
		},
	}
	var response AdsPagesResponse
	c, err := curl.New(l.url).Debug(false).Post().JsonData(data)
	if err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-page-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口构建失败："+err.Error())
		return
	}
	if err = c.Request(&response, curl.Authorization(param.accessToken), curl.JsonHeader()); err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-page-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口请求失败："+err.Error())
		return
	}
	if response.Code != "0" {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-page-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口响应失败："+response.Message)
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

func (l *AdsQueryLogic) query(param *queryParam, page int64) {
	defer func() {
		l.doneChan <- struct{}{}
	}()
	data := AdsRequest{
		TimeGranularity: StateTimeDaily,
		StartDate:       l.statDay,
		EndDate:         l.statDay,
		Page:            page,
		PageSize:        l.pageSize,
		GroupBy:         RequestGroupBy,
		OrderType:       OrderAsc,
		Filtering: Filtering{
			Currency: CurrencyUsd,
		},
	}
	var response AdsResponse
	c, err := curl.New(l.url).Debug(false).Post().JsonData(data)
	if err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口构建失败："+err.Error())
		l.setFailed(param, page)
		return
	}
	if err = c.Request(&response, curl.Authorization(param.accessToken), curl.JsonHeader()); err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口请求失败："+err.Error())
		l.setFailed(param, page)
		return
	}
	if response.Code != "0" {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "API 接口响应失败："+response.Message)
		l.setFailed(param, page)
		return
	}
	if err = l.saveAdsData(param.accountId, response.Data.List); err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-request").Log(logrus.Fields{
			"stat_day": l.statDay,
		}, "数据存储失败："+err.Error())
		l.setFailed(param, page)
		return
	}

	return
}

func (l *AdsQueryLogic) setFailed(param *queryParam, page int64) {
	for {
		if l.workers < maxGoroutine {
			if param.failed < failRetryTimes {
				l.runChan <- &run{
					queryParam: &queryParam{
						accountId:    param.accountId,
						clientId:     param.clientId,
						clientSecret: param.clientSecret,
						accessToken:  param.accessToken,
						failed:       param.failed + 1,
					},
					page: page,
				}
			} else {
				// log
				hlog.NewLog(logrus.ErrorLevel, "jobs-ads-request").Log(logrus.Fields{
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

func (l *AdsQueryLogic) saveAdsData(accountId int64, data []*AdsList) (err error) {
	if len(data) > 0 {
		l.mu.Lock() // 防止 Mysql 事务抱死锁错误
		defer l.mu.Unlock()
		m := make([]*model.ReportAdsSource, len(data))
		now := time.Now()
		for i, d := range data {
			day, err := time.Parse("2006-01-02 15", d.StatDatetime)
			if err != nil {
				day, _ = time.Parse("2006-01-02 15", l.statDay+" 00")
			}
			m[i] = &model.ReportAdsSource{
				StatDay:           day,
				StatHour:          uint8(day.Hour()),
				Country:           d.Country,
				AccountId:         accountId,
				AppId:             "C" + d.AppId,
				AdType:            d.AdType,
				PlacementId:       d.PlacementId,
				AdRequests:        d.ReachedAdRequests,
				MatchedAdRequests: d.MatchedReachedAdRequests,
				ShowCount:         d.ShowCount,
				ClickCount:        d.ClickCount,
				Earnings:          d.Earnings,
				Timestamp:         model.Timestamp{CreatedAt: now, UpdatedAt: now},
			}
		}
		//if producer, err := queue.NewKafkaProducer("report_ads_source", "ads"); err == nil {
		//	_ = producer.SendMessages(func(s string, encoder sarama.StringEncoder) (msg []*sarama.ProducerMessage) {
		//		for _, source := range m {
		//			if marshal, err := json.Marshal(source); err == nil {
		//				msg = append(msg, &sarama.ProducerMessage{
		//					Topic: s,
		//					Key:   encoder,
		//					Value: sarama.ByteEncoder(marshal),
		//				})
		//			}
		//		}
		//		return
		//	})
		//}
		return model.NewRAS(vars.DBMysql).BatchInsert(m)
	}
	return nil
}

func (l *AdsQueryLogic) setClientInfo() error {
	info, err := model.NewAct(vars.DBMysql).ClientInfo(vars.AccountTypeAds)
	if err != nil {
		return err
	}
	l.clientInfo = make(map[int64]*client)
	for _, clientInfo := range info {
		l.clientInfo[clientInfo.Id] = &client{
			clientId: clientInfo.ClientId,
			secret:   clientInfo.Secret,
		}
	}
	return nil
}

func (l *AdsQueryLogic) getClientInfo(accountId int64) (string, string) {
	if c, ok := l.clientInfo[accountId]; ok {
		return c.clientId, c.secret
	} else {
		return "", ""
	}
}
