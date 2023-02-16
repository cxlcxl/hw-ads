package serviceexternal

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
)

func QueryAccounts(uid int64) (markets, ads []int64, err error) {
	userActs, err := model.NewUserAccount(vars.DBMysql).FindAccountsByUserId(uid)
	if err != nil {
		return
	}
	for _, act := range userActs {
		if act.AccountType == vars.AccountTypeAds {
			ads = append(ads, act.AccountId)
		}
		if act.AccountType == vars.AccountTypeMarket {
			markets = append(markets, act.AccountId)
		}
	}
	return
}

// Ads 取出外部用户绑定的变现数据账户
func Ads(searchActs []int64, uid int64) (ads []int64, err error) {
	var accounts []int64
	_, accounts, err = QueryAccounts(uid)
	if err != nil {
		return
	}
	if len(searchActs) == 0 {
		return accounts, nil
	}
	for _, act := range searchActs {
		if utils.InArray(act, ads) {
			ads = append(ads, act)
		}
	}
	return
}

// Markets 取出外部用户绑定的投放数据账户
func Markets(searchActs []int64, uid int64) (markets []int64, err error) {
	accounts, _, err := QueryAccounts(uid)
	if err != nil {
		return
	}
	if len(searchActs) == 0 {
		return accounts, err
	}
	for _, act := range searchActs {
		if utils.InArray(act, accounts) {
			markets = append(markets, act)
		}
	}
	return
}
