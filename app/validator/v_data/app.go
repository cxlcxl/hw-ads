package v_data

import "bs.mobgi.cc/app/vars"

type VAppCreate struct {
	AppName string `json:"app_name" binding:"required"`
	AppId   string `json:"app_id" binding:"required"`
	PkgName string `json:"pkg_name" binding:"required"`
	Channel int64  `json:"channel"`
	AppType string `json:"app_type"`
	Tags    string `json:"tags"`
}

type VAppUpdate struct {
	Id      int64  `json:"id" binding:"required"`
	AppName string `json:"app_name" binding:"required"`
	PkgName string `json:"pkg_name" binding:"required"`
	Channel int64  `json:"channel"`
	AppType string `json:"app_type"`
	Tags    string `json:"tags"`
}

type VAppList struct {
	Pagination
	AppId      string  `json:"app_id,optional"`
	AppName    string  `json:"app_name,optional"`
	AppType    string  `json:"app_type,optional"`
	Channel    int64   `json:"channel,optional"`
	AccountIds []int64 `json:"account_ids"`

	User *vars.LoginUser
}

type VAppCampaignList struct {
	AppName  string `form:"app_name,optional"`
	Page     int64  `form:"page" binding:"required"`
	PageSize int64  `form:"page_size" binding:"required"`
}

type CampaignAppInfo struct {
	AppName      string `json:"app_name"`
	AppId        string `json:"app_id"`
	AdvertiserId string `json:"advertiser_id"`
	IconUrl      string `json:"icon_url"`
	AccountId    int64  `json:"account_id" binding:"required"`
}

type VAppPull struct {
	AccountId    int64  `json:"account_id" binding:"required"`
	AdvertiserId string `json:"advertiser_id" binding:"required"`
}
