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
