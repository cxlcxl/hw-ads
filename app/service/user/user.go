package serviceuser

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/library/curl"
	"errors"
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
	user, err = model.NewUser(vars.DBMysql).FindUserBySso(u.Email, u.SsoUid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		s := utils.GenerateSecret(0)
		user = &model.User{
			SsoUid:     u.SsoUid,
			Email:      u.Email,
			Username:   u.Username,
			Mobile:     u.Mobile,
			Secret:     s,
			State:      1,
			IsInternal: 1,
			Pass:       utils.Password(vars.SystemDefaultPass, s),
			CreatedAt:  time.Now(),
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

func ResetPass(params *v_data.VResetPass) error {
	if params.OldPass == params.Pass {
		return errors.New("新密码不能与旧密码一样")
	}
	user, err := model.NewUser(vars.DBMysql).FindUserById(params.User.UserId)
	if err != nil {
		return err
	}
	if utils.Password(params.OldPass, user.Secret) != user.Pass {
		return errors.New("旧密码错误")
	}
	password := utils.Password(params.Pass, user.Secret)
	err = model.NewUser(vars.DBMysql).UpdateUser(map[string]interface{}{"pass": password}, params.User.UserId)
	if err == nil && user.SsoUid != "" {
		go UpdatePass(user.SsoUid, params.Pass)
	}
	return err
}

func UserCreate(params *v_data.VUserCreate) (err error) {
	s := utils.GenerateSecret(0)
	user := &model.User{
		Email:      params.Email,
		Username:   params.Username,
		Mobile:     params.Mobile,
		State:      1,
		RoleId:     params.RoleId,
		IsInternal: params.IsInternal,
		Secret:     s,
		Pass:       utils.Password(params.Pass, s),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err = model.NewUser(vars.DBMysql).CreateUser(user); err != nil {
		return err
	}
	if params.IsInternal == 0 {
		userAccounts := buildUserActs(user.Id, params.MarketAccounts, params.AdsAccounts, params.IsInternal)
		if err = model.NewUserAccount(vars.DBMysql).SaveUserAccount(user.Id, userAccounts); err != nil {
			return err
		}
	}
	return nil
}

func UserUpdate(params *v_data.VUserUpdate) (err error) {
	user, err := model.NewUser(vars.DBMysql).FindUserById(params.Id)
	if err != nil {
		return
	}
	d := map[string]interface{}{
		"username":    params.Username,
		"email":       params.Email,
		"mobile":      params.Mobile,
		"role_id":     params.RoleId,
		"state":       params.State,
		"is_internal": params.IsInternal,
		"updated_at":  time.Now(),
	}
	updatePass := false
	if params.Pass != "" {
		d["pass"] = utils.Password(params.Pass, user.Secret)
		updatePass = true
	}
	err = model.NewUser(vars.DBMysql).UpdateUser(d, params.Id)
	if err != nil {
		return
	}
	if updatePass && user.SsoUid != "" {
		go UpdatePass(user.SsoUid, params.Pass)
	}
	userAccounts := buildUserActs(params.Id, params.MarketAccounts, params.AdsAccounts, params.IsInternal)
	err = model.NewUserAccount(vars.DBMysql).SaveUserAccount(params.Id, userAccounts)
	if err != nil {
		return
	}
	return nil
}

func BindAccount(user *model.User) map[string]interface{} {
	rs := map[string]interface{}{
		"id":          user.Id,
		"email":       user.Email,
		"username":    user.Username,
		"is_internal": user.IsInternal,
		"mobile":      user.Mobile,
		"role_id":     user.RoleId,
		"sso_uid":     user.SsoUid,
		"state":       user.State,
		"created_at":  user.CreatedAt.Format(vars.DateTimeFormat),
		"updated_at":  user.UpdatedAt.Format(vars.DateTimeFormat),
	}
	if user.IsInternal == 0 {
		acts, err := model.NewUserAccount(vars.DBMysql).FindAccountsByUserId(user.Id)
		if err == nil {
			var market, ads []int64
			for _, act := range acts {
				if act.AccountType == vars.AccountTypeMarket {
					market = append(market, act.AccountId)
				}
				if act.AccountType == vars.AccountTypeAds {
					ads = append(ads, act.AccountId)
				}
			}
			rs["market_accounts"] = market
			rs["ads_accounts"] = ads
		}
	}
	return rs
}

func buildUserActs(userId int64, market, ads []int64, internal uint8) []*model.UserAccount {
	userAccounts := make([]*model.UserAccount, 0)
	if internal == 1 {
		return userAccounts
	}
	for _, account := range market {
		userAccounts = append(userAccounts, &model.UserAccount{
			AccountId:   account,
			UserId:      userId,
			AccountType: vars.AccountTypeMarket,
		})
	}
	for _, account := range ads {
		userAccounts = append(userAccounts, &model.UserAccount{
			AccountId:   account,
			UserId:      userId,
			AccountType: vars.AccountTypeAds,
		})
	}
	return userAccounts
}
