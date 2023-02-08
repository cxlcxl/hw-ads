package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/service/jwt"
	serviceuser "bs.mobgi.cc/app/service/user"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type User struct{}

func (l *User) UserList(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VUserList)
	offset := utils.GetPages(params.Page, params.PageSize)
	users, total, err := model.NewUser(vars.DBMysql).List(params.Username, params.Email, params.State, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{"total": total, "list": users})
}

func (l *User) Login(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VLogin)
	user, err := model.NewUser(vars.DBMysql).FindUserByEmail(params.Email)
	if err != nil {
		response.Fail(ctx, "登录失败: "+err.Error())
		return
	}
	if user.State != 1 {
		response.Fail(ctx, "账号已失效不可登陆")
		return
	}
	if user.Pass != utils.Password(params.Pass, user.Secret) {
		response.Fail(ctx, "密码错误")
		return
	}
	token, err := jwt.CreateUserToken(user.Id, user.RoleId, user.Email, user.Username, user.Mobile)
	if err != nil {
		response.Fail(ctx, "登录失败: "+err.Error())
		return
	}
	if _, err = jwt.ParseUserToken(token); err != nil {
		response.Fail(ctx, "TOKEN 生成失败: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{"token": token})
}

func (l User) Profile(ctx *gin.Context) {
	user, exists := ctx.Get(vars.LoginUserKey)
	if !exists {
		response.Fail(ctx, "用户信息获取失败")
		return
	}
	permissions, _ := model.NewPR(vars.DBMysql).GetRolePermissions(user.(*vars.LoginUser).RoleId)
	response.Success(ctx, gin.H{
		"user_id":     user.(*vars.LoginUser).UserId,
		"username":    user.(*vars.LoginUser).Username,
		"email":       user.(*vars.LoginUser).Email,
		"mobile":      user.(*vars.LoginUser).Mobile,
		"role_id":     user.(*vars.LoginUser).RoleId,
		"permissions": permissions,
	})
}

func (l User) Logout(ctx *gin.Context) {
	response.Success(ctx, nil)
}

func (l *User) UserInfo(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	user, err := model.NewUser(vars.DBMysql).FindUserById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	response.Success(ctx, user)
}

func (l *User) UserCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VUserCreate)
	s := utils.GenerateSecret(0)
	user := &model.User{
		Email:     params.Email,
		Username:  params.Username,
		Mobile:    params.Mobile,
		State:     1,
		RoleId:    params.RoleId,
		Secret:    s,
		Pass:      utils.Password(params.Pass, s),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := model.NewUser(vars.DBMysql).CreateUser(user)
	if err != nil {
		response.Fail(ctx, "创建失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *User) UserUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VUserUpdate)
	user, err := model.NewUser(vars.DBMysql).FindUserById(params.Id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	d := map[string]interface{}{
		"username":   params.Username,
		"email":      params.Email,
		"mobile":     params.Mobile,
		"role_id":    params.RoleId,
		"state":      params.State,
		"updated_at": time.Now(),
	}
	updatePass := false
	if params.Pass != "" {
		d["pass"] = utils.Password(params.Pass, user.Secret)
		updatePass = true
	}
	err = model.NewUser(vars.DBMysql).UpdateUser(d, params.Id)
	if err != nil {
		response.Fail(ctx, "修改失败："+err.Error())
		return
	}
	if updatePass && user.SsoUid != "" {
		go serviceuser.UpdatePass(user.SsoUid, params.Pass)
	}
	response.Success(ctx, nil)
}
