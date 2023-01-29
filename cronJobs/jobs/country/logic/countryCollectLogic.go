package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"time"
)

type CountryCollectLogic struct {
	statDay string
	day     time.Time
}

func NewCountryCollectLogic(day string) *CountryCollectLogic {
	t, _ := time.Parse(vars.DateFormat, day)
	return &CountryCollectLogic{
		statDay: day,
		day:     t,
	}
}

func (l *CountryCollectLogic) CountryCollect() (err error) {
	sources, err := model.NewRMS(vars.DBMysql).CollectSources(l.statDay)
	if err != nil {
		return err
	}
	if len(sources) == 0 {
		return
	}
	collects := make([]*model.ReportMarketCollect, len(sources))
	for i, source := range sources {
		collects[i] = &model.ReportMarketCollect{
			StatDay:              l.day,
			Country:              source.Country,
			AccountId:            source.AccountId,
			AppId:                source.AppId,
			AppName:              source.AppName,
			Cost:                 source.Cost,
			ShowCount:            source.ShowCount,
			ClickCount:           source.ClickCount,
			DownloadCount:        source.DownloadCount,
			InstallCount:         source.InstallCount,
			ActivateCount:        source.ActivateCount,
			RetainCount:          source.RetainCount,
			ThreeRetainCount:     source.ThreeRetainCount,
			SevenRetainCount:     source.SevenRetainCount,
			ClickThroughRate:     getRate(float64(source.ClickCount), source.ShowCount, 4),
			ClickDownloadRate:    getRate(float64(source.DownloadCount), source.ClickCount, 4),
			DownloadActivateRate: getRate(float64(source.ActivateCount), source.DownloadCount, 4),
			Cpm:                  getRate(source.Cost, source.ShowCount, 6),
			Cpc:                  getRate(source.Cost, source.ClickCount, 6),
			Cpd:                  getRate(source.Cost, source.DownloadCount, 6),
			Cpi:                  getRate(source.Cost, source.InstallCount, 6),
			Cpa:                  getRate(source.Cost, source.ActivateCount, 6),
			SevenRetainCost:      getRate(source.Cost, source.SevenRetainCount, 6),
			RetainCost:           getRate(source.Cost, source.RetainCount, 6),
			ThreeRetainCost:      getRate(source.Cost, source.ThreeRetainCount, 6),
		}
	}
	return model.NewRMC(vars.DBMysql).BatchInsert(collects)
}
