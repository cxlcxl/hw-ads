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
			StatDay:          l.day,
			Country:          source.Country,
			AccountId:        source.AccountId,
			AppId:            source.AppId,
			AppName:          source.AppName,
			Cost:             source.Cost,
			ShowCount:        source.ShowCount,
			ClickCount:       source.ClickCount,
			DownloadCount:    source.DownloadCount,
			InstallCount:     source.InstallCount,
			ActivateCount:    source.ActivateCount,
			RetainCount:      source.RetainCount,
			ThreeRetainCount: source.ThreeRetainCount,
			SevenRetainCount: source.SevenRetainCount,
		}
	}
	return model.NewRMC(vars.DBMysql).BatchInsert(collects)
}
