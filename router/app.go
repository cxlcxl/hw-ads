package router

import (
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initAppApis(g *gin.RouterGroup) {
	group := g.Group("/app")
	{
		group.POST("/pull", (validator.BsValidator{}).VAppPull)
		group.POST("/create", (validator.BsValidator{}).VAppCreate)
		group.POST("/update", (validator.BsValidator{}).VAppUpdate)

		group.GET("/:id", (validator.BsValidator{}).VAppInfo)
		group.GET("/list", (validator.BsValidator{}).VAppList)
		group.GET("/campaign-list", (validator.BsValidator{}).VAppCampaignList)
	}
}
