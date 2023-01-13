package serviceaccount

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
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
