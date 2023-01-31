package serviceaccount

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/curl"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

func GetToken(accountId int64) (*vars.AccountTokenInfo, error) {
	if token, err := model.NewToken(vars.DBMysql).FindByAccountId(accountId); err != nil {
		return nil, err
	} else {
		return &vars.AccountTokenInfo{
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
	state := utils.MD5(fmt.Sprintf("%d-%s-%s", id, clientId, time.Now().String()))
	baseUrl := vars.YmlConfig.GetString("MarketingApis.Authorize.CodeUrl")
	if !strings.HasSuffix(baseUrl, "?") {
		baseUrl += "?"
	}
	params := curl.HttpBuildQuery(map[string]string{
		"response_type": "code",
		"access_type":   "offline",
		"client_id":     clientId,
		"state":         state,
		"redirect_uri":  vars.YmlConfig.GetString("MarketingApis.Authorize.RedirectUri"),
		"scope":         vars.YmlConfig.GetString("MarketingApis.Authorize.Scope"),
	})
	authorizeValue := fmt.Sprintf("%d-%s-%s", id, clientId, secret)
	if !vars.DBRedis.SetString("authorize:"+state, authorizeValue, 300) {
		return "", err
	}
	return baseUrl + params, nil
}
