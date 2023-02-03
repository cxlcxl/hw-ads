package servicereport

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
)

// ReportAds 变现报表
func ReportAds(params *v_data.VReportAds) (data []*AdsReport, total int64, err error) {
	countries := formatCountries(params.Countries)
	var _ads []*model.Ads
	groups := make([]string, 0)
	for _, dimension := range params.Dimensions {
		if dimension != "account_id" {
			groups = append(groups, dimension)
		} else {
			groups = append(groups, "ads_account_id")
		}
	}
	offset := utils.GetPages(params.Page, params.PageSize)
	_ads, total, err = model.NewRAS(vars.DBMysql).AnalysisAds(
		params.AccountIds, params.AppIds, params.DateRange, countries, _adsColumns(params.Dimensions),
		groups, int(offset), int(params.PageSize),
	)
	if err != nil {
		return nil, 0, err
	}
	// 3. 数据整理
	data, err = formatAdsData(params, _ads)
	return
}

func _adsColumns(dimensions []string) []string {
	rs := append(AdsSQLColumnsMap, "stat_day")
	if utils.InArray("account_id", dimensions) {
		rs = append(rs, "ads_account_id")
	}
	if utils.InArray("app_id", dimensions) {
		rs = append(rs, "app_id")
	}
	if utils.InArray("country", dimensions) {
		rs = append(rs, "country")
	}
	return rs
}

func formatAdsData(params *v_data.VReportAds, _ads []*model.Ads) (data []*AdsReport, err error) {
	// 3.1 检查是否需要填充账户名称，需要则填充
	_accountMap := make(map[int64]string)
	if utils.InArray("account_id", params.Dimensions) {
		_accountMap = accountMap(vars.AccountTypeAds)
	}
	// 3.2 检查是否需要填充国家地区，需要则填充
	areaMap := make(map[string]string)
	if utils.InArray("country", params.Dimensions) {
		areaMap = regionCountryMap()
	}
	// 3.3 检查应用
	_appMap := make(map[string]string)
	if utils.InArray("app_id", params.Dimensions) {
		_appMap = appMap()
	}
	data = make([]*AdsReport, len(_ads))
	for i, ads := range _ads {
		f0, f1, f2, f3 := calculateAdsRate(ads)
		data[i] = &AdsReport{
			StatDay:           ads.StatDay.Format(vars.DateFormat),
			Country:           ads.Country,
			AccountId:         ads.AdsAccountId,
			AppId:             ads.AppId,
			AppName:           _appMap[ads.AppId],
			AccountName:       _accountMap[ads.AdsAccountId],
			RegionCountry:     areaMap[ads.Country],
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
	}
	return
}

func appMap() map[string]string {
	rs := make(map[string]string)
	apps, err := model.NewApp(vars.DBMysql).AllApps()
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
	if utils.InArray("account_id", dimensions) {
		rs = append(rs, AccountColumn)
	}
	if utils.InArray("app_id", dimensions) {
		rs = append(rs, AppColumn)
	}
	if utils.InArray("country", dimensions) {
		rs = append(rs, CountryColumn)
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

func calculateAdsRate(ads *model.Ads) (rmr float64, rsr float64, ctr float64, ecpm float64) {
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
