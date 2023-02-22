package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() error {
	r := gin.Default()

	if vars.YmlConfig.GetBool("HttpServer.AllowCrossDomain") {
		r.Use(corsNext())
	}

	group := r.Group(vars.ApiPrefix)
	{
		initNoAuthApis(group)

		group.Use(middleware.CheckUserLogin(), middleware.AccessLog())
		{
			initRbacApis(group)
			initAccountApis(group)
			initAppApis(group)

			initMarketingApis(group)
			initReportApis(group)
			initSettingsApis(group)

			group.GET("/regions", (&handlers.Region{}).Regions)
			g := group.Group("/region", middleware.CheckPermission())
			{
				g.GET("/area", (&handlers.Region{}).Areas)
				g.GET("/country", (validator.BsValidator{}).VCountries)
				g.POST("", (validator.BsValidator{}).VRegionCreate)
				g.POST("/area-set", (validator.BsValidator{}).VRegionAreaSet)
			}
		}
	}

	return r.Run(vars.YmlConfig.GetString("HttpServer.Port"))
}

// 允许跨域
func corsNext() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusAccepted)
		}
		c.Next()
	}
}
