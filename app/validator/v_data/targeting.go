package v_data

type VTargetingList struct {
	CampaignId string `form:"campaign_id" binding:"required"`
	AccountId  int64  `form:"account_id" binding:"required"`
}

type VTargetingCreate struct {
	CampaignId         string     `json:"campaign_id" binding:"required"`
	TargetingType      string     `json:"targeting_type"`
	TargetingName      string     `json:"targeting_name" binding:"required"`
	Gender             string     `json:"gender"`
	Age                []string   `json:"age"`
	NetworkType        []string   `json:"network_type"`
	Location           string     `json:"location"`
	LocationType       string     `json:"location_type"`
	IncludeLocation    []string   `json:"include_location"`
	ExcludeLocation    []string   `json:"exclude_location"`
	Carrier            string     `json:"carrier"`
	Carriers           [][]string `json:"carriers"`
	AppCategory        string     `json:"app_category"`
	AppCategories      []string   `json:"app_categories"`
	AppInterest        string     `json:"app_interest"`
	AppInterests       [][]string `json:"app_interests"`
	Audience           string     `json:"audience"`
	Audiences          []string   `json:"audiences"`
	NotAudience        []string   `json:"not_audience"`
	SeriesType         string     `json:"series_type"`
	Series             []string   `json:"series"`
	MediaAppCategory   string     `json:"media_app_category"`
	AppCategoryOfMedia [][]string `json:"app_category_of_media"`
	LanguageCheck      string     `json:"language_check"`
	Language           []string   `json:"language"`
	InstalledApps      string     `json:"installed_apps"`
}
