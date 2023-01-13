package jobs

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/curl"
	"errors"
	"fmt"
	"log"
	"time"
)

type AccessToken struct {
	Error            int64  `json:"error"`
	ErrorDescription string `json:"error_description"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
}

type QueryParam struct {
	AccountId    int64
	AccessToken  string
	AdvertiserId string
}

type QueueData interface {
	GenerateMsg(fn func(interface{}))
}

func GetTokens(tokenChan chan *QueryParam) {
	list, err := model.NewToken(vars.DBMysql).GetAccessTokenList()
	if err != nil {
		return
	}
	for _, tokens := range list {
		if tokens.ExpiredAt.Before(time.Now()) {
			at, err := Refresh(tokens)
			if err != nil {
				log.Println("Token 刷新失败，账户 ID：", tokens.AccountId, err)
				continue
			}
			tokenChan <- &QueryParam{
				AccountId:    tokens.AccountId,
				AdvertiserId: tokens.AdvertiserId,
				AccessToken:  fmt.Sprintf("%s %s", tokens.TokenType, at),
			}
		} else {
			tokenChan <- &QueryParam{
				AccountId:    tokens.AccountId,
				AdvertiserId: tokens.AdvertiserId,
				AccessToken:  fmt.Sprintf("%s %s", tokens.TokenType, tokens.AccessToken),
			}
		}
	}

	close(tokenChan)
}

// Refresh token 过期时刷新（投放类型账户刷新方式）
func Refresh(token *model.Token) (t string, err error) {
	if token.RefreshToken == "" {
		return "", errors.New("没有 RefreshToken 无法刷新此 AccessToken")
	}
	data := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": token.RefreshToken,
		"client_id":     vars.YmlConfig.GetString("MarketingApis.Client.Id"),
		"client_secret": vars.YmlConfig.GetString("MarketingApis.Client.Secret"),
	}
	post := curl.New(vars.YmlConfig.GetString("MarketingApis.Refresh")).Post().QueryData(data)
	var at AccessToken
	err = post.Request(&at, curl.FormHeader())
	if err != nil {
		return "", err
	}
	if at.Error != 0 {
		return "", errors.New("华为接口调用失败：" + at.ErrorDescription)
	}
	_ = model.NewToken(vars.DBMysql).TokenUpdate(map[string]interface{}{
		"access_token": at.AccessToken,
		"expired_at":   time.Unix(time.Now().Unix()+at.ExpiresIn-20, 0).Format(vars.DateTimeFormat),
		"updated_at":   time.Now().Format(vars.DateTimeFormat),
		"token_type":   at.TokenType,
	}, token.Id)

	return at.AccessToken, nil
}

// RefreshV3 token 过期时刷新（变现类型账户刷新方式）
func RefreshV3(token *model.Token, clientId, secret string) (string, error) {
	if clientId == "" || secret == "" {
		return "", errors.New("ClientId/ClientSecret 信息不完整无法刷新 Token")
	}
	data := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     clientId,
		"client_secret": secret,
	}
	post := curl.New(vars.YmlConfig.GetString("MarketingApis.RefreshV3")).Debug(false).Post().QueryData(data)
	var at AccessToken
	if err := post.Request(&at, curl.FormHeader()); err != nil {
		return "", err
	}
	if at.Error != 0 {
		return "", errors.New("华为接口调用失败：" + at.ErrorDescription)
	}
	_ = model.NewToken(vars.DBMysql).TokenUpdate(map[string]interface{}{
		"token_type":   at.TokenType,
		"access_token": at.AccessToken,
		"expired_at":   time.Unix(time.Now().Unix()+at.ExpiresIn-20, 0).Format(vars.DateTimeFormat),
		"updated_at":   time.Now().Format(vars.DateTimeFormat),
	}, token.Id)

	return at.AccessToken, nil
}
