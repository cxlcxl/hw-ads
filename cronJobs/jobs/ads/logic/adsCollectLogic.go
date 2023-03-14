package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/hlog"
	"github.com/sirupsen/logrus"
	"time"
)

type AdsCollectLogic struct {
	statDay string
	day     time.Time
	appMap  map[string][]int64
	actId   int64
}

func NewAdsCollectLogic(day string, accountId int64) *AdsCollectLogic {
	t, _ := time.Parse(vars.DateFormat, day)
	return &AdsCollectLogic{
		statDay: day,
		day:     t,
		appMap:  make(map[string][]int64),
		actId:   accountId,
	}
}

func (l *AdsCollectLogic) AdsCollect() (err error) {
	if err = l.getApps(); err != nil {
		return err
	}
	sources, err := model.NewRAS(vars.DBMysql).CollectSources(l.statDay, l.actId)
	if err != nil {
		return err
	}
	if len(sources) == 0 {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-collect-request").Log(logrus.Fields{
			"account_id": l.actId,
			"stat_day":   l.statDay,
		}, "没有请求到源数据")
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
					Ads: model.Ads{
						StatDay:           source.StatDay,
						Country:           source.Country,
						AdsAccountId:      source.AccountId,
						AccountId:         id,
						AppId:             source.AppId,
						AdRequests:        source.AdRequests,
						MatchedAdRequests: source.MatchedAdRequests,
						ShowCount:         source.ShowCount,
						ClickCount:        source.ClickCount,
						Earnings:          source.Earnings,
					},
				})
			}
		}
		collects = append(collects, &model.ReportAdsCollect{
			StatDay:           source.StatDay,
			Country:           source.Country,
			AppId:             source.AppId,
			AdsAccountId:      source.AccountId,
			AdRequests:        source.AdRequests,
			MatchedAdRequests: source.MatchedAdRequests,
			ShowCount:         source.ShowCount,
			ClickCount:        source.ClickCount,
			Earnings:          source.Earnings,
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
