package model

import (
	"bs.mobgi.cc/app/cache"
	"fmt"
	"gorm.io/gorm"
)

type OverseasArea struct {
	connectDb
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

var (
	regionAreaKey    = "db:region:area"
	regionKey        = "db:regions"
	areaCountriesKey = "db:area_countries"
)

func NewOverseasArea(db *gorm.DB) *OverseasArea {
	return &OverseasArea{connectDb: connectDb{DB: db}}
}

func (m *OverseasArea) TableName() string {
	return "overseas_areas"
}

func (m *OverseasArea) Areas() (areas []*OverseasArea, err error) {
	err = cache.New(m.DB).Query(regionAreaKey, &areas, func(db *gorm.DB, v interface{}) error {
		return db.Table(m.TableName()).Find(v).Error
	})
	return
}

func (m *OverseasArea) AreaCreate(area *OverseasArea) (err error) {
	err = m.Table(m.TableName()).Create(&area).Error
	if err == nil {
		_ = cache.New(m.DB).DelQueryRowCache(regionAreaKey, "")
	}
	return
}

type Region struct {
	Id       int64     `json:"id"`
	CName    string    `json:"c_name"`
	Pid      int64     `json:"pid"`
	CCode    string    `json:"c_code"`
	Children []*Region `json:"children" gorm:"-"`
}

func (m *OverseasArea) GetRegions() (regions []*Region, err error) {
	err = cache.New(m.DB).Query(regionKey, &regions, func(db *gorm.DB, v interface{}) error {
		sql := fmt.Sprintf(
			"select id,name as c_name,0 as pid,id as c_code from %s "+
				"union all "+
				"select t1.id,t1.c_name,t0.area_id as pid,t0.c_code "+
				"from %s t0 "+
				"left join overseas_regions t1 on t0.c_code = t1.c_code "+
				"where t1.level = 0",
			m.TableName(), NewOverseasAreaRegion(nil).TableName(),
		)
		return db.Raw(sql).Find(v).Error
	})
	return
}

type AreaCountry struct {
	AreaName    string `json:"area_name"`
	CountryName string `json:"country_name"`
	CCode       string `json:"c_code"`
}

func (m *OverseasArea) AreaCountries() (areas []*AreaCountry, err error) {
	err = cache.New(m.DB).Query(areaCountriesKey, &areas, func(db *gorm.DB, v interface{}) error {
		sql := "select t0.c_code,IF(t2.name is null, ' - ', t2.name) as area_name,t0.c_name as country_name " +
			"from overseas_regions t0 " +
			"left join overseas_area_regions t1 on t0.c_code = t1.c_code " +
			"left join overseas_areas t2 on t1.area_id = t2.id " +
			"where t0.level = 0"
		return db.Raw(sql).Find(v).Error
	})
	return
}
