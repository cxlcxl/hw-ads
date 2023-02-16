package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"errors"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VRoleList(ctx *gin.Context) {
	var params v_data.VRoleList
	bindData(ctx, &params, (&handlers.RolePermission{}).RoleList)
}

func (v BsValidator) VRoleCreate(ctx *gin.Context) {
	var params v_data.VRoleCreate
	bindData(ctx, &params, (&handlers.RolePermission{}).RoleCreate)
}

func (v BsValidator) VRoleUpdate(ctx *gin.Context) {
	var params v_data.VRoleUpdate
	bindData(ctx, &params, (&handlers.RolePermission{}).RoleUpdate)
}

func (v BsValidator) VRoleInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.RolePermission{}).RoleInfo)
}

func (v BsValidator) VRolePermissions(ctx *gin.Context) {
	var params v_data.VRolePermissions
	bindData(ctx, &params, (&handlers.RolePermission{}).RolePermissions)
}

func (v BsValidator) VUserList(ctx *gin.Context) {
	var params v_data.VUserList
	bindData(ctx, &params, (&handlers.User{}).UserList)
}

func (v BsValidator) VUserCreate(ctx *gin.Context) {
	var params v_data.VUserCreate
	bindData(ctx, &params, (&handlers.User{}).UserCreate, isInternal)
}

func (v BsValidator) VUserUpdate(ctx *gin.Context) {
	var params v_data.VUserUpdate
	bindData(ctx, &params, (&handlers.User{}).UserUpdate, isInternal)
}

func (v BsValidator) VSelfUpdate(ctx *gin.Context) {
	var params v_data.VSelfUpdate
	bindData(ctx, &params, (&handlers.User{}).SelfUpdate, fillUser)
}

func (v BsValidator) VUserInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.User{}).UserInfo)
}

func (v BsValidator) VLogin(ctx *gin.Context) {
	var params v_data.VLogin
	bindData(ctx, &params, (&handlers.User{}).Login)
}

func (v BsValidator) VSsoLogin(ctx *gin.Context) {
	var params v_data.VSsoLoginData
	bindData(ctx, &params, (&handlers.Sso{}).ValidTicket)
}

func (v BsValidator) VResetPass(ctx *gin.Context) {
	var params v_data.VResetPass
	bindData(ctx, &params, (&handlers.User{}).ResetPass, fillUser)
}

func (v BsValidator) VPermissionCreate(ctx *gin.Context) {
	var params v_data.VPermissionCreate
	bindData(ctx, &params, (&handlers.RolePermission{}).PermissionCreate)
}

func (v BsValidator) VPermissionUpdate(ctx *gin.Context) {
	var params v_data.VPermissionUpdate
	bindData(ctx, &params, (&handlers.RolePermission{}).PermissionUpdate)
}

func (v BsValidator) VPermissionDestroy(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.RolePermission{}).PermissionDestroy)
}

func isInternal(_ *gin.Context, p interface{}) error {
	var (
		internal uint8
		markets  []int64
		ads      []int64
	)
	switch p.(type) {
	case *v_data.VUserCreate:
		internal = p.(*v_data.VUserCreate).IsInternal
		markets = p.(*v_data.VUserCreate).MarketAccounts
		ads = p.(*v_data.VUserCreate).AdsAccounts
		break
	case *v_data.VUserUpdate:
		internal = p.(*v_data.VUserUpdate).IsInternal
		markets = p.(*v_data.VUserUpdate).MarketAccounts
		ads = p.(*v_data.VUserUpdate).AdsAccounts
		break
	default:
		return errors.New("参数错误")
	}
	if internal == 0 {
		if len(markets) == 0 || len(ads) == 0 {
			return errors.New("外部用户需要绑定投放与变现账户用来限定数据")
		}
	}
	return nil
}
