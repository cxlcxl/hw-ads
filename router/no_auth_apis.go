package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initNoAuthApis(r *gin.RouterGroup) {
	r.GET("/sso-login", (&handlers.Sso{}).SsoLoginPath)
	r.POST("/login", (validator.BsValidator{}).VLogin) //用户登陆
	r.POST("/sso-login", (validator.BsValidator{}).VSsoLogin)

	r.GET("/settings/log/:key", (validator.BsValidator{}).VSettingsLogDownload)
}
