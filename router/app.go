package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initAppApis(g *gin.RouterGroup) {
	group := g.Group("/app")
	{
		group.GET("/all", (&handlers.App{}).AllApp)
		group.GET("/campaign-list", (validator.BsValidator{}).VAppCampaignList)
		group.GET("/relation", (&handlers.App{}).AppRelations)
	}
	gp := g.Group("/app", middleware.CheckPermission())
	{
		gp.POST("/pull", (validator.BsValidator{}).VAppPull)
		gp.POST("/create", (validator.BsValidator{}).VAppCreate)
		gp.POST("/update", (validator.BsValidator{}).VAppUpdate)

		gp.GET("/:id", (validator.BsValidator{}).VAppInfo)
		gp.POST("/list", (validator.BsValidator{}).VAppList)
	}
}
