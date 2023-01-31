package servicereport

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"fmt"
)

// ReportComprehensive 综合报表
func ReportComprehensive(params *v_data.VReportComprehensive) (data interface{}, total int64, err error) {
	if len(params.ShowColumns) == 0 {
		params.ShowColumns = getFetchColumns()
	}

	groups := append(params.Dimensions, "stat_day")
	// 1. 汇总出投放报表数据
	markets, err := model.NewRMC(vars.DBMysql).AnalysisComprehensive(
		params.AppIds, params.DateRange, params.AccountIds,
		marketColumns(params.ShowColumns, params.Dimensions),
		groups,
	)
	if err != nil {
		return nil, 0, err
	}
	// 2. 汇总变现报表数据
	ads, err := model.NewRAC(vars.DBMysql).AnalysisComprehensive(
		params.AppIds, params.DateRange, params.AccountIds,
		adsColumns(params.ShowColumns, params.Dimensions),
		groups,
	)
	if err != nil {
		return nil, 0, err
	}
	// 3. 组合出数据整合的唯一条件
	// 4. 数据整理
	// 5. 分页返回
	fmt.Println(markets, ads)
	//utils.CeilPages()

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
