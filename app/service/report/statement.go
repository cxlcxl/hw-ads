package servicereport

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
	"errors"
)

func ValidatorDimension(dimensions []string) error {
	if utils.InArray(vars.ReportDimensionArea, dimensions) && utils.InArray(vars.ReportDimensionCountry, dimensions) {
		return errors.New("国家和地区维度不能同时筛选")
	}
	return nil
}

type ReportColumn struct {
	Key    string `json:"key"`
	Label  string `json:"label"`
	Align  string `json:"align"`
	Show   bool   `json:"show"`
	Min    uint8  `json:"min"`
	Fix    bool   `json:"fix"`
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
	Sort   string `json:"sort"`
}

// ComprehensiveReport 最终综合报表页展示的数据字段「需要和 ComprehensiveColumns 的 key 一致」
type ComprehensiveReport struct {
	StatDay              string  `json:"stat_day"`
	Country              string  `json:"country"`
	AccountId            int64   `json:"account_id"`
	AppId                string  `json:"app_id"`
	AppName              string  `json:"app_name"`
	AccountName          string  `json:"account_name"`
	AreaName             string  `json:"area_name"`
	CountryName          string  `json:"country_name"`
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
	AreaName          string  `json:"area_name"`
	CountryName       string  `json:"country_name"`
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
		{Key: "roi", Label: "ROI", Align: "right", Min: 100, Suffix: "%", Sort: "1"},
		{Key: "cost", Label: "花费($)", Align: "right", Min: 100, Sort: "custom"},
		{Key: "earnings", Label: "收入($)", Align: "right", Min: 100, Sort: "custom"},
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
	DateColumn    = &ReportColumn{Key: "stat_day", Label: "日期", Align: "center", Min: 90, Fix: true, Show: true}
	AccountColumn = &ReportColumn{Key: "account_name", Label: "账户", Align: "left", Min: 120, Fix: true, Show: true}
	AppColumn     = &ReportColumn{Key: "app_name", Label: "应用", Align: "left", Min: 130, Fix: true, Show: true}
	AreaColumn    = &ReportColumn{Key: "area_name", Label: "地区", Align: "left", Min: 90, Fix: true, Show: true}
	CountryColumn = &ReportColumn{Key: "country_name", Label: "国家", Align: "left", Min: 120, Fix: true, Show: true}

	// MarketSQLColumns 综合报表投放查询汇总字段「as 需要和数据库模型字段一直」
	MarketSQLColumns = []string{
		"`stat_day`",
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
	// AdsSQLColumns 综合报表变现查询汇总字段「as 需要和数据库模型字段一直」
	AdsSQLColumns = []string{
		"`stat_day`",
		"round(sum(`earnings`), 3) as `earnings`",
		"sum(`ad_requests`) as `ad_requests`",
		"sum(`matched_ad_requests`) as `matched_ad_requests`",
		"sum(`show_count`) as `ad_show_count`",
		"sum(`click_count`) as `ad_click_count`",
	}
	ComprehensiveGranularityAll = []string{
		"NOW() as `stat_day`",
		"round(sum(t0.`cost`), 3) as `cost`",
		"sum(t0.`show_count`) as `show_count`",
		"sum(t0.`click_count`) as `click_count`",
		"sum(t0.`install_count`) as `install_count`",
		"sum(t0.`download_count`) as `download_count`",
		"sum(t0.`activate_count`) as `activate_count`",
		"sum(t0.`retain_count`) as `retain_count`",
		"sum(t0.`three_retain_count`) as `three_retain_count`",
		"sum(t0.`seven_retain_count`) as `seven_retain_count`",
		"round(sum(t1.`earnings`), 3) as `earnings`",
		"sum(t1.`ad_requests`) as `ad_requests`",
		"sum(t1.`matched_ad_requests`) as `matched_ad_requests`",
		"sum(t1.`ad_show_count`) as `ad_show_count`",
		"sum(t1.`ad_click_count`) as `ad_click_count`",
	}
	ComprehensiveGranularityDate = []string{
		"t0.*",
		"t1.`earnings`",
		"t1.`ad_requests`",
		"t1.`matched_ad_requests`",
		"t1.`ad_show_count`",
		"t1.`ad_click_count`",
	}
	// AdsReportColumns 综合报表变现查询汇总字段「as 需要和数据库模型字段一直」
	AdsReportColumns = []string{
		"`stat_day`",
		"round(sum(`earnings`), 3) as `earnings`",
		"sum(`ad_requests`) as `ad_requests`",
		"sum(`matched_ad_requests`) as `matched_ad_requests`",
		"sum(`show_count`) as `show_count`",
		"sum(`click_count`) as `click_count`",
	}

	sortableColumns = map[string]string{
		"cost":     "t0.`cost`",
		"earnings": "t1.`earnings`",
	}
	orderBy = map[string]string{"descending": "desc", "ascending": "asc"}
)

func formatCountries(countries [][]string, dimensions []string) []string {
	rs := make([]string, 0)
	if !utils.InArray(vars.ReportDimensionCountry, dimensions) || len(countries) == 0 {
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

func accountMap(accountType int) map[int64]string {
	rs := make(map[int64]string)
	accounts, err := model.NewAct(vars.DBMysql).AllAccounts(nil)
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

func regionCountryMap() map[string]*model.AreaCountry {
	rs := make(map[string]*model.AreaCountry)
	areas, err := model.NewOverseasArea(vars.DBMysql).AreaCountries()
	if err != nil {
		return rs
	}
	for _, area := range areas {
		rs[area.CCode] = area
	}
	return rs
}

func areaMap() map[int64]string {
	rs := make(map[int64]string)
	areas, err := model.NewOverseasArea(vars.DBMysql).Areas()
	if err != nil {
		return rs
	}
	for _, area := range areas {
		rs[area.Id] = area.Name
	}
	return rs
}

func queryCountriesByAreaIds(areas []int64) []string {
	cCodes, err := model.NewOverseasAreaRegion(vars.DBMysql).FindCCodesByAreaIds(areas)
	if err != nil {
		return nil
	}
	return cCodes
}
