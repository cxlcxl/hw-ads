package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ReportAdsCollect struct {
	connectDb

	Id                  int64     `json:"id"`
	StatDay             time.Time `json:"stat_day"`                // 日: 日粒度，例如2021-09-08
	Country             string    `json:"country"`                 // 国家代码，使用华为开发者文档中的广告代码库
	AccountId           int64     `json:"account_id"`              // 变现类型账户ID
	AppId               string    `json:"app_id"`                  // 应用ID（此处一般标识三方应用ID）
	AdRequests          int64     `json:"ad_requests"`             // 到达服务器的请求数量
	MatchedAdRequests   int64     `json:"matched_ad_requests"`     // 匹配到的到达广告请求数量
	ShowCount           int64     `json:"show_count"`              // 展示数
	ClickCount          int64     `json:"click_count"`             // 点击数
	AdRequestsMatchRate float64   `json:"ad_requests_match_rate"`  //填充率
	AdRequestsShowRate  float64   `json:"ad_requests_show_rate"`   // 请求展示率',
	ClickThroughRate    float64   `json:"click_through_rate"`      // 点击率',
	Earnings            float64   `json:"earnings"`                // 收入',
	ECpm                float64   `json:"ecpm" gorm:"column:ecpm"` // ECPM',
}

// NewRAC ReportAdsCollect 实例
func NewRAC(db *gorm.DB) *ReportAdsCollect {
	return &ReportAdsCollect{connectDb: connectDb{DB: db}}
}

func (m *ReportAdsCollect) TableName() string {
	return "report_ads_collects"
}

func (m *ReportAdsCollect) BatchInsert(realizations []*ReportAdsCollect) (err error) {
	if len(realizations) == 0 {
		return nil
	}
	updateColumns := []string{
		"ad_requests", "matched_ad_requests", "ad_requests_match_rate", "show_count", "ad_requests_show_rate",
		"click_count", "click_through_rate", "earnings", "ecpm",
	}

	err = m.Table(m.TableName()).Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(updateColumns),
	}).CreateInBatches(realizations, 500).Error
	return err
}
