package logic

import "bs.mobgi.cc/cronJobs/job_data/statements"

var (
	maxGoroutine   int64 = 10
	failRetryTimes uint8 = 5 // 单个 API 失败重试最大次数
)

type app struct {
	appId   string
	appName string
}

type run struct {
	*queryParam
	page int64
}

type queryParam struct {
	accountId   int64
	accessToken string
	failed      uint8 // 失败次数
}

const (
	FilterTypeCreative = "CREATIVE"
	FilterTypeAdgroup  = "ADGROUP"

	StateTimeHour    = "STAT_TIME_GRANULARITY_HOURLY"
	StateTimeDaily   = "STAT_TIME_GRANULARITY_DAILY"
	StateTimeMonth   = "STAT_TIME_GRANULARITY_MONTHLY"
	StateTimeSummary = "STAT_TIME_GRANULARITY_SUMMARY"

	OrderAsc  = "ASC"
	OrderDesc = "DESC"
)

type CountryRequest struct {
	TimeGranularity string    `json:"time_granularity"`
	StartDate       string    `json:"start_date"`
	EndDate         string    `json:"end_date"`
	IsAbroad        bool      `json:"is_abroad"`
	Page            int64     `json:"page"`
	PageSize        int64     `json:"page_size"`
	OrderType       string    `json:"order_type"`
	Filtering       Filtering `json:"filtering"`
}

type Filtering struct {
	OtherFilterType string `json:"other_filter_type"`
}
type CountryResponse struct {
	statements.BaseResponse
	Data struct {
		List     []*Country       `json:"list"`
		PageInfo statements.Pages `json:"page_info"`
	} `json:"data"`
}

type CountryPagesResponse struct {
	statements.BaseResponse
	Data struct {
		PageInfo statements.Pages `json:"page_info"`
	} `json:"data"`
}

type Country struct {
	StatDatetime          string `json:"stat_datetime"`
	AdvertiserId          string `json:"advertiser_id"`
	CampaignId            string `json:"campaign_id"`
	CampaignName          string `json:"campaign_name"`
	AdgroupId             string `json:"adgroup_id"`
	AdgroupName           string `json:"adgroup_name"`
	CreativeId            string `json:"creative_id"`
	CreativeName          string `json:"creative_name"`
	Country               string `json:"country"`
	PackageName           string `json:"package_name"`
	ShowCount             int64  `json:"show_count"`
	ClickCount            int64  `json:"click_count"`
	Cpc                   string `json:"cpc"`
	Cpm                   string `json:"thousand_show_cost"`
	Cost                  string `json:"cost"`
	Cpd                   string `json:"download_cost"`
	Cpi                   string `json:"install_cost"`
	Cpa                   string `json:"active_cost_normalized"`
	DownloadCount         int64  `json:"download_count"`
	InstallCount          int64  `json:"install_count"`
	ActiveCountNormalized int64  `json:"active_count_normalized"`
	RetainCountNormalized int64  `json:"retain_count_normalized"`
	RetainCostNormalized  string `json:"retain_cost_normalized"`
	//SevenDayRetainCount   int64  `json:"seven_day_retain_count"`
	//ThreeDayRetainCount   int64  `json:"three_day_retain_count"`
	//ThreeDayRetainCost    string `json:"three_day_retain_cost"`
	//SevenDayRetainCost    string `json:"seven_day_retain_cost"`
	//AppCustomCount        int64  `json:"app_custom_count"`
	//AppCustomCost         string `json:"app_custom_cost"`
	//WebCustomCount        int64  `json:"web_custom_count"`
	//WebCustomCost         string `json:"web_custom_cost"`
	//PlayCount             int64  `json:"play_count"`
	//PlayOverCount         int64  `json:"play_over_count"`
	//BrowseCount           int64  `json:"browse_count"`
	//BrowseCost            string `json:"browse_cost"`
	//ShareCount            int64  `json:"share_count"`
	//ShareCost             string `json:"share_cost"`
	//PayCountNormalized    int64  `json:"pay_count_normalized"`
	//PayCostNormalized     string `json:"pay_cost_normalized"`
	//Iaa                   string `json:"attribution_income_iaa"`
	//ThirtyRoi             string `json:"ad_income_thirty_day_roi"`
	//FifteenRoi            string `json:"ad_income_fifteen_day_roi"`
	//LtvHms                string `json:"ad_income_seven_day_ltv_hms"`
}
