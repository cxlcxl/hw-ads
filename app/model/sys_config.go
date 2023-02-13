package model

import (
	"bs.mobgi.cc/app/cache"
	"gorm.io/gorm"
)

type SysConfig struct {
	connectDb

	Id     int64  `json:"id"`
	Key    string `json:"key" gorm:"column:_k"`
	Val    string `json:"val" gorm:"column:_v"`
	Desc   string `json:"desc" gorm:"column:_desc"`
	State  uint8  `json:"state"`
	Bak1   string `json:"bak1"`
	Bak2   string `json:"bak2"`
	Remark string `json:"remark"`
}

var (
	configKey = "db:configs:"
)

// NewSysConfig 实例
func NewSysConfig(db *gorm.DB) *SysConfig {
	return &SysConfig{connectDb: connectDb{DB: db}}
}

func (m *SysConfig) TableName() string {
	return "sys_configs"
}

func (m *SysConfig) List(k, desc string, state uint8, offset, limit int64) (configs []*SysConfig, total int64, err error) {
	query := m.Table(m.TableName()).Where("state = ?", state)
	if k != "" {
		query = query.Where("_k like ?", "%"+k+"%")
	}
	if desc != "" {
		query = query.Where("_desc like ?", "%"+desc+"%")
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	err = query.Offset(int(offset)).Limit(int(limit)).Find(&configs).Error
	return
}

func (m *SysConfig) FindConfigsByKey(k string) (configs []*SysConfig, err error) {
	err = cache.New(m.DB).QueryRow(configKey, &configs, k, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Where("_k = ? and state = 1", id).First(v).Error
	})
	return
}

func (m *SysConfig) FindOneByKey(k string) (config *SysConfig, err error) {
	err = m.Table(m.TableName()).Where("_k = ? and state = 1", k).First(&config).Error
	return
}

func (m *SysConfig) FindOneById(id int64) (config *SysConfig, err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).First(&config).Error
	return
}

func (m *SysConfig) CreateConfig(c SysConfig) (err error) {
	err = m.Table(m.TableName()).Create(&c).Error
	return
}

func (m *SysConfig) UpdateConfig(id int64, v map[string]interface{}) (err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).Updates(v).Error
	return
}
