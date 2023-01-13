package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
)

type Campaign struct{}

func (h *Campaign) Resources(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"opt_status":        vars.OptStatus,
		"product_type":      vars.ProductType,
		"campaign_type":     vars.CampaignType,
		"sync_flow":         vars.SyncFlow,
		"daily_budget_opts": vars.DailyBudgetOpts,
		"show_status":       vars.CampaignShowStatus,
		"campaign_daily":    vars.CampaignDaily,
		"flow_resource":     vars.FlowResource,
		"balance_status":    vars.BalanceStatus,
	})
}

func (h *Campaign) List(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VCampaignList)
	offset := utils.GetPages(params.Page, params.PageSize)
	list, total, err := model.NewCampaign(vars.DBMysql).CampaignList(
		params.AppId,
		params.CampaignId,
		params.CampaignName,
		params.CampaignType,
		offset,
		params.PageSize,
	)
	if err != nil {
		response.Fail(ctx, "查询失败："+err.Error())
		return
	}
	response.Success(ctx, gin.H{"total": total, "campaigns": list})
}

func (h *Campaign) CampaignInfo(ctx *gin.Context, v string) {
	campaign, err := model.NewCampaign(vars.DBMysql).FindByCampaignId(v)
	if err != nil {
		response.Fail(ctx, "查询失败："+err.Error())
		return
	}
	app, err := model.NewApp(vars.DBMysql).FindAppByAppId(campaign.AppId)
	if err != nil {
		response.Fail(ctx, "计划关联应用查询失败："+err.Error())
		return
	}
	campaign.App = model.BelongApp{
		ProductId: app.ProductId,
		AppName:   app.AppName,
		IconUrl:   app.IconUrl,
	}
	response.Success(ctx, campaign)
}
