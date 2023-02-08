package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	serviceauth "bs.mobgi.cc/app/service/auth"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RolePermission struct{}

func (l *RolePermission) RoleList(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VRoleList)
	roles, err := model.NewRole(vars.DBMysql).List(params.RoleName, params.State)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, roles)
}

func (l *RolePermission) RoleCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VRoleCreate)
	role := &model.Role{
		RoleName: params.RoleName,
		State:    1,
	}
	err := model.NewRole(vars.DBMysql).CreateRole(role)
	if err != nil {
		response.Fail(ctx, "创建失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *RolePermission) RoleUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VRoleUpdate)
	err := serviceauth.RoleUpdate(params)
	if err != nil {
		response.Fail(ctx, "修改失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *RolePermission) RoleInfo(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误："+err.Error())
		return
	}
	role, err := model.NewRole(vars.DBMysql).FindRoleById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	permissions, err := model.NewPR(vars.DBMysql).GetRolePermissions(id)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	response.Success(ctx, gin.H{
		"id":          role.Id,
		"role_name":   role.RoleName,
		"state":       role.State,
		"sys":         role.Sys,
		"permissions": permissions,
	})
}

func (l *RolePermission) PermissionCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VPermissionCreate)
	err := model.NewPermission(vars.DBMysql).PermissionCreate(&model.Permission{
		Per:    params.Permission,
		PName:  params.PName,
		Method: params.Method,
		Pid:    params.Pid,
	})
	if err != nil {
		response.Fail(ctx, "添加失败:"+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *RolePermission) PermissionUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VPermissionUpdate)
	err := model.NewPermission(vars.DBMysql).PermissionUpdate(params.Id, map[string]interface{}{
		"permission": params.Permission,
		"p_name":     params.PName,
		"method":     params.Method,
		"pid":        params.Pid,
	})
	if err != nil {
		response.Fail(ctx, "修改失败:"+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *RolePermission) PermissionDestroy(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	if err = model.NewPermission(vars.DBMysql).PermissionDestroy(id); err != nil {
		response.Fail(ctx, "删除失败:"+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *RolePermission) PermissionTree(ctx *gin.Context) {
	permissions, err := model.NewPermission(vars.DBMysql).Permissions()
	if err != nil {
		response.Fail(ctx, "请求失败:"+err.Error())
		return
	}
	if len(permissions) == 0 {
		response.Success(ctx, []interface{}{})
		return
	}
	response.Success(ctx, tree(permissions, 0))
}

func tree(ps []*model.Permission, pid int64) (rs []*model.Permission) {
	for _, p := range ps {
		if p.Pid == pid {
			p.Children = tree(ps, p.Id)
			rs = append(rs, p)
		}
	}
	return
}

func (l *RolePermission) RolePermissions(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VRolePermissions)
	permissions, err := model.NewPR(vars.DBMysql).GetRolePermissions(params.Id)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	response.Success(ctx, permissions)
}
