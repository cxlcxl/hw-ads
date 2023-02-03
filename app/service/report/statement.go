package servicereport

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
)

type ReportColumn struct {
	Key    string `json:"key"`
	Label  string `json:"label"`
	Align  string `json:"align"`
	Show   bool   `json:"show"`
	Min    uint8  `json:"min"`
	Fix    bool   `json:"fix"`
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

// ComprehensiveReport 最终综合报表页展示的数据字段「需要和 ComprehensiveColumns 的 key 一致」
type ComprehensiveReport struct {
	StatDay              string  `json:"stat_day"`
	Country              string  `json:"country"`
	AccountId            int64   `json:"account_id"`
	AppId                string  `json:"app_id"`
	AppName              string  `json:"app_name"`
	AccountName          string  `json:"account_name"`
	RegionCountry        string  `json:"region_country"`
	ROI                  float64 `json:"roi"`
	Cost                 float64 `json:"cost"`
	ShowCount            int64   `json:"show_count"`
	ClickCount           int64   `json:"click_count"`
	DownloadCount        int64   `json:"download_count"`
	InstallCount         int64   `json:"install_count"`
	ActivateCount        int64   `json:"activate_count"`
	RetainCount          int64   `json:"retain_count"`
	ThreeRetainCount     int64   `json:"three_retain_count"`
	SevenRetainCount     int64   `json:"seven_retain_count"`
	ClickThroughRate     float64 `json:"click_through_rate"`
	ClickDownloadRate    float64 `json:"click_download_rate"`
	DownloadActivateRate float64 `json:"download_activate_rate"`
	Cpm                  float64 `json:"cpm"`
	Cpc                  float64 `json:"cpc"`
	Cpd                  float64 `json:"cpd"`
	Cpi                  float64 `json:"cpi"`
	Cpa                  float64 `json:"cpa"`
	SevenRetainCost      float64 `json:"seven_retain_cost"`
	RetainCost           float64 `json:"retain_cost"`
	ThreeRetainCost      float64 `json:"three_retain_cost"`
	AdRequests           int64   `json:"ad_requests"`             // 到达服务器的请求数量
	MatchedAdRequests    int64   `json:"matched_ad_requests"`     // 匹配到的到达广告请求数量
	AdShowCount          int64   `json:"ad_show_count"`           // 展示数
	AdClickCount         int64   `json:"ad_click_count"`          // 点击数
	AdRequestsMatchRate  float64 `json:"ad_requests_match_rate"`  //填充率
	AdRequestsShowRate   float64 `json:"ad_requests_show_rate"`   // 请求展示率',
	AdClickThroughRate   float64 `json:"ad_click_through_rate"`   // 点击率',
	Earnings             float64 `json:"earnings"`                // 收入',
	ECpm                 float64 `json:"ecpm" gorm:"column:ecpm"` // ECPM',
	ARPU                 float64 `json:"arpu"`
}

// AdsReport 最终综合报表页展示的数据字段「需要和 AdsColumns 的 key 一致」
type AdsReport struct {
	StatDay           string  `json:"stat_day"`
	Country           string  `json:"country"`
	AccountId         int64   `json:"account_id"`
	AppId             string  `json:"app_id"`
	AppName           string  `json:"app_name"`
	AccountName       string  `json:"account_name"`
	RegionCountry     string  `json:"region_country"`
	AdRequests        int64   `json:"ad_requests"`             // 到达服务器的请求数量
	MatchedAdRequests int64   `json:"matched_ad_requests"`     // 匹配到的到达广告请求数量
	ShowCount         int64   `json:"show_count"`              // 展示数
	ClickCount        int64   `json:"click_count"`             // 点击数
	RequestsMatchRate float64 `json:"requests_match_rate"`     //填充率
	RequestsShowRate  float64 `json:"requests_show_rate"`      // 请求展示率',
	ClickThroughRate  float64 `json:"click_through_rate"`      // 点击率',
	Earnings          float64 `json:"earnings"`                // 收入',
	ECpm              float64 `json:"ecpm" gorm:"column:ecpm"` // ECPM',
	ARPU              float64 `json:"arpu"`
}

var (
	// ComprehensiveColumns 综合报表展示字段
	ComprehensiveColumns = []*ReportColumn{
		{Key: "roi", Label: "ROI", Align: "right", Min: 100, Suffix: "%"},
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
		{Key: "click_through_rate", Label: "点击率", Align: "right", Min: 100, Suffix: "%"},
		{Key: "download_activate_rate", Label: "激活转化率", Align: "right", Min: 100, Suffix: "%"},
		{Key: "click_download_rate", Label: "下载转化率", Align: "right", Min: 100, Suffix: "%"},
		{Key: "ad_requests", Label: "请求(变现)", Align: "right", Min: 100},
		{Key: "matched_ad_requests", Label: "填充(变现)", Align: "right", Min: 100},
		{Key: "ad_requests_match_rate", Label: "填充率(变现)", Align: "right", Min: 100, Suffix: "%"},
		{Key: "ad_show_count", Label: "曝光量(变现)", Align: "right", Min: 100},
		{Key: "ad_click_count", Label: "点击量(变现)", Align: "right", Min: 100},
		{Key: "ad_requests_show_rate", Label: "展示率(变现)", Align: "right", Min: 100, Suffix: "%"},
		{Key: "ad_click_through_rate", Label: "点击率(变现)", Align: "right", Min: 100, Suffix: "%"},
		//{Key: "arpu", Label: "ARPU(变现:$)", Align: "right", Min: 100},
	}
	// AdsColumns 变现报表展示字段
	AdsColumns = []*ReportColumn{
		{Key: "earnings", Label: "收入($)", Align: "right", Min: 100},
		{Key: "ecpm", Label: "eCPM(变现:$)", Align: "right", Min: 100},
		{Key: "ad_requests", Label: "请求(变现)", Align: "right", Min: 100},
		{Key: "matched_ad_requests", Label: "填充(变现)", Align: "right", Min: 100},
		{Key: "requests_match_rate", Label: "填充率(变现)", Align: "right", Min: 100, Suffix: "%"},
		{Key: "show_count", Label: "曝光量(变现)", Align: "right", Min: 100},
		{Key: "click_count", Label: "点击量(变现)", Align: "right", Min: 100},
		{Key: "requests_show_rate", Label: "展示率(变现)", Align: "right", Min: 100, Suffix: "%"},
		{Key: "click_through_rate", Label: "点击率(变现)", Align: "right", Min: 100, Suffix: "%"},
		//{Key: "arpu", Label: "ARPU(变现:$)", Align: "right", Min: 100},
	}
	AccountColumn = &ReportColumn{Key: "account_name", Label: "账户", Align: "left", Min: 120, Fix: true, Show: true}
	AppColumn     = &ReportColumn{Key: "app_name", Label: "应用", Align: "left", Min: 130, Fix: true, Show: true}
	CountryColumn = &ReportColumn{Key: "region_country", Label: "区域&国家", Align: "left", Min: 130, Fix: true, Show: true}

	// ComprehensiveMarketSQLColumnsMap 综合报表投放查询汇总字段「as 需要和数据库模型字段一直」
	ComprehensiveMarketSQLColumnsMap = []string{
		"round(sum(`cost`), 3) as `cost`",
		"sum(`show_count`) as `show_count`",
		"sum(`click_count`) as `click_count`",
		"sum(`install_count`) as `install_count`",
		"sum(`download_count`) as `download_count`",
		"sum(`activate_count`) as `activate_count`",
		"sum(`retain_count`) as `retain_count`",
		"sum(`three_retain_count`) as `three_retain_count`",
		"sum(`seven_retain_count`) as `seven_retain_count`",
	}
	// AdsSQLColumnsMap 综合报表变现查询汇总字段「as 需要和数据库模型字段一直」
	AdsSQLColumnsMap = []string{
		"round(sum(`earnings`), 3) as `earnings`",
		"sum(`ad_requests`) as `ad_requests`",
		"sum(`matched_ad_requests`) as `matched_ad_requests`",
		"sum(`show_count`) as `show_count`",
		"sum(`click_count`) as `click_count`",
	}
)

func formatCountries(countries [][]string) []string {
	rs := make([]string, 0)
	if len(countries) == 0 {
		return rs
	}
	for _, country := range countries {
		if len(country) != 2 {
			continue
		}
		rs = append(rs, country[1])
	}
	return rs
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

func adsColumns(dimensions []string) []string {
	rs := append(AdsSQLColumnsMap, "stat_day")
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
	return rs
}

func accountMap(accountType int) map[int64]string {
	rs := make(map[int64]string)
	accounts, err := model.NewAct(vars.DBMysql).AllAccounts()
	if err != nil {
		return rs
	}
	for _, account := range accounts {
		if account.AccountType == int64(accountType) {
			rs[account.Id] = account.AccountName
		}
	}
	return rs
}

func regionCountryMap() map[string]string {
	rs := make(map[string]string)
	areas, err := model.NewOverseasArea(vars.DBMysql).AreaCountries()
	if err != nil {
		return rs
	}
	for _, area := range areas {
		rs[area.CCode] = area.CName
	}
	return rs
}
