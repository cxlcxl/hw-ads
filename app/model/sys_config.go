package model

import (
	"bs.mobgi.cc/app/cache"
	"gorm.io/gorm"
)

type SysConfig struct {
	connectDb

	Id    int64  `json:"id"`
	Key   string `json:"key" gorm:"column:_k"`
	Val   string `json:"val" gorm:"column:_v"`
	Desc  string `json:"desc" gorm:"column:_desc"`
	State uint8  `json:"state"`
	Bak1  string `json:"bak1"`
	Bak2  string `json:"bak2"`
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
