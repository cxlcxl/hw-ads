package v_data

type VTrackingList struct {
	AppId string `form:"app_id" binding:"required"`
}

type VTrackingRefresh struct {
	AppId string `form:"app_id" binding:"required"`
}

type VDictQuery struct {
	DictKey string `form:"dict_key" binding:"required"`
}
