package servicereport

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"fmt"
)

// ReportComprehensive 综合报表
func ReportComprehensive(params *v_data.VReportComprehensive) (data []*ComprehensiveReport, total int64, err error) {
	countries := formatCountries(params.Countries)
	// 1. 汇总出投放报表数据
	offset := utils.GetPages(params.Page, params.PageSize)
	markets, total, err := model.NewRMC(vars.DBMysql).AnalysisComprehensive(
		params.AccountIds, params.AppIds, params.DateRange, countries,
		marketColumns(params.Dimensions),
		params.Dimensions, int(offset), int(params.PageSize),
	)
	if err != nil {
		return nil, 0, err
	}
	if len(markets) == 0 {
		return
	}
	// 2. 汇总变现报表数据
	appIds := adsWhere(params, markets)
	var _ads []*model.Ads
	groups := make([]string, 0)
	for _, dimension := range params.Dimensions {
		if dimension != "account_id" {
			groups = append(groups, dimension)
		}
	}
	if utils.InArray("account_id", params.Dimensions) {
		// 只包含国家分组，可以直接单查询
		_ads, err = model.NewRACC(vars.DBMysql).AnalysisComprehensive(
			params.AccountIds, appIds, params.DateRange, countries, adsColumns(params.Dimensions), groups,
		)
	} else {
		_ads, err = model.NewRAC(vars.DBMysql).AnalysisComprehensive(
			appIds, params.DateRange, countries, adsColumns(params.Dimensions), groups,
		)
	}
	if err != nil {
		return nil, 0, err
	}
	// 3. 数据整理
	data, err = formatComprehensiveData(params, markets, _ads)
	return
}

func marketColumns(dimensions []string) []string {
	rs := append(ComprehensiveMarketSQLColumnsMap, "stat_day")
	if utils.InArray("account_id", dimensions) {
		rs = append(rs, "account_id")
	}
	if utils.InArray("app_id", dimensions) {
		rs = append(rs, "app_id")
		rs = append(rs, "app_name")
	}
	if utils.InArray("country", dimensions) {
		rs = append(rs, "country")
	}
	return rs
}

func formatComprehensiveData(
	params *v_data.VReportComprehensive, markets []*model.ReportMarketCollect, _ads []*model.Ads,
) (data []*ComprehensiveReport, err error) {
	adsMap := adsFormat(_ads)
	// 3.1 检查是否需要填充账户名称，需要则填充
	_accountMap := make(map[int64]string)
	if utils.InArray("account_id", params.Dimensions) {
		_accountMap = accountMap(vars.AccountTypeMarket)
	}
	// 3.2 检查是否需要填充国家地区，需要则填充
	areaMap := make(map[string]string)
	if utils.InArray("country", params.Dimensions) {
		areaMap = regionCountryMap()
	}
	for _, market := range markets {
		day := market.StatDay.Format(vars.DateFormat)
		unique := fmt.Sprintf("_%s_%s_%d_%s_", day, market.AppId, market.AccountId, market.Country)
		ads, ok := adsMap[unique]
		if !ok {
			ads = &model.Ads{}
		}
		var roi float64
		if market.Cost == 0 {
			if ads.Earnings == 0 {
				roi = 0
			} else {
				roi = 100
			}
		} else {
			roi = utils.Round(ads.Earnings/market.Cost*100, 2)
		}

		calculateMarketingRateCost(market)
		calculateMediationRateEarnings(ads)
		data = append(data, &ComprehensiveReport{
			StatDay:              day,
			Country:              market.Country,
			AccountId:            market.AccountId,
			AppId:                market.AppId,
			AppName:              market.AppName,
			AccountName:          _accountMap[market.AccountId],
			RegionCountry:        areaMap[market.Country],
			ROI:                  roi,
			Cost:                 market.Cost,
			ShowCount:            market.ShowCount,
			ClickCount:           market.ClickCount,
			DownloadCount:        market.DownloadCount,
			InstallCount:         market.InstallCount,
			ActivateCount:        market.ActivateCount,
			Cpm:                  market.Cpm,
			Cpd:                  market.Cpd,
			Cpi:                  market.Cpi,
			Cpa:                  market.Cpa,
			Cpc:                  market.Cpc,
			ClickThroughRate:     market.ClickThroughRate,
			ClickDownloadRate:    market.ClickDownloadRate,
			DownloadActivateRate: market.DownloadActivateRate,
			RetainCost:           market.RetainCost,
			AdRequests:           ads.AdRequests,
			MatchedAdRequests:    ads.MatchedAdRequests,
			AdShowCount:          ads.ShowCount,
			AdClickCount:         ads.ClickCount,
			Earnings:             utils.Round(ads.Earnings, 3),
			ECpm:                 ads.ECpm,
			AdRequestsMatchRate:  ads.AdRequestsMatchRate,
			AdRequestsShowRate:   ads.AdRequestsShowRate,
			AdClickThroughRate:   ads.ClickThroughRate,
		})
	}
	return
}

func adsFormat(_ads []*model.Ads) map[string]*model.Ads {
	rs := make(map[string]*model.Ads)
	for _, ad := range _ads {
		unique := fmt.Sprintf("_%s_%s_%d_%s_", ad.StatDay.Format(vars.DateFormat), ad.AppId, ad.AccountId, ad.Country)
		rs[unique] = ad
	}
	return rs
}

func getRate(a float64, b, c int64) float64 {
	if b == 0 {
		return 0
	}
	return utils.Round(a/float64(b), int(c))
}

func calculateMarketingRateCost(market *model.ReportMarketCollect) {
	// 点击率
	if market.ShowCount == 0 {
		market.ClickThroughRate = 0
	} else {
		market.ClickThroughRate = getRate(float64(market.ClickCount)*100, market.ShowCount, 2)
	}

	// 点击下载率
	if market.ClickCount == 0 {
		market.ClickDownloadRate = 0
	} else {
		market.ClickDownloadRate = getRate(float64(market.DownloadCount)*100, market.ClickCount, 2)
	}
	// 下载激活率
	if market.DownloadCount == 0 {
		market.DownloadActivateRate = 0
	} else {
		market.DownloadActivateRate = getRate(float64(market.ActivateCount)*100, market.DownloadCount, 2)
	}

	if market.ShowCount == 0 {
		market.Cpm = 0
	} else {
		market.Cpm = getRate(market.Cost, market.ShowCount, 6)
	}

	if market.ClickCount == 0 {
		market.Cpc = 0
	} else {
		market.Cpc = getRate(market.Cost, market.ClickCount, 6)
	}

	if market.DownloadCount == 0 {
		market.Cpd = 0
	} else {
		market.Cpd = getRate(market.Cost, market.DownloadCount, 6)
	}

	if market.InstallCount == 0 {
		market.Cpi = 0
	} else {
		market.Cpi = getRate(market.Cost, market.InstallCount, 6)
	}

	if market.ActivateCount == 0 {
		market.Cpa = 0
	} else {
		market.Cpa = getRate(market.Cost, market.ActivateCount, 6)
	}

	if market.RetainCount == 0 {
		market.RetainCost = 0
	} else {
		market.RetainCost = getRate(market.Cost, market.RetainCount, 6)
	}
}

func calculateMediationRateEarnings(ads *model.Ads) {
	if ads.AdRequests == 0 {
		ads.AdRequestsMatchRate = 0
	} else {
		ads.AdRequestsMatchRate = getRate(float64(ads.MatchedAdRequests)*100, ads.AdRequests, 2)
	}

	if ads.MatchedAdRequests == 0 {
		ads.AdRequestsShowRate = 0
	} else {
		ads.AdRequestsShowRate = getRate(float64(ads.ShowCount)*100, ads.MatchedAdRequests, 2)
	}

	if ads.ShowCount == 0 {
		ads.ClickThroughRate = 0
		ads.ECpm = 0
	} else {
		ads.ECpm = getRate(ads.Earnings*1000, ads.ShowCount, 4)
		ads.ClickThroughRate = getRate(float64(ads.ClickCount)*100, ads.ShowCount, 2)
	}
}

func ReportComprehensiveColumns(columns, dimensions []string) (rs []*ReportColumn) {
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
	for _, column := range ComprehensiveColumns {
		if forceShow || utils.InArray(column.Key, columns) {
			column.Show = true
		} else {
			column.Show = false
		}
		rs = append(rs, column)
	}
	return
}
