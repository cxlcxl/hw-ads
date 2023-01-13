package router

import (
	"bs.mobgi.cc/app/handlers/h_rbac"
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initRbacApis(r *gin.RouterGroup) {
	r.GET("/sso-login", (&h_rbac.Sso{}).SsoLoginPath)
	r.POST("/login", (validator.BsValidator{}).VLogin) //用户登陆
	r.POST("/sso-login", (validator.BsValidator{}).VSsoLogin)

	r.Use(middleware.CheckUserLogin())
	{
		r.GET("/profile", (&h_rbac.User{}).Profile) //个人信息
		r.POST("/logout", (&h_rbac.User{}).Logout)

		u := r.Group("/user")
		{
			//角色列表
			u.GET("/list", (validator.BsValidator{}).VUserList)
			//用户创建
			u.POST("/create", (validator.BsValidator{}).VUserCreate)
			//用户修改
			u.POST("/update", (validator.BsValidator{}).VUserUpdate)
			//用户信息
			u.GET("/:id", (validator.BsValidator{}).VUserInfo)
			//用户删除
			u.POST("/destroy")
		}
		role := r.Group("/role")
		{
			//角色列表
			role.GET("/list", (validator.BsValidator{}).VRoleList)
			//角色创建
			role.POST("/create", (validator.BsValidator{}).VRoleCreate)
			//角色修改
			role.POST("/update", (validator.BsValidator{}).VRoleUpdate)
			//角色信息
			role.GET("/:id", (validator.BsValidator{}).VRoleInfo)
			//角色删除
			role.POST("/destroy")
		}
		permission := r.Group("/permission")
		{
			permission.GET("/list", (validator.BsValidator{}).VPermissionList)
			permission.POST("/create")
			permission.POST("/update")
			permission.GET("/info")
			permission.POST("/destroy")
		}
	}
}
