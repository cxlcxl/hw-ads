package v_data

type Pagination struct {
	Page     int64 `json:"page" form:"page" binding:"required"`
	PageSize int64 `json:"page_size" form:"page_size" binding:"required"`
}
