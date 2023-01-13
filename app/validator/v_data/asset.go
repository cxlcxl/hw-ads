package v_data

type VAssetList struct {
	Page      int64  `form:"page" binding:"required"`
	PageSize  int64  `form:"page_size" binding:"required"`
	AssetType string `form:"asset_type,optional"`
	Width     int64  `form:"width,optional"`
	Height    int64  `form:"height,optional"`
	AppId     string `form:"app_id,optional"`
}

type VAssetSync struct {
	AppId string `form:"app_id" binding:"required"`
}
