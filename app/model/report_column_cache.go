package model

import (
	"bs.mobgi.cc/app/cache"
	"gorm.io/gorm"
)

type ReportColumn struct {
	connectDb

	Id       int64  `json:"id"`
	Columns  string `json:"columns"`
	CacheKey string `json:"cache_key"`
}

var (
	columnKey = "report:column:"
)

// NewReportColumn 实例
func NewReportColumn(db *gorm.DB) *ReportColumn {
	return &ReportColumn{connectDb: connectDb{DB: db}}
}

func (m *ReportColumn) TableName() string {
	return "report_column_caches"
}

func (m *ReportColumn) GetColumnCache(k string) (c *ReportColumn, err error) {
	err = cache.New(m.DB).QueryRow(columnKey, &c, k, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Where("cache_key = ?", id).First(v).Error
	})
	return
}

func (m *ReportColumn) UpdateColumnCache(k string, val map[string]interface{}) error {
	return cache.New(m.DB).SetRow(columnKey, val, k, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Where("cache_key = ?", id).Updates(v).Error
	})
}
