package model

import (
	"bs.mobgi.cc/app/cache"
	"errors"
	"gorm.io/gorm"
)

type ReportColumn struct {
	connectDb

	Columns   string `json:"columns"`
	ColumnKey string `json:"column_key"`
}

var (
	columnKey = "report:column:"
)

// NewReportColumn 实例
func NewReportColumn(db *gorm.DB) *ReportColumn {
	return &ReportColumn{connectDb: connectDb{DB: db}}
}

func (m *ReportColumn) TableName() string {
	return "report_columns"
}

func (m *ReportColumn) GetColumn(k string) (c *ReportColumn, err error) {
	err = cache.New(m.DB).QueryRow(columnKey, &c, k, func(db *gorm.DB, v interface{}, id interface{}) error {
		if err = db.Table(m.TableName()).Where("column_key = ?", id).First(v).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return db.Table(m.TableName()).Create(&ReportColumn{ColumnKey: id.(string)}).Error
			} else {
				return err
			}
		}
		return nil
	})
	return
}

func (m *ReportColumn) UpdateColumn(k string, val map[string]interface{}) error {
	return cache.New(m.DB).SetRow(columnKey, val, k, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Where("column_key = ?", id).Updates(v).Error
	})
}
