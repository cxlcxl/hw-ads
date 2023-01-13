package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	serviceaccount "bs.mobgi.cc/app/service/account"
	servicemarketing "bs.mobgi.cc/app/service/marketing"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/curl"
	"github.com/gin-gonic/gin"
)

type Targeting struct{}

func (h *Targeting) List(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VTargetingList)
	targets, err := model.NewTargeting(vars.DBMysql).GetTargets(params.AccountId)
	if err != nil {
		response.Fail(ctx, "没有本地定向包，需要同步："+err.Error())
		return
	}
	rs, err := servicemarketing.UnFormatTargetData(targets)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, rs)
}

func (h *Targeting) Location(ctx *gin.Context) {
	continents, err := model.NewContinent(vars.DBMysql).Continents()
	if err != nil {
		response.Fail(ctx, "大洲信息未填充")
		return
	}
	countries, err := model.NewOverseasRegion(vars.DBMysql).GetCountries()
	if err != nil {
		response.Fail(ctx, "国家信息为空")
		return
	}
	response.Success(ctx, gin.H{"continents": continents, "countries": countries})
}

func (h *Targeting) Create(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VTargetingCreate)

	total, err := model.NewTargeting(vars.DBMysql).CheckExistsByTargetName(params.TargetingName)
	if err != nil {
		response.Fail(ctx, "查询定向包失败："+err.Error())
		return
	}
	if total > 0 {
		response.Fail(ctx, "定向包名已存在")
		return
	}
	campaign, err := model.NewCampaign(vars.DBMysql).FindByCampaignId(params.CampaignId)
	if err != nil {
		response.Fail(ctx, "查询计划失败："+err.Error())
		return
	}
	data, err := servicemarketing.FormatAdsData(params, campaign.AdvertiserId)
	if err != nil {
		response.Fail(ctx, "格式化数据错误："+err.Error())
		return
	}
	c, err := curl.New(vars.YmlConfig.GetString("MarketingApis.Targeting.Create")).Post().JsonData(data)
	if err != nil {
		response.Fail(ctx, "请求构建错误："+err.Error())
		return
	}
	t, err := serviceaccount.GetToken(campaign.AccountId)
	if err != nil {
		response.Fail(ctx, "Token 查询错误："+err.Error())
		return
	}
	var rs servicemarketing.TargetingCreateResp
	if err = c.Request(&rs, curl.Authorization(t.AccessToken)); err != nil {
		response.Fail(ctx, "华为API调用失败："+err.Error())
		return
	}
	if rs.Code != "200" {
		response.Fail(ctx, "华为API响应失败："+err.Error())
		return
	}

	targeting := servicemarketing.FormatRDBData(params, campaign.AdvertiserId, rs.Data.TargetingId, campaign.AccountId)
	if err = model.NewTargeting(vars.DBMysql).TargetingCreate(targeting); err != nil {
		response.Fail(ctx, "本地写入失败："+err.Error())
		return
	}
	response.Success(ctx, rs.Data.TargetingId)
}
