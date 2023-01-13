package router

import (
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initMarketingApis(g *gin.RouterGroup) {
	group := g.Group("/marketing")
	{
		group.GET("/resource/pricing", (validator.BsValidator{}).VResourcePricing)

		group.GET("/tracking/list", (validator.BsValidator{}).VTrackingList)
		group.GET("/tracking/refresh", (validator.BsValidator{}).VTrackingRefresh)

		group.GET("/dictionary/query", (validator.BsValidator{}).VDictQuery)

		initCampaignApis(group)
		initAssetApis(group)
		initTargetingApis(group)
		initPositionApis(group)
	}
}
