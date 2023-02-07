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
	countries := formatCountries(params.Countries, params.Dimensions)
	// 1. 汇总报表数据
	offset := utils.GetPages(params.Page, params.PageSize)
	groups := append(params.Dimensions, "stat_day")
	comprehensives, total, err := model.NewRMC(vars.DBMysql).ReportComprehensive(
		params.DateRange, params.AccountIds, params.AppIds, countries,
		marketColumns(params.Dimensions), adsColumns(params.Dimensions),
		groups, comprehensiveOrders(params.Order, params.By), uint64(offset), uint64(params.PageSize),
	)
	if err != nil {
		return nil, 0, err
	}
	if len(comprehensives) == 0 {
		return
	}
	// 2. 数据整理
	data, err = formatComprehensiveData(params, comprehensives)
	return
}

func formatComprehensiveData(
	params *v_data.VReportComprehensive, comprehensives []*model.ComprehensiveReport,
) (data []*ComprehensiveReport, err error) {
	// 3.1 检查是否需要填充账户名称，需要则填充
	_accountMap := make(map[int64]string)
	if utils.InArray("account_id", params.Dimensions) {
		_accountMap = accountMap(vars.AccountTypeMarket)
	}
	// 3.2 检查是否需要填充国家地区，需要则填充
	areaMap := make(map[string]*model.AreaCountry)
	if utils.InArray("country", params.Dimensions) {
		areaMap = regionCountryMap()
	}
	for _, report := range comprehensives {
		var roi float64
		if report.Cost == 0 {
			if report.Earnings == 0 {
				roi = 0
			} else {
				roi = 100
			}
		} else {
			roi = utils.Round(report.Earnings/report.Cost*100, 2)
		}
		area, ok := areaMap[report.Country]
		if !ok {
			area = &model.AreaCountry{}
		}
		comprehensiveReport := &ComprehensiveReport{
			StatDay:           report.StatDay.Format(vars.DateFormat),
			Country:           report.Country,
			AccountId:         report.AccountId,
			AppId:             report.AppId,
			AppName:           report.AppName,
			AccountName:       _accountMap[report.AccountId],
			AreaName:          area.AreaName,
			CountryName:       area.CountryName,
			ROI:               roi,
			Cost:              report.Cost,
			ShowCount:         report.ShowCount,
			ClickCount:        report.ClickCount,
			DownloadCount:     report.DownloadCount,
			InstallCount:      report.InstallCount,
			ActivateCount:     report.ActivateCount,
			AdRequests:        report.AdRequests,
			MatchedAdRequests: report.MatchedAdRequests,
			AdShowCount:       report.ShowCount,
			AdClickCount:      report.ClickCount,
			Earnings:          utils.Round(report.Earnings, 3),
		}
		calculateRates(report, comprehensiveReport)
		data = append(data, comprehensiveReport)
	}
	return
}

func marketColumns(dimensions []string) (rs []string) {
	rs = append(MarketSQLColumns)
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

func getRate(a float64, b, c int64) float64 {
	if b == 0 {
		return 0
	}
	return utils.Round(a/float64(b), int(c))
}

func calculateRates(report *model.ComprehensiveReport, v *ComprehensiveReport) {
	if report.AdRequests == 0 {
		v.AdRequestsMatchRate = 0
	} else {
		v.AdRequestsMatchRate = getRate(float64(report.MatchedAdRequests)*100, report.AdRequests, 2)
	}

	if report.MatchedAdRequests == 0 {
		v.AdRequestsShowRate = 0
	} else {
		v.AdRequestsShowRate = getRate(float64(report.AdShowCount)*100, report.MatchedAdRequests, 2)
	}

	if report.AdShowCount == 0 {
		v.AdClickThroughRate = 0
		v.ECpm = 0
	} else {
		v.ECpm = getRate(report.Earnings*1000, report.AdShowCount, 4)
		v.AdClickThroughRate = getRate(float64(report.AdClickCount)*100, report.AdShowCount, 2)
	}
	// 点击率
	if report.ShowCount == 0 {
		v.ClickThroughRate = 0
	} else {
		v.ClickThroughRate = getRate(float64(report.ClickCount)*100, report.ShowCount, 2)
	}
	// 点击下载率
	if report.ClickCount == 0 {
		v.ClickDownloadRate = 0
	} else {
		v.ClickDownloadRate = getRate(float64(report.DownloadCount)*100, report.ClickCount, 2)
	}
	// 下载激活率
	if report.DownloadCount == 0 {
		v.DownloadActivateRate = 0
	} else {
		v.DownloadActivateRate = getRate(float64(report.ActivateCount)*100, report.DownloadCount, 2)
	}

	if report.ShowCount == 0 {
		v.Cpm = 0
	} else {
		v.Cpm = getRate(report.Cost, report.ShowCount, 6)
	}

	if report.ClickCount == 0 {
		v.Cpc = 0
	} else {
		v.Cpc = getRate(report.Cost, report.ClickCount, 6)
	}

	if report.DownloadCount == 0 {
		v.Cpd = 0
	} else {
		v.Cpd = getRate(report.Cost, report.DownloadCount, 6)
	}

	if report.InstallCount == 0 {
		v.Cpi = 0
	} else {
		v.Cpi = getRate(report.Cost, report.InstallCount, 6)
	}

	if report.ActivateCount == 0 {
		v.Cpa = 0
	} else {
		v.Cpa = getRate(report.Cost, report.ActivateCount, 6)
	}

	if report.RetainCount == 0 {
		v.RetainCost = 0
	} else {
		v.RetainCost = getRate(report.Cost, report.RetainCount, 6)
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
		rs = append(rs, AreaColumn)
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

func comprehensiveOrders(order, by string) (orders []string) {
	_order, ok := sortableColumns[order]
	if !ok {
		return
	}
	if _by, ok := orderBy[by]; ok {
		orders = append(orders, fmt.Sprintf("%s %s", _order, _by))
	}
	return
}
