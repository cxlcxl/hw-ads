package test

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	_ "bs.mobgi.cc/bootstrap"
	"testing"
)

func TestAnalysisComprehensive(t *testing.T) {
	empty := []string{}
	dates := []string{"2023-01-27", "2023-01-27"}
	selected := []string{
		"stat_day", "app_id",
		"round(sum(`earnings`), 3) as `earnings`",
		"sum(`ad_requests`) as `ad_requests`",
		"sum(`matched_ad_requests`) as `matched_ad_requests`",
		"sum(`show_count`) as `show_count`",
		"sum(`click_count`) as `click_count`",
	}
	comprehensive, err := model.NewRAC(vars.DBMysql).AnalysisComprehensive(empty, dates, selected, empty, empty)
	t.Log(comprehensive, err)
}
