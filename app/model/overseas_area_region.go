package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
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

func (m *OverseasAreaRegion) AreaSet(v *OverseasAreaRegion) (err error) {
	err = m.Table(m.TableName()).Create(&v).Error
	return
}

func (m *OverseasAreaRegion) FindCCodesByAreaIds(areaIds []int64) (codes []string, err error) {
	err = m.Table(m.TableName()).Where("area_id in ?", areaIds).Select("c_code").Find(&codes).Error
	return
}

func (m *OverseasAreaRegion) FindByAreaIds(areaIds []int64) (as []*OverseasAreaRegion, err error) {
	query := m.Table(m.TableName())
	if len(areaIds) > 0 {
		query = query.Where("area_id in ?", areaIds)
	}
	err = query.Find(&as).Error
	return
}

// AreaColumnParse 选择区域维度时字段查询解析
func (m *OverseasAreaRegion) AreaColumnParse(areaIds []int64) (c string) {
	areas, err := m.FindByAreaIds(areaIds)
	if err != nil {
		return
	}
	columns := make(map[int64][]string)
	for _, area := range areas {
		columns[area.AreaId] = append(columns[area.AreaId], area.CCode)
	}
	if len(columns) == 0 {
		return
	}
	v := []string{"(CASE"}
	for areaId, countries := range columns {
		if len(countries) > 0 {
			v = append(v, fmt.Sprintf(
				"WHEN `country` IN ('%s') THEN %d", strings.Join(countries, "','"),
				areaId,
			))
		}
	}
	v = append(v, "ELSE 0 END) as area_id")
	return strings.Join(v, " ")
}
