package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"time"
)

type Role struct{}

func (l *Role) RoleList(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VRoleList)
	roles, err := model.NewRole(vars.DBMysql).List(params.RoleName, params.State)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, roles)
}

func (l *Role) RoleCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VRoleCreate)
	role := &model.Role{
		RoleName:  params.RoleName,
		State:     1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := model.NewRole(vars.DBMysql).CreateRole(role)
	if err != nil {
		response.Fail(ctx, "创建失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}
func (l *Role) RoleUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VRoleUpdate)
	d := map[string]interface{}{
		"role_name": params.RoleName,
		"state":     params.State,
	}
	err := model.NewRole(vars.DBMysql).UpdateRole(d, params.Id)
	if err != nil {
		response.Fail(ctx, "创建失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}
func (l *Role) RoleInfo(ctx *gin.Context, v string) {
	response.Success(ctx, nil)
}
