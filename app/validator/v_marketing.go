package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VResourcePricing(ctx *gin.Context) {
	response.Success(ctx, vars.Pricing)
}

func (v BsValidator) VTrackingList(ctx *gin.Context) {
	var params v_data.VTrackingList
	bindData(ctx, &params, (&handlers.Marketing{}).TrackingList)
}

func (v BsValidator) VTrackingRefresh(ctx *gin.Context) {
	var params v_data.VTrackingRefresh
	bindData(ctx, &params, (&handlers.Marketing{}).TrackingRefresh)
}

func (v BsValidator) VDictQuery(ctx *gin.Context) {
	var params v_data.VDictQuery
	bindData(ctx, &params, (&handlers.Marketing{}).DictQuery)
}
