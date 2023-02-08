package router

import (
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initReportApis(g *gin.RouterGroup) {
	group := g.Group("/report", middleware.CheckPermission())
	{
		group.POST("/comprehensive", (validator.BsValidator{}).VReportComprehensive)
		group.POST("/ads", (validator.BsValidator{}).VReportAds)
		group.POST("/column", (validator.BsValidator{}).VReportColumn)
	}
}
