package v_data

import "bs.mobgi.cc/app/vars"

type VAccountList struct {
	AccountId   int64  `form:"account_id,optional"`
	AccountName string `form:"account_name,optional"`
	AccountType int64  `form:"account_type,optional"`
	State       int64  `form:"state"`
	Page        int64  `form:"page"`
	PageSize    int64  `form:"page_size"`

	User *vars.LoginUser
}

type VAccountParents struct {
	ParentId    int64  `form:"parent_id,optional"`
	AccountName string `form:"account_name,optional"`
}

type VAccountSearch struct {
	AccountName string `form:"account_name,optional"`
}

type VAccountAuth struct {
	AuthorizationCode string `json:"authorization_code" binding:"required"`
	State             string `json:"state" binding:"required"`
}

type VAccountCreate struct {
	ParentId     int64  `json:"parent_id"`
	AccountName  string `json:"account_name" binding:"required"`
	AccountType  uint8  `json:"account_type" binding:"required"`
	AdvertiserId string `json:"advertiser_id"`
	DeveloperId  string `json:"developer_id"`
	ClientId     string `json:"client_id"`
	Secret       string `json:"secret"`
	State        int64  `json:"state"`

	User *vars.LoginUser
}

type VAccountUpdate struct {
	Id int64 `json:"id" binding:"required"`
	VAccountCreate
}
