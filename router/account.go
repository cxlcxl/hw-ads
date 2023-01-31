package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initAccountApis(g *gin.RouterGroup) {
	group := g.Group("/account")
	{
		group.POST("/token")
		group.POST("/refresh/:id", (validator.BsValidator{}).VAccountRefreshToken)
		group.POST("/create", (validator.BsValidator{}).VAccountCreate)
		group.POST("/update", (validator.BsValidator{}).VAccountUpdate)

		group.GET("/:id", (validator.BsValidator{}).VAccountInfo)
		group.GET("/auth", (&handlers.Account{}).AccountAuth)
		group.GET("/list", (validator.BsValidator{}).VAccountList)
		group.GET("/search", (validator.BsValidator{}).VAccountSearch)
		group.GET("/default", (validator.BsValidator{}).VAccountDefault)
		group.GET("/all", (validator.BsValidator{}).VAllAccounts)
		group.GET("/parents", (validator.BsValidator{}).VAccountParents)
	}
}
