package logic

import "bs.mobgi.cc/cronJobs/job_data/statements"

var (
	maxPageSize    int64 = 50
	maxGoroutine   int64 = 10
	failRetryTimes uint8 = 3 // 单个 API 失败重试最大次数
	agApp                = "AG_APP_FOR_DISPLAY_NETWORK"
)

type queryParam struct {
	accountId   int64
	accessToken string
	failed      uint8
}

type run struct {
	*queryParam
	page int64
}

type AppRequest struct {
	Page      int64        `json:"page"`
	PageSize  int64        `json:"page_size"`
	Filtering ReqFiltering `json:"filtering"`
}

type ReqFiltering struct {
	ProductType string `json:"product_type"`
	AgAppType   string `json:"ag_app_type"`
}
type ProductInfo struct {
	ProductType string  `json:"product_type"`
	ProductId   string  `json:"product_id"`
	ProductInfo AppInfo `json:"product_info"`
}
type AppInfo struct {
	App App `json:"app"`
}
type App struct {
	IconUrl     string `json:"icon_url"`
	PackageName string `json:"package_name"`
	Description string `json:"description"`
	AppId       string `json:"app_id"`
	ProductName string `json:"product_name"`
}
type AppResponse struct {
	statements.BaseResponse
	Data AdsAppInfo `json:"data"`
}
type AdsAppInfo struct {
	Total int64          `json:"total"`
	Data  []*ProductInfo `json:"data"`
}
