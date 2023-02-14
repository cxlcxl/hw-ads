package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initSettingsApis(g *gin.RouterGroup) {
	g.GET("/settings/version", (&handlers.Settings{}).VersionInfo)

	group := g.Group("/settings", middleware.CheckPermission())
	{

		group.GET("/cron", (&handlers.Settings{}).Cron)
		group.GET("/cron/:id", (validator.BsValidator{}).VSettingsCronInfo)
		group.POST("/cron/:id", (validator.BsValidator{}).VSettingsCronUpdate)
		group.POST("/cron/schedule", (validator.BsValidator{}).VSettingsCronSchedule)

		group.GET("/configs", (validator.BsValidator{}).VSettingsConfigs)
		group.POST("/config", (validator.BsValidator{}).VSettingsConfigCreate)
		group.POST("/config/:id", (validator.BsValidator{}).VSettingsConfigUpdate)
		group.GET("/config/:id", (validator.BsValidator{}).VSettingsConfig)
	}
}
