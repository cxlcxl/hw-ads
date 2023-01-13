package model

import (
	"gorm.io/gorm"
)

type OverseasRegion struct {
	connectDb

	Id          int64  `json:"id"`
	CId         string `json:"c_id"`
	Pid         string `json:"pid"`
	CCode       string `json:"c_code"`
	CName       string `json:"c_name"`
	Level       int64  `json:"level"`
	ContinentId int64  `json:"continent_id"`
}

func NewOverseasRegion(db *gorm.DB) *OverseasRegion {
	return &OverseasRegion{connectDb: connectDb{DB: db}}
}

func (m *OverseasRegion) TableName() string {
	return "overseas_regions"
}

func (m *OverseasRegion) GetCountries() (regions []*OverseasRegion, err error) {
	err = m.Table(m.TableName()).Where("level = 0").Order("c_name asc").Find(&regions).Error
	return
}
