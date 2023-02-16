package handlers

import (
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/service/jwt"
	serviceuser "bs.mobgi.cc/app/service/user"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/curl"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Sso struct{}

func (l *Sso) SsoLoginPath(ctx *gin.Context) {
	host := vars.YmlConfig.GetString("Sso.LoginHost")
	appCode := vars.YmlConfig.GetString("Sso.AppCode")
	response.Success(ctx, fmt.Sprintf("%s%s", host, appCode))
}

type SsoLoginRes struct {
	Code    int64                 `json:"code"`
	Message string                `json:"message"`
	Data    *vars.SsoLoginResData `json:"data"`
}

func (l *Sso) ValidTicket(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VSsoLoginData)
	d := map[string]interface{}{
		"app_code": vars.YmlConfig.GetString("Sso.AppCode"),
		"ticket":   params.Ticket,
	}
	c, err := curl.New(vars.YmlConfig.GetString("Sso.UserHost")).Post().JsonData(d)
	if err != nil {
		response.Fail(ctx, "登陆请求失败："+err.Error())
		return
	}
	var res SsoLoginRes
	if err = c.Request(&res, curl.JsonHeader()); err != nil {
		response.Fail(ctx, "登陆失败："+err.Error())
		return
	}
	if res.Code != 200 {
		response.Fail(ctx, "登陆失败："+res.Message)
		return
	}
	user, err := serviceuser.SsoLogin(res.Data)
	if err != nil {
		response.Fail(ctx, "登陆失败："+err.Error())
		return
	}
	var token string
	if token, err = jwt.CreateUserToken(user.Id, user.RoleId, user.Email, user.Username, user.Mobile, 1); err == nil {
		_, err = jwt.ParseToken(token)
		if err != nil {
			response.Fail(ctx, err.Error())
			return
		}
		data := gin.H{
			"user_id":  user.Id,
			"username": res.Data.Username,
			"email":    res.Data.Email,
			"token":    token,
		}
		response.Success(ctx, data)
	}
}
