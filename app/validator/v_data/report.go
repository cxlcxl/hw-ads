package v_data

import "bs.mobgi.cc/app/vars"

type VReportComprehensive struct {
	DateRange   []string   `json:"date_range" binding:"required"`
	Dimensions  []string   `json:"dimensions"`
	AccountIds  []int64    `json:"account_ids"`
	AppIds      []string   `json:"app_ids"`
	Countries   [][]string `json:"countries"`
	ShowColumns []string   `json:"show_columns" binding:"required"`
	Order       string     `json:"order"`
	By          string     `json:"by"`
	Pagination
	User *vars.LoginUser
}

type VReportAds struct {
	DateRange   []string   `json:"date_range" binding:"required"`
	Dimensions  []string   `json:"dimensions"`
	AccountIds  []int64    `json:"account_ids"`
	AppIds      []string   `json:"app_ids"`
	Countries   [][]string `json:"countries"`
	ShowColumns []string   `json:"show_columns" binding:"required"`
	Pagination
	User *vars.LoginUser
}

type VReportColumn struct {
	Columns []string `json:"columns" binding:"required"`
	Module  string   `json:"module" binding:"required"`
	User    *vars.LoginUser
}
