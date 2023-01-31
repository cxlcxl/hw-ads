package v_data

type VAssetList struct {
	AssetType string `form:"asset_type,optional"`
	Width     int64  `form:"width,optional"`
	Height    int64  `form:"height,optional"`
	AppId     string `form:"app_id,optional"`

	Pagination
}

type VAssetSync struct {
	AppId string `form:"app_id" binding:"required"`
}
