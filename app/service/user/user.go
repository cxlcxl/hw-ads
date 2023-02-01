package serviceuser

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/library/curl"
	"log"
	"time"

	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
)

func UpdatePass(uid, pass string) {
	curlPwd := map[string]interface{}{
		"uid":  uid,
		"pass": utils.MD5(pass),
	}
	c, err := curl.New(vars.YmlConfig.GetString("Sso.ForcePassHost")).Post().JsonData(curlPwd)
	if err != nil {
		log.Println("调用 sso 密码修改失败", err.Error())
		return
	}
	_ = c.Request(nil, curl.JsonHeader())
}

func SsoLogin(u *vars.SsoLoginResData) (user *model.User, err error) {
	user, err = model.NewUser(vars.DBMysql).FindUserByEmail(u.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		s := utils.GenerateSecret(0)
		user = &model.User{
			SsoUid:    u.SsoUid,
			Email:     u.Email,
			Username:  u.Username,
			Mobile:    u.Mobile,
			Secret:    s,
			State:     1,
			Pass:      utils.Password(vars.SystemDefaultPass, s),
			CreatedAt: time.Now(),
		}
		err = model.NewUser(vars.DBMysql).CreateUser(user)
	} else {
		if user.SsoUid == "" {
			d := map[string]interface{}{"sso_uid": u.SsoUid}
			err = model.NewUser(vars.DBMysql).UpdateUser(d, user.Id)
		}
	}
	return
}
