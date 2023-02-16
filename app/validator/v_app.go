package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VAppList(ctx *gin.Context) {
	var params v_data.VAppList
	bindData(ctx, &params, (&handlers.App{}).AppList, fillUser)
}

func (v BsValidator) VAppCampaignList(ctx *gin.Context) {
	var params v_data.VAppCampaignList
	bindData(ctx, &params, (&handlers.App{}).AppCampaignList)
}

func (v BsValidator) VAppPull(ctx *gin.Context) {
	var params v_data.VAppPull
	bindData(ctx, &params, (&handlers.App{}).AppPull)
}

func (v BsValidator) VAppUpdate(ctx *gin.Context) {
	var params v_data.VAppUpdate
	bindData(ctx, &params, (&handlers.App{}).AppUpdate)
}

func (v BsValidator) VAppCreate(ctx *gin.Context) {
	var params v_data.VAppCreate
	bindData(ctx, &params, (&handlers.App{}).AppCreate)
}

func (v BsValidator) VAppInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.App{}).AppInfo)
}
