package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VRoleList(ctx *gin.Context) {
	var params v_data.VRoleList
	bindData(ctx, &params, (&handlers.Role{}).RoleList)
}
func (v BsValidator) VRoleCreate(ctx *gin.Context) {
	var params v_data.VRoleCreate
	bindData(ctx, &params, (&handlers.Role{}).RoleCreate)
}
func (v BsValidator) VRoleUpdate(ctx *gin.Context) {
	var params v_data.VRoleUpdate
	bindData(ctx, &params, (&handlers.Role{}).RoleUpdate)
}
func (v BsValidator) VRoleInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Role{}).RoleInfo)
}

func (v BsValidator) VUserList(ctx *gin.Context) {
	var params v_data.VUserList
	bindData(ctx, &params, (&handlers.User{}).UserList)
}

func (v BsValidator) VUserCreate(ctx *gin.Context) {
	var params v_data.VUserCreate
	bindData(ctx, &params, (&handlers.User{}).UserCreate)
}

func (v BsValidator) VUserUpdate(ctx *gin.Context) {
	var params v_data.VUserUpdate
	bindData(ctx, &params, (&handlers.User{}).UserUpdate)
}

func (v BsValidator) VUserInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.User{}).UserInfo)
}

func (v BsValidator) VLogin(ctx *gin.Context) {
	var params v_data.VLogin
	bindData(ctx, &params, (&handlers.User{}).Login)
}

func (v BsValidator) VPermissionList(ctx *gin.Context) {
	var params v_data.VPermissionList
	bindData(ctx, &params, (&handlers.User{}).PermissionList)
}

func (v BsValidator) VSsoLogin(ctx *gin.Context) {
	var params v_data.VSsoLoginData
	bindData(ctx, &params, (&handlers.Sso{}).ValidTicket)
}
