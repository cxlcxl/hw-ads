package validator

import (
	"bs.mobgi.cc/app/handlers/h_rbac"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VRoleList(ctx *gin.Context) {
	var params v_data.VRoleList
	bindData(ctx, &params, (&h_rbac.Role{}).RoleList)
}
func (v BsValidator) VRoleCreate(ctx *gin.Context) {
	var params v_data.VRoleCreate
	bindData(ctx, &params, (&h_rbac.Role{}).RoleCreate)
}
func (v BsValidator) VRoleUpdate(ctx *gin.Context) {
	var params v_data.VRoleUpdate
	bindData(ctx, &params, (&h_rbac.Role{}).RoleUpdate)
}
func (v BsValidator) VRoleInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&h_rbac.Role{}).RoleInfo)
}

func (v BsValidator) VUserList(ctx *gin.Context) {
	var params v_data.VUserList
	bindData(ctx, &params, (&h_rbac.User{}).UserList)
}

func (v BsValidator) VUserCreate(ctx *gin.Context) {
	var params v_data.VUserCreate
	bindData(ctx, &params, (&h_rbac.User{}).UserCreate)
}

func (v BsValidator) VUserUpdate(ctx *gin.Context) {
	var params v_data.VUserUpdate
	bindData(ctx, &params, (&h_rbac.User{}).UserUpdate)
}

func (v BsValidator) VUserInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&h_rbac.User{}).UserInfo)
}

func (v BsValidator) VLogin(ctx *gin.Context) {
	var params v_data.VLogin
	bindData(ctx, &params, (&h_rbac.User{}).Login)
}

func (v BsValidator) VPermissionList(ctx *gin.Context) {
	var params v_data.VPermissionList
	bindData(ctx, &params, (&h_rbac.User{}).PermissionList)
}

func (v BsValidator) VSsoLogin(ctx *gin.Context) {
	var params v_data.VSsoLoginData
	bindData(ctx, &params, (&h_rbac.Sso{}).ValidTicket)
}
