package v_data

type VReportComprehensive struct {
	DateRange   []string `json:"date_range" binding:"required"`
	Dimensions  []string `json:"dimensions"`
	AccountIds  []int64  `json:"account_ids"`
	AppIds      []string `json:"app_ids"`
	ShowColumns []string `json:"show_columns" binding:"required"`
	Pagination
}
