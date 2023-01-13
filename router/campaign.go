package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initCampaignApis(g *gin.RouterGroup) {
	group := g.Group("/campaign")
	{
		group.GET("/resources", (&handlers.Campaign{}).Resources)
		group.GET("/list", (validator.BsValidator{}).VCampaignList)
		group.GET("/:campaign_id", (validator.BsValidator{}).VCampaignInfo)
	}
}
