package test

import (
	"bs.mobgi.cc/app/model"
	servicereport "bs.mobgi.cc/app/service/report"
	"bs.mobgi.cc/app/vars"
	_ "bs.mobgi.cc/bootstrap"
	"testing"
)

func TestAnalysisComprehensive(t *testing.T) {
	empty := []string{}
	dates := []string{"2023-01-27", "2023-01-27"}
	d, total, err := model.NewRMC(vars.DBMysql).ReportComprehensive(
		dates,
		[]int64{},
		empty, empty, servicereport.MarketSQLColumns, servicereport.AdsSQLColumns, []string{"stat_day", "account_id", "app_id"},
		empty, 0, 15,
	)
	t.Log(d, total, err)
}
