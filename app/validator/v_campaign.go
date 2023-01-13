package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VCampaignList(ctx *gin.Context) {
	var params v_data.VCampaignList
	bindData(ctx, &params, emptyValidator, (&handlers.Campaign{}).List)
}

func (v BsValidator) VCampaignInfo(ctx *gin.Context) {
	bindRouteData(ctx, "campaign_id", (&handlers.Campaign{}).CampaignInfo)
}
