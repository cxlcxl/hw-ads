package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initTargetingApis(g *gin.RouterGroup) {
	group := g.Group("/targeting", middleware.CheckPermission())
	{
		group.GET("/list", (validator.BsValidator{}).VTargetingList)
		group.GET("/location", (&handlers.Targeting{}).Location)

		group.POST("/create", (validator.BsValidator{}).VTargetingCreate)
	}
}
