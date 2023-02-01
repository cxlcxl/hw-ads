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
	if len(params.ShowColumns) == 0 {
		params.ShowColumns = getFetchColumns()
	}

	// 1. 汇总出投放报表数据
	offset := utils.GetPages(params.Page, params.PageSize)
	markets, total, err := model.NewRMC(vars.DBMysql).AnalysisComprehensive(
		params.AccountIds, params.AppIds, params.DateRange,
		marketColumns(params.ShowColumns, params.Dimensions),
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
	var _ads []*model.Comprehensive
	groups := make([]string, 0)
	for _, dimension := range params.Dimensions {
		if dimension != "account_id" {
			groups = append(groups, dimension)
		}
	}
	if utils.InArray("account_id", params.Dimensions) {
		// 只包含国家分组，可以直接单查询
		_ads, err = model.NewRACC(vars.DBMysql).AnalysisComprehensive(
			params.AccountIds, appIds, params.DateRange, adsColumns(params.ShowColumns, params.Dimensions), groups,
		)
	} else {
		_ads, err = model.NewRAC(vars.DBMysql).AnalysisComprehensive(
			appIds, params.DateRange, adsColumns(params.ShowColumns, params.Dimensions), groups,
		)
	}
	if err != nil {
		return nil, 0, err
	}
	// 3. 数据整理
	data, err = formatComprehensiveData(params, markets, _ads)
	return
}

func getFetchColumns() (columns []string) {
	for _, column := range ComprehensiveColumns {
		columns = append(columns, column.Key)
	}
	return
}

func marketColumns(columns, dimensions []string) []string {
	rs := []string{"stat_day"}
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
	for _, column := range columns {
		if v, ok := ComprehensiveMarketSQLColumnsMap[column]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

func adsColumns(columns, dimensions []string) []string {
	rs := []string{"stat_day"}
	if utils.InArray("account_id", dimensions) {
		rs = append(rs, "account_id")
	} else {
		rs = append(rs, "0 as account_id")
	}
	if utils.InArray("app_id", dimensions) {
		rs = append(rs, "app_id")
	}
	if utils.InArray("country", dimensions) {
		rs = append(rs, "country")
	}
	for _, column := range columns {
		if v, ok := ComprehensiveAdsSQLColumnsMap[column]; ok {
			rs = append(rs, v)
		}
	}
	return rs
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
	var updates []string
	for _, column := range ComprehensiveColumns {
		var c *ReportColumn
		c = column
		if forceShow || utils.InArray(column.Key, columns) {
			c.Show = true
			updates = append(updates, c.Key)
		}
		rs = append(rs, c)
	}
	return
}

func adsWhere(params *v_data.VReportComprehensive, markets []*model.ReportMarketCollect) (appIds []string) {
	if utils.InArray("app_id", params.Dimensions) {
		tmp := make(map[string]struct{})
		for _, market := range markets {
			if _, ok := tmp[market.AppId]; ok {
				continue
			}
			tmp[market.AppId] = struct{}{}
			appIds = append(appIds, market.AppId)
		}
	}
	return
}

func accountMap() map[int64]string {
	rs := make(map[int64]string)
	accounts, err := model.NewAct(vars.DBMysql).AllAccounts()
	if err != nil {
		return rs
	}
	for _, account := range accounts {
		if account.AccountType == vars.AccountTypeMarket {
			rs[account.Id] = account.AccountName
		}
	}
	return rs
}

func formatComprehensiveData(
	params *v_data.VReportComprehensive, markets []*model.ReportMarketCollect, _ads []*model.Comprehensive,
) (data []*ComprehensiveReport, err error) {
	adsMap := adsFormat(_ads)
	// 3.1 检查是否需要填充账户名称，需要则填充
	_accountMap := make(map[int64]string)
	if utils.InArray("account_id", params.Dimensions) {
		_accountMap = accountMap()
	}
	for _, market := range markets {
		day := market.StatDay.Format(vars.DateFormat)
		unique := fmt.Sprintf("_%s_%s_%d_%s_", day, market.AppId, market.AccountId, market.Country)
		ads, ok := adsMap[unique]
		if !ok {
			ads = &model.Comprehensive{}
		}
		data = append(data, &ComprehensiveReport{
			StatDay:           day,
			Country:           market.Country,
			AccountId:         market.AccountId,
			AppId:             market.AppId,
			AppName:           market.AppName,
			AccountName:       _accountMap[market.AccountId],
			Cost:              market.Cost,
			ShowCount:         market.ShowCount,
			ClickCount:        market.ClickCount,
			DownloadCount:     market.DownloadCount,
			InstallCount:      market.InstallCount,
			ActivateCount:     market.ActivateCount,
			AdRequests:        ads.AdRequests,
			MatchedAdRequests: ads.MatchedAdRequests,
			AdShowCount:       ads.ShowCount,
			AdClickCount:      ads.ClickCount,
			Earnings:          utils.Round(ads.Earnings, 3),
		})
	}
	return
}

func adsFormat(_ads []*model.Comprehensive) map[string]*model.Comprehensive {
	rs := make(map[string]*model.Comprehensive)
	for _, ad := range _ads {
		unique := fmt.Sprintf("_%s_%s_%d_%s_", ad.StatDay.Format(vars.DateFormat), ad.AppId, ad.AccountId, ad.Country)
		rs[unique] = ad
	}
	return rs
}
