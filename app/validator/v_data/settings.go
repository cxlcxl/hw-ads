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
