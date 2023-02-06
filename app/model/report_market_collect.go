package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ReportMarketCollect struct {
	connectDb

	Id               int64     `json:"id"`
	StatDay          time.Time `json:"stat_day"`           // 日: 日粒度，例如2021-09-08
	Country          string    `json:"country"`            // 国家代码，使用华为开发者文档中的广告代码库
	AccountId        int64     `json:"account_id"`         // 应用所属账户ID
	AppId            string    `json:"app_id"`             // 应用ID（此处一般标识三方应用ID）
	AppName          string    `json:"app_name"`           // 应用名称
	Cost             float64   `json:"cost"`               // 花费
	ShowCount        int64     `json:"show_count"`         // 展示数
	ClickCount       int64     `json:"click_count"`        // 点击数
	DownloadCount    int64     `json:"download_count"`     // 下载数
	InstallCount     int64     `json:"install_count"`      // 安装数
	ActivateCount    int64     `json:"activate_count"`     // 激活数
	RetainCount      int64     `json:"retain_count"`       // 留存数
	ThreeRetainCount int64     `json:"three_retain_count"` // 三日留存数
	SevenRetainCount int64     `json:"seven_retain_count"` // 七日留存数
}

// NewRMC ReportMarketCollect 实例
func NewRMC(db *gorm.DB) *ReportMarketCollect {
	return &ReportMarketCollect{connectDb: connectDb{DB: db}}
}

func (m *ReportMarketCollect) TableName() string {
	return "report_market_collects"
}

func (m *ReportMarketCollect) BatchInsert(ms []*ReportMarketCollect) (err error) {
	if len(ms) == 0 {
		return nil
	}
	updateColumns := []string{
		"cost", "show_count", "click_count", "download_count", "install_count", "activate_count", "retain_count",
	}

	err = m.Table(m.TableName()).Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(updateColumns),
	}).CreateInBatches(ms, 300).Error
	return err
}

// AnalysisComprehensive 综合报表投放数据部分
func (m *ReportMarketCollect) AnalysisComprehensive(
	actIds []int64, appIds, dates, countries, selects, groups []string, offset, size int,
) (markets []*ReportMarketCollect, total int64, err error) {
	query := m.Table(m.TableName()).
		Select(selects).
		Where("stat_day between ? and ?", dates[0], dates[1]).
		Group("stat_day").Order("stat_day desc")
	if len(appIds) > 0 {
		query = query.Where("app_id in ?", appIds)
	}
	if len(actIds) > 0 {
		query = query.Where("account_id in ?", actIds)
	}
	if len(countries) > 0 {
		query = query.Where("country in ?", countries)
	}

	for _, group := range groups {
		query = query.Group(group).Order(group + " asc")
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	err = query.Offset(offset).Limit(size).Find(&markets).Error
	return
}
