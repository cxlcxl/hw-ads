package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ReportAdsCollectAct struct {
	connectDb

	Id int64 `json:"id"`
	Ads
}

type Ads struct {
	StatDay           time.Time `json:"stat_day"`            // 日: 日粒度，例如2021-09-08
	Country           string    `json:"country"`             // 国家代码，使用华为开发者文档中的广告代码库
	AccountId         int64     `json:"account_id"`          // 投放账户ID
	AdsAccountId      int64     `json:"ads_account_id"`      // 变现类型账户ID
	AppId             string    `json:"app_id"`              // 应用ID（此处一般标识三方应用ID）
	AdRequests        int64     `json:"ad_requests"`         // 到达服务器的请求数量
	MatchedAdRequests int64     `json:"matched_ad_requests"` // 匹配到的到达广告请求数量
	ShowCount         int64     `json:"show_count"`          // 展示数
	ClickCount        int64     `json:"click_count"`         // 点击数
	Earnings          float64   `json:"earnings"`            // 收入',
}

// NewRACC ReportAdsCollectAct 实例
func NewRACC(db *gorm.DB) *ReportAdsCollectAct {
	return &ReportAdsCollectAct{connectDb: connectDb{DB: db}}
}

func (m *ReportAdsCollectAct) TableName() string {
	return "report_ads_collect_accounts"
}

func (m *ReportAdsCollectAct) BatchInsert(realizations []*ReportAdsCollectAct) (err error) {
	if len(realizations) == 0 {
		return nil
	}
	updateColumns := []string{"ad_requests", "matched_ad_requests", "show_count", "click_count", "earnings"}

	err = m.Table(m.TableName()).Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(updateColumns),
	}).CreateInBatches(realizations, 500).Error
	return err
}
