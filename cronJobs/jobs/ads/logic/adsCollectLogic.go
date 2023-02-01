package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"time"
)

type AdsCollectLogic struct {
	statDay string
	day     time.Time
	appMap  map[string][]int64
}

func NewAdsCollectLogic(day string) *AdsCollectLogic {
	t, _ := time.Parse(vars.DateFormat, day)
	return &AdsCollectLogic{
		statDay: day,
		day:     t,
		appMap:  make(map[string][]int64),
	}
}

func (l *AdsCollectLogic) AdsCollect() (err error) {
	if err = l.getApps(); err != nil {
		return err
	}
	sources, err := model.NewRAS(vars.DBMysql).CollectSources(l.statDay)
	if err != nil {
		return err
	}
	if len(sources) == 0 {
		return
	}
	collects := make([]*model.ReportAdsCollect, 0)
	collectActs := make([]*model.ReportAdsCollectAct, 0)
	for _, source := range sources {
		if actIds, ok := l.appMap[source.AppId]; !ok {
			continue
		} else {
			for _, id := range actIds {
				collectActs = append(collectActs, &model.ReportAdsCollectAct{
					Comprehensive: model.Comprehensive{
						StatDay:             source.StatDay,
						Country:             source.Country,
						AdsAccountId:        source.AccountId,
						AccountId:           id,
						AppId:               source.AppId,
						AdRequests:          source.AdRequests,
						MatchedAdRequests:   source.MatchedAdRequests,
						ShowCount:           source.ShowCount,
						ClickCount:          source.ClickCount,
						Earnings:            source.Earnings,
						AdRequestsMatchRate: getRate(float64(source.MatchedAdRequests), source.AdRequests, 4),
						AdRequestsShowRate:  getRate(float64(source.ShowCount), source.MatchedAdRequests, 4),
						ClickThroughRate:    getRate(float64(source.ClickCount), source.ShowCount, 4),
						ECpm:                getRate(source.Earnings*1000, source.ShowCount, 4),
					},
				})
			}
		}
		collects = append(collects, &model.ReportAdsCollect{
			StatDay:             source.StatDay,
			Country:             source.Country,
			AppId:               source.AppId,
			AdsAccountId:        source.AccountId,
			AdRequests:          source.AdRequests,
			MatchedAdRequests:   source.MatchedAdRequests,
			ShowCount:           source.ShowCount,
			ClickCount:          source.ClickCount,
			Earnings:            source.Earnings,
			AdRequestsMatchRate: getRate(float64(source.MatchedAdRequests), source.AdRequests, 4),
			AdRequestsShowRate:  getRate(float64(source.ShowCount), source.MatchedAdRequests, 4),
			ClickThroughRate:    getRate(float64(source.ClickCount), source.ShowCount, 4),
			ECpm:                getRate(source.Earnings*1000, source.ShowCount, 4),
		})

	}
	return model.NewRAC(vars.DBMysql).BatchInsert(collects, collectActs)
}

func (l *AdsCollectLogic) getApps() error {
	apps, err := model.NewAppAct(vars.DBMysql).CollectAdsApps()
	if err != nil {
		return err
	}
	for _, app := range apps {
		l.appMap[app.AppId] = append(l.appMap[app.AppId], app.AccountId)
	}
	return nil
}
