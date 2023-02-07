package logic

import "bs.mobgi.cc/cronJobs/job_data/statements"

var (
	maxGoroutine   int64 = 10
	failRetryTimes uint8 = 3 // 单个 API 失败重试最大次数
	RequestGroupBy       = []string{"STAT_BREAK_DOWNS_COUNTRY", "STAT_BREAK_DOWNS_APP_ID", "STAT_BREAK_DOWNS_PLACEMENT_ID"}
)

type queryParam struct {
	accountId    int64
	clientId     string
	clientSecret string
	accessToken  string
	failed       uint8
}

type client struct {
	clientId string
	secret   string
}

type run struct {
	*queryParam
	page int64
}

const (
	StateTimeDaily = "STAT_TIME_GRANULARITY_DAILY"
	OrderAsc       = "ASC"
	CurrencyUsd    = "USD"
)

type AdsRequest struct {
	StartDate       string    `json:"start_date"`
	EndDate         string    `json:"end_date"`
	GroupBy         []string  `json:"group_by"`
	TimeGranularity string    `json:"time_granularity"`
	Page            int64     `json:"page"`
	PageSize        int64     `json:"page_size"`
	OrderType       string    `json:"order_type"`
	Filtering       Filtering `json:"filtering"`
}

type Filtering struct {
	Currency     string   `json:"currency"`
	AppIds       []string `json:"app_ids"`
	AdTypes      []string `json:"ad_types"`
	PlacementIds []string `json:"placement_ids"`
}

type Pages struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"page_size"`
	TotalNum  int64 `json:"total_number"`
	TotalPage int64 `json:"total_page"`
}

type AdsResponse struct {
	statements.BaseResponse
	Data struct {
		List []*AdsList `json:"list"`
	} `json:"data"`
}

type AdsPagesResponse struct {
	statements.BaseResponse
	Data struct {
		PageInfo Pages `json:"page_info"`
	} `json:"data"`
}

type AdsList struct {
	PlacementId              string  `json:"placement_id"`
	AppId                    string  `json:"app_id"`
	Country                  string  `json:"country"`
	AdType                   string  `json:"ad_type"`
	StatDatetime             string  `json:"stat_datetime"`
	Earnings                 float64 `json:"earnings"`
	ReachedAdRequests        int64   `json:"reached_ad_requests"`
	MatchedReachedAdRequests int64   `json:"matched_reached_ad_requests"`
	ShowCount                int64   `json:"show_count"`
	ClickCount               int64   `json:"click_count"`
	//AdRequestsMatchRate      float64 `json:"ad_requests_match_rate"`
	//AdRequestsShowRate       float64 `json:"ad_requests_show_rate"`
	//ClickThroughRate         float64 `json:"click_through_rate"`
}
