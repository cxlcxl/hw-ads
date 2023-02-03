package v_data

type VCountries struct {
	AreaId int64  `form:"area_id" binding:"required,numeric"`
	K      string `form:"k"`
	Pagination
}

type VRegionCreate struct {
	T        string `json:"t" binding:"required"`
	AreaId   int64  `json:"area_id"`
	AreaName string `json:"area_name"`
	CCode    string `json:"c_code"`
	CName    string `json:"c_name"`
}

type VRegionAreaSet struct {
	AreaId int64  `json:"area_id"`
	CCode  string `json:"c_code"`
}
