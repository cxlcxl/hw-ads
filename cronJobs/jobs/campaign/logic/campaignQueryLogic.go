package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/job_data/statements"
	"bs.mobgi.cc/cronJobs/jobs"
	"bs.mobgi.cc/library/curl"
	"fmt"
	"time"
)

type CampaignQueryLogic struct {
	tokenChan chan *jobs.QueryParam
	statDay   string
	pageSize  int64
	campaigns []*model.Campaign
	worker    int
}

func NewCampaignQueryLogic(day string) *CampaignQueryLogic {
	return &CampaignQueryLogic{
		tokenChan: make(chan *jobs.QueryParam),
		campaigns: make([]*model.Campaign, 0),
		statDay:   day,
		pageSize:  50,
		worker:    0,
	}
}

func (l *CampaignQueryLogic) CampaignQuery() (err error) {
	if err != nil {
		return err
	}
	go jobs.GetTokens(l.tokenChan)

	for token := range l.tokenChan {
		l.query(token, 1)
	}

	if len(l.campaigns) > 0 {
		err = model.NewCampaign(vars.DBMysql).BatchInsert(l.campaigns)
	}
	return
}

func (l *CampaignQueryLogic) query(param *jobs.QueryParam, page int64) {
	data := statements.CampaignRequest{
		AdvertiserId: param.AdvertiserId,
		Page:         page,
		PageSize:     l.pageSize,
		Filtering: statements.CampaignFiltering{
			UpdatedBeginTime: l.statDay + " 00:00:00",
			UpdatedEndTime:   l.statDay + " 23:59:59",
		},
	}
	var response statements.CampaignResponse
	c, err := curl.New(vars.YmlConfig.GetString("Promotion.Campaign")).Get().JsonData(data)
	if err != nil {
		fmt.Println("参数生成失败", err)
		return
	}
	if err = c.Request(&response, curl.Authorization(param.AccessToken)); err != nil {
		fmt.Println("接口请求失败", err)
		return
	}
	if response.Code != "200" {
		fmt.Println("接口请求失败", response.Code, response.Message)
		return
	}
	if len(response.Data.Data) == 0 {
		return
	}
	now := time.Now()
	for _, datum := range response.Data.Data {
		l.campaigns = append(l.campaigns, &model.Campaign{
			CampaignId:                datum.CampaignId,
			AppId:                     "",
			CampaignName:              datum.CampaignName,
			AccountId:                 param.AccountId,
			AdvertiserId:              param.AdvertiserId,
			OptStatus:                 datum.CampaignStatus,
			CampaignDailyBudgetStatus: datum.CampaignDailyBudgetStatus,
			ProductType:               datum.ProductType,
			ShowStatus:                datum.ShowStatus,
			UserBalanceStatus:         datum.UserBalanceStatus,
			FlowResource:              datum.FlowResource,
			SyncFlowResource:          datum.SyncFlowResourceSearchAd,
			CampaignType:              datum.CampaignType,
			TodayDailyBudget:          datum.TodayDailyBudget,
			TomorrowDailyBudget:       datum.TomorrowDailyBudget,
			MarketingGoal:             datum.MarketingGoal,
			CreatedAt:                 now,
			UpdatedAt:                 now,
		})
	}

	if page > 1 {
		return
	}

	sumPages := utils.CeilPages(response.Data.Total, l.pageSize)
	if sumPages > 1 {
		var i int64 = 2
		for ; i <= sumPages; i++ {
			l.query(param, i)

			time.Sleep(time.Millisecond * 500)
		}
	}

	return
}
