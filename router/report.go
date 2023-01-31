package router

import (
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initReportApis(g *gin.RouterGroup) {
	group := g.Group("/report")
	{
		group.POST("/comprehensive", (validator.BsValidator{}).VReportComprehensive)
	}
}
