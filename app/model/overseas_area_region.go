package model

import (
	"gorm.io/gorm"
)

type OverseasAreaRegion struct {
	connectDb
	AreaId int64  `json:"area_id"`
	CCode  string `json:"c_code"`
}

func NewOverseasAreaRegion(db *gorm.DB) *OverseasAreaRegion {
	return &OverseasAreaRegion{connectDb: connectDb{DB: db}}
}

func (m *OverseasAreaRegion) TableName() string {
	return "overseas_area_regions"
}

func (m *OverseasAreaRegion) FindCCodesByAreaId(areaId int64) (codes []string, err error) {
	err = m.Table(m.TableName()).Where("area_id = ?", areaId).Select("c_code").Find(&codes).Error
	return
}

func (m *OverseasAreaRegion) AreaSet(v *OverseasAreaRegion) (err error) {
	err = m.Table(m.TableName()).Create(&v).Error
	return
}
