package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs"
	"bs.mobgi.cc/library/curl"
	"fmt"
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
	}
}

func (l *AdsQueryLogic) AdsQuery() (err error) {
	tokenList, err := model.NewToken(vars.DBMysql).ReportAccessTokens(vars.AccountTypeAds)
	if err != nil {
		return err
	}
	if len(tokenList) == 0 {
		fmt.Println("没有可用的 Token")
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
				fmt.Println("Token 刷新失败，账户 ID：", tokens.AccountId, err)
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
		fmt.Println("API 接口构建失败：" + err.Error())
		return
	}
	if err = c.Request(&response, curl.Authorization(param.accessToken), curl.JsonHeader()); err != nil {
		fmt.Println("API 接口请求失败：" + err.Error())
		return
	}
	if response.Code != "0" {
		fmt.Println("API 接口响应失败：" + response.Message)
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
		fmt.Println("API 接口构建失败：" + err.Error())
		l.setFailed(param, page)
		return
	}
	if err = c.Request(&response, curl.Authorization(param.accessToken), curl.JsonHeader()); err != nil {
		fmt.Println("API 接口请求失败：" + err.Error())
		l.setFailed(param, page)
		return
	}
	if response.Code != "0" {
		fmt.Println("API 接口响应失败：" + response.Message)
		l.setFailed(param, page)
		return
	}
	if err = l.saveAdsData(param.accountId, response.Data.List); err != nil {
		fmt.Println("数据存储失败：" + err.Error())
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
				// TODO log
			}

			break
		}
	}
	return
}

func (l *AdsQueryLogic) saveAdsData(accountId int64, data []*AdsList) (err error) {
	if len(data) > 0 {
		m := make([]*model.ReportAdsSource, len(data))
		now := time.Now()
		for i, d := range data {
			day, err := time.Parse("2006-01-02 15", d.StatDatetime)
			if err != nil {
				day, _ = time.Parse("2006-01-02 15", l.statDay+" 00")
			}
			eCpm := 0.0
			if d.ShowCount > 0 {
				eCpm = getRate(d.Earnings*1000, d.ShowCount)
			}
			m[i] = &model.ReportAdsSource{
				StatDay:             day,
				StatHour:            uint8(day.Hour()),
				Country:             d.Country,
				AccountId:           accountId,
				AppId:               "C" + d.AppId,
				AdType:              d.AdType,
				PlacementId:         d.PlacementId,
				AdRequests:          d.ReachedAdRequests,
				MatchedAdRequests:   d.MatchedReachedAdRequests,
				ShowCount:           d.ShowCount,
				ClickCount:          d.ClickCount,
				AdRequestsMatchRate: d.AdRequestsMatchRate,
				AdRequestsShowRate:  d.AdRequestsShowRate,
				ClickThroughRate:    d.ClickThroughRate,
				Earnings:            d.Earnings,
				ECpm:                eCpm,
				Timestamp:           model.Timestamp{CreatedAt: now, UpdatedAt: now},
			}
		}

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

func getRate(a float64, b int64) float64 {
	if b == 0 {
		return 0
	}
	return utils.Round(a/float64(b), 6)
}
