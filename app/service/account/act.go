package serviceaccount

import (
	"bs.mobgi.cc/app/model"
	serviceexternal "bs.mobgi.cc/app/service/external"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/curl"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type AccountTokenInfo struct {
	AccountId    int64  `json:"account_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int64  `json:"expired_at"`
	TokenType    string `json:"token_type"`
}

func GetToken(accountId int64) (*AccountTokenInfo, error) {
	if token, err := model.NewToken(vars.DBMysql).FindByAccountId(accountId); err != nil {
		return nil, err
	} else {
		return &AccountTokenInfo{
			AccountId:    accountId,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			ExpiredAt:    token.ExpiredAt.Unix(),
			TokenType:    token.TokenType,
		}, nil
	}
}

func AuthorizeCodeUrl(id int64, clientId, secret string) (url string, err error) {
	token, err := GetToken(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	if token != nil && time.Now().Before(time.Unix(token.ExpiredAt-300, 0)) {
		return "", errors.New("TOKEN 尚未过期，无需重新认证")
	}
	state := utils.MD5(fmt.Sprintf("%d-%s-%d", id, clientId, time.Now().Unix()))
	baseUrl := vars.YmlConfig.GetString("MarketingApis.Authorize.CodeUrl")
	if !strings.HasSuffix(baseUrl, "?") {
		baseUrl += "?"
	}
	conf, err := model.NewSysConfig(vars.DBMysql).FindOneByKey("MarketingApis.Authorize.RedirectUri")
	if err != nil {
		return "", errors.New("获取重定向授权地址失败，请检查是否填写地址配置")
	}
	params := curl.HttpBuildQuery(map[string]string{
		"response_type": "code",
		"access_type":   "offline",
		"client_id":     clientId,
		"state":         state,
		"redirect_uri":  conf.Val,
		"scope":         vars.YmlConfig.GetString("MarketingApis.Authorize.Scope"),
	})
	authorizeValue := fmt.Sprintf("%d-%s-%s", id, clientId, secret)
	if !vars.DBRedis.SetString("authorize:"+state, authorizeValue, 300) {
		return "", err
	}
	return baseUrl + params, nil
}

func SetToken(token *model.Token) (err error) {
	t, err := model.NewToken(vars.DBMysql).FindByAccountId(token.AccountId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if t == nil {
		return model.NewToken(vars.DBMysql).TokenCreate(token)
	} else {
		v := map[string]interface{}{
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
			"expired_at":    token.ExpiredAt,
			"updated_at":    token.UpdatedAt,
			"token_type":    token.TokenType,
		}
		return model.NewToken(vars.DBMysql).TokenUpdate(v, t.Id)
	}
}

// ExternalUserAccountIds 外部用户查询与之绑定的账号ID「投放与变现均返回」
func ExternalUserAccountIds(internal uint8, userId int64) (accountIds []int64, err error) {
	if internal == 1 {
		return
	}
	accounts, ads, err := serviceexternal.QueryAccounts(userId)
	if err != nil {
		return
	}
	accountIds = append(accounts, ads...)
	return
}
