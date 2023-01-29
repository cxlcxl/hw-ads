package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"time"
)

type AdsCollectLogic struct {
	statDay string
	day     time.Time
}

func NewAdsCollectLogic(day string) *AdsCollectLogic {
	t, _ := time.Parse(vars.DateFormat, day)
	return &AdsCollectLogic{
		statDay: day,
		day:     t,
	}
}

func (l *AdsCollectLogic) AdsCollect() (err error) {
	sources, err := model.NewRAS(vars.DBMysql).CollectSources(l.statDay)
	if err != nil {
		return err
	}
	if len(sources) == 0 {
		return
	}
	collects := make([]*model.ReportAdsCollect, len(sources))
	for i, source := range sources {
		collects[i] = &model.ReportAdsCollect{
			StatDay:             l.day,
			Country:             source.Country,
			AccountId:           source.AccountId,
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
		}
	}
	return model.NewRAC(vars.DBMysql).BatchInsert(collects)
}
