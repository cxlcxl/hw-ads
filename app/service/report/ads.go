package servicereport

import (
	"bs.mobgi.cc/app/model"
	serviceexternal "bs.mobgi.cc/app/service/external"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
)

// ReportAds 变现报表
func ReportAds(params *v_data.VReportAds) (data []*AdsReport, total int64, err error) {
	countries := formatCountries(params.Countries, params.Dimensions)
	offset := utils.GetPages(params.Page, params.PageSize)
	actIds := params.AccountIds
	if params.User.IsInternal == 0 {
		if actIds, err = serviceexternal.Ads(params.AccountIds, params.User.UserId); err != nil {
			return
		}
	}
	var _ads []*model.ReportAds
	_ads, total, err = model.NewRAS(vars.DBMysql).AnalysisAds(
		actIds, params.Areas, params.AppIds, params.DateRange, countries, _adsColumns(params.Dimensions),
		params.Dimensions, int(offset), int(params.PageSize),
	)
	if err != nil {
		return nil, 0, err
	}
	// 3. 数据整理
	data, err = formatAdsData(params, _ads)
	return
}

func _adsColumns(dimensions []string) []string {
	rs := append(AdsReportColumns)
	if utils.InArray("account_id", dimensions) {
		rs = append(rs, "account_id")
	}
	if utils.InArray("app_id", dimensions) {
		rs = append(rs, "app_id")
	}
	if utils.InArray("country", dimensions) {
		rs = append(rs, "country")
	}
	return rs
}

func formatAdsData(params *v_data.VReportAds, _ads []*model.ReportAds) (data []*AdsReport, err error) {
	// 3.1 检查是否需要填充账户名称，需要则填充
	_accountMap := make(map[int64]string)
	if utils.InArray("account_id", params.Dimensions) {
		_accountMap = accountMap(vars.AccountTypeAds)
	}
	// 3.2 检查是否需要填充国家地区，需要则填充
	_countryMap := make(map[string]*model.AreaCountry)
	if utils.InArray("country", params.Dimensions) {
		_countryMap = regionCountryMap()
	}
	_areaMap := make(map[int64]string)
	if utils.InArray("area_id", params.Dimensions) {
		_areaMap = areaMap()
	}
	// 3.3 检查应用
	_appMap := make(map[string]string)
	if utils.InArray("app_id", params.Dimensions) {
		_appMap = appMap()
	}
	data = make([]*AdsReport, len(_ads))
	for i, ads := range _ads {
		f0, f1, f2, f3 := calculateAdsRate(ads)
		appName := _appMap[ads.AppId]
		if appName == "" {
			appName = ads.AppId
		}
		adsReport := &AdsReport{
			StatDay:           ads.StatDay.Format(vars.DateFormat),
			Country:           ads.Country,
			AccountId:         ads.AccountId,
			AppId:             ads.AppId,
			AppName:           appName,
			AccountName:       _accountMap[ads.AccountId],
			AdRequests:        ads.AdRequests,
			MatchedAdRequests: ads.MatchedAdRequests,
			ShowCount:         ads.ShowCount,
			ClickCount:        ads.ClickCount,
			RequestsMatchRate: f0,
			RequestsShowRate:  f1,
			ClickThroughRate:  f2,
			Earnings:          ads.Earnings,
			ECpm:              f3,
			ARPU:              0,
		}
		if area, ok := _countryMap[ads.Country]; ok {
			adsReport.AreaName = area.AreaName
			adsReport.CountryName = area.CountryName
		}
		if area, ok := _areaMap[ads.AreaId]; ok {
			adsReport.AreaName = area
		} else {
			adsReport.AreaName = "-"
		}
		data[i] = adsReport
	}
	return
}

func appMap() map[string]string {
	rs := make(map[string]string)
	apps, err := model.NewApp(vars.DBMysql).AllApps(nil)
	if err != nil {
		return rs
	}
	for _, app := range apps {
		rs[app.AppId] = app.AppName
	}
	return rs
}

func ReportAdsColumns(columns, dimensions []string) (rs []*ReportColumn) {
	var forceShow bool
	if len(columns) == 0 {
		forceShow = true
	}
	rs = append(rs, DateColumn)
	if utils.InArray("account_id", dimensions) {
		rs = append(rs, AccountColumn)
	}
	if utils.InArray("app_id", dimensions) {
		rs = append(rs, AppColumn)
	}
	if utils.InArray("country", dimensions) {
		rs = append(rs, AreaColumn)
		rs = append(rs, CountryColumn)
	}
	if utils.InArray("area_id", dimensions) {
		rs = append(rs, AreaColumn)
	}
	for _, column := range AdsColumns {
		if forceShow || utils.InArray(column.Key, columns) {
			column.Show = true
		} else {
			column.Show = false
		}
		rs = append(rs, column)
	}
	return
}

func calculateAdsRate(ads *model.ReportAds) (rmr float64, rsr float64, ctr float64, ecpm float64) {
	if ads.AdRequests == 0 {
		rmr = 0
	} else {
		rmr = getRate(float64(ads.MatchedAdRequests)*100, ads.AdRequests, 2)
	}

	if ads.MatchedAdRequests == 0 {
		rsr = 0
	} else {
		rsr = getRate(float64(ads.ShowCount)*100, ads.MatchedAdRequests, 2)
	}

	if ads.ShowCount == 0 {
		ctr = 0
		ecpm = 0
	} else {
		ecpm = getRate(ads.Earnings*1000, ads.ShowCount, 4)
		ctr = getRate(float64(ads.ClickCount)*100, ads.ShowCount, 2)
	}
	return
}
