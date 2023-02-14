package serviceapp

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
)

func GetAppBelongsAccounts(apps []*model.App) []*model.App {
	if len(apps) == 0 {
		return nil
	}
	appIds := make([]string, len(apps))
	for i, app := range apps {
		appIds[i] = app.AppId
	}

	vs, err := model.NewAppAct(vars.DBMysql).FindAccountIdsByAppIds(appIds)
	if err != nil {
		return apps
	}
	tmp := make(map[string][]int64)
	for _, v := range vs {
		tmp[v.AppId] = append(tmp[v.AppId], v.AccountId)
	}
	for i, app := range apps {
		if actIds, ok := tmp[app.AppId]; ok {
			apps[i].AccountIds = actIds
		}
	}
	return apps
}
