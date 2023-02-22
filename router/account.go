package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initAccountApis(g *gin.RouterGroup) {
	g.Use(middleware.CheckUserLogin())
	{
		gp := g.Group("/account", middleware.CheckPermission())
		{
			gp.POST("/refresh/:id", (validator.BsValidator{}).VAccountRefreshToken)
			gp.POST("/create", (validator.BsValidator{}).VAccountCreate)
			gp.POST("/update", (validator.BsValidator{}).VAccountUpdate)

			gp.GET("/:id", (validator.BsValidator{}).VAccountInfo)
			gp.GET("/auth", (&handlers.Account{}).AccountAuth)
			gp.GET("/list", (validator.BsValidator{}).VAccountList)
		}

		group := g.Group("/account")
		{
			group.POST("/token", (validator.BsValidator{}).VAccountAuth)

			group.GET("/search", (validator.BsValidator{}).VAccountSearch)
			group.GET("/default", (validator.BsValidator{}).VAccountDefault)
			group.GET("/all", (validator.BsValidator{}).VAllAccounts)
			group.GET("/parents", (validator.BsValidator{}).VAccountParents)
		}
	}
}
