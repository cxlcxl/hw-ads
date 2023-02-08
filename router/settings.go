package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initSettingsApis(g *gin.RouterGroup) {
	group := g.Group("/settings", middleware.CheckPermission())
	{
		group.GET("/cron", (&handlers.Settings{}).Cron)
		group.GET("/cron/:id", (validator.BsValidator{}).VSettingsCronInfo)
		group.POST("/cron/:id", (validator.BsValidator{}).VSettingsCronUpdate)
		group.POST("/cron/schedule", (validator.BsValidator{}).VSettingsCronSchedule)
	}
}
