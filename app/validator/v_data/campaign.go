package v_data

type VCampaignList struct {
	Page         int64  `form:"page" binding:"required"`
	PageSize     int64  `form:"page_size" binding:"required"`
	CampaignName string `form:"campaign_name,optional"`
	CampaignId   string `form:"campaign_id,optional"`
	CampaignType string `form:"show_status,optional"`
	AppId        string `form:"app_id,optional"`
}
