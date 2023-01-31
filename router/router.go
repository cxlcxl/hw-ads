package router

import (
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() error {
	r := gin.Default()

	if vars.YmlConfig.GetBool("HttpServer.AllowCrossDomain") {
		r.Use(corsNext())
	}

	group := r.Group("/api")
	{
		initRbacApis(group)
		initAccountApis(group)
		initAppApis(group)

		initMarketingApis(group)
		initReportApis(group)
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
