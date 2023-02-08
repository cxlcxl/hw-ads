package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initAssetApis(g *gin.RouterGroup) {
	group := g.Group("/asset", middleware.CheckPermission())
	{
		group.GET("/list", (validator.BsValidator{}).VAssetList)
		group.GET("/dimension", (&handlers.Asset{}).AssetDimension)
		group.GET("/sync", (validator.BsValidator{}).VAssetSync)
	}
}
