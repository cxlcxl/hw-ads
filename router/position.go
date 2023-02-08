package router

import (
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/validator"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
)

func initPositionApis(g *gin.RouterGroup) {
	group := g.Group("/position", middleware.CheckPermission())
	{
		group.GET("category", func(ctx *gin.Context) {
			response.Success(ctx, vars.CreativeCategory)
		})

		group.GET("/query", (validator.BsValidator{}).VPositionQuery)
		group.GET("/placement", (validator.BsValidator{}).VPositionPlacement)
		group.GET("/price", (validator.BsValidator{}).VPositionPrice)
	}
}
