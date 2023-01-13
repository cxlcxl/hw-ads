package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/job_data/statements"
	"fmt"
	"time"
)

var (
	positionElementTypes = []string{
		"image",
		"icon",
		"video",
		"title",
		"description",
		"corprate_name",
		"landing_page_url",
		"impression_tracking_url",
		"click_tracking_url",
	}
)

type requestPriceChan struct {
	Params statements.CreativeSizePriceRequest
	Token  string
}

type requestElementChan struct {
	Params statements.PositionElementRequest
	Token  string
}

func getTokenMap() (tokens map[int64]string, validAccountIds []int64, err error) {
	list, err := model.NewToken(vars.DBMysql).GetAccessTokenList()
	if err != nil {
		return nil, nil, err
	}
	tokens = make(map[int64]string)
	for _, token := range list {
		if token.ExpiredAt.After(time.Now()) {
			tokens[token.AccountId] = fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
			validAccountIds = append(validAccountIds, token.AccountId)
		}
	}
	return
}
