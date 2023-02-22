package v_data

type VSettingsCronUpdate struct {
	Id          int64
	JobSchedule string `json:"job_schedule" binding:"required"`
	OrderBy     uint8  `json:"order_by" binding:"required,numeric"`
	PauseRule   int    `json:"pause_rule"`
	Version     int    `json:"version" binding:"required,numeric"`
	StatDay     string `json:"stat_day" binding:"required"`
	Remark      string `json:"remark"`
}

type VSettingsCronSchedule struct {
	ApiModule string `json:"api_module" binding:"required"`
	PauseRule int64  `json:"pause_rule"`
	StatDay   string `json:"stat_day" binding:"required"`
}

type VSettingsLog struct {
	D string `json:"d" binding:"required"`
}

type VSettingsConfigs struct {
	Key   string `form:"_k"`
	Desc  string `form:"_desc"`
	State uint8  `form:"state"`
	Pagination
}

type VSettingsConfigCreate struct {
	Key    string `json:"key" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Val    string `json:"val" binding:"required"`
	Bak1   string `json:"bak1"`
	Bak2   string `json:"bak2"`
	Remark string `json:"remark"`
}

type VSettingsConfigUpdate struct {
	Id    int64 `json:"id" binding:"required"`
	State uint8 `json:"state" binding:"numeric"`
	VSettingsConfigCreate
}
