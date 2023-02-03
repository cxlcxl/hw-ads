package model

import (
	"fmt"
	"gorm.io/gorm"
)

// OverseasRegion
// 海外区域数据来源地址：https://developer.huawei.com/consumer/cn/doc/distribution/promotion/marketing-api-tool-targeting7-0000001286343134#ZH-CN_TOPIC_0000001286343134__li1079145116117
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

func (m *OverseasRegion) RegionCreate(region *OverseasRegion) (err error) {
	err = m.Table(m.TableName()).Create(&region).Error
	return
}

type Country struct {
	CCode    string `json:"c_code"`
	CName    string `json:"c_name"`
	AreaName string `json:"area_name"`
}

func (m *OverseasRegion) FindCountries(areaId int64, k string, offset, limit int) (regions []*Country, total int64, err error) {
	columns := "select t0.c_code,t0.c_name,t2.name as area_name "
	count := "select count(1) as total "
	sql := fmt.Sprintf(
		"from %s t0 "+
			"left join %s t1 on t0.c_code = t1.c_code "+
			"left join %s t2 on t1.area_id = t2.id "+
			"where t0.level = 0 ",
		m.TableName(), NewOverseasAreaRegion(nil).TableName(), NewOverseasArea(nil).TableName(),
	)
	val := make([]interface{}, 0)
	if areaId > 0 {
		sql += " and t1.area_id = ?"
		val = append(val, areaId)
	} else if areaId == -2 {
		// 筛选未分地区的国家
		sql += " and t1.area_id is null"
	}
	if k != "" {
		sql += "c_code = ? or c_name like ?"
		val = append(val, k, "%"+k+"%")
	}
	if err = m.Raw(count+sql+"", val...).Scan(&total).Error; err != nil {
		return
	}
	val = append(val, limit, offset)
	err = m.Raw(columns+sql+" order by t0.c_code asc,t0.id asc limit ? offset ?", val...).Find(&regions).Error
	return
}
