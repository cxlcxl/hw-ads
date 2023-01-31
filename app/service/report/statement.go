package servicereport

type ReportColumn struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Align string `json:"align"`
	Show  bool   `json:"show"`
	Min   uint8  `json:"min"`
	Fix   bool   `json:"fix"`
}

var (
	// ComprehensiveColumns 综合报表展示字段
	ComprehensiveColumns = []*ReportColumn{
		{Key: "roi", Label: "ROI", Align: "right", Min: 100},
		{Key: "cost", Label: "花费($)", Align: "right", Min: 100},
		{Key: "earnings", Label: "收入($)", Align: "right", Min: 100},
		{Key: "ecpm", Label: "eCPM(变现:$)", Align: "right", Min: 100},
		{Key: "show_count", Label: "展示量", Align: "right", Min: 100},
		{Key: "cpm", Label: "CPM($)", Align: "right", Min: 100},
		{Key: "click_count", Label: "点击量", Align: "right", Min: 100},
		{Key: "cpc", Label: "CPC($)", Align: "right", Min: 100},
		{Key: "download_count", Label: "下载数", Align: "right", Min: 100},
		{Key: "cpd", Label: "CPD($)", Align: "right", Min: 100},
		{Key: "install_count", Label: "安装数", Align: "right", Min: 100},
		{Key: "cpi", Label: "CPI($)", Align: "right", Min: 100},
		{Key: "activate_count", Label: "激活量", Align: "right", Min: 100},
		{Key: "cpa", Label: "CPA($)", Align: "right", Min: 100},
		{Key: "click_through_rate", Label: "点击率", Align: "right", Min: 100},
		{Key: "download_activate_rate", Label: "激活转化率", Align: "right", Min: 100},
		{Key: "click_download_rate", Label: "下载转化率", Align: "right", Min: 100},
		{Key: "ad_requests", Label: "请求(变现)", Align: "right", Min: 100},
		{Key: "matched_ad_requests", Label: "填充(变现)", Align: "right", Min: 100},
		{Key: "ad_requests_match_rate", Label: "填充率(变现)", Align: "right", Min: 100},
		{Key: "ad_show_count", Label: "曝光量(变现)", Align: "right", Min: 100},
		{Key: "ad_click_count", Label: "点击量(变现)", Align: "right", Min: 100},
		{Key: "ad_requests_show_rate", Label: "展示率(变现)", Align: "right", Min: 100},
		{Key: "ad_click_through_rate", Label: "点击率(变现)", Align: "right", Min: 100},
		{Key: "arpu", Label: "ARPU(变现:$)", Align: "right", Min: 100},
	}
	AccountColumn                    = &ReportColumn{Key: "account_name", Label: "账户", Align: "left", Min: 120, Fix: true, Show: true}
	AppColumn                        = &ReportColumn{Key: "app_name", Label: "应用", Align: "center", Min: 100, Fix: true, Show: true}
	CountryColumn                    = &ReportColumn{Key: "region_country", Label: "区域&国家", Align: "left", Min: 120, Fix: true, Show: true}
	ComprehensiveMarketSQLColumnsMap = map[string]string{
		"cost":               "sum(`cost`) as `cost`",
		"show_count":         "sum(`show_count`) as `show_count`",
		"click_count":        "sum(`click_count`) as `click_count`",
		"install_count":      "sum(`install_count`) as `install_count`",
		"download_count":     "sum(`download_count`) as `download_count`",
		"activate_count":     "sum(`activate_count`) as `activate_count`",
		"retain_count":       "sum(`retain_count`) as `retain_count`",
		"three_retain_count": "sum(`three_retain_count`) as `three_retain_count`",
		"seven_retain_count": "sum(`seven_retain_count`) as `seven_retain_count`",
	}
	ComprehensiveAdsSQLColumnsMap = map[string]string{
		"earnings":            "sum(`earnings`) as `earnings`",
		"ad_requests":         "sum(`ad_requests`) as `ad_requests`",
		"matched_ad_requests": "sum(`matched_ad_requests`) as `matched_ad_requests`",
		"ad_show_count":       "sum(`show_count`) as `ad_show_count`",
		"ad_click_count":      "sum(`click_count`) as `ad_click_count`",
	}
)
