package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ReportAdsCollect struct {
	connectDb
	Id                int64     `json:"id"`
	StatDay           time.Time `json:"stat_day"`            // 日: 日粒度，例如2021-09-08
	Country           string    `json:"country"`             // 国家代码，使用华为开发者文档中的广告代码库
	AdsAccountId      int64     `json:"ads_account_id"`      // 变现类型账户ID
	AppId             string    `json:"app_id"`              // 应用ID（此处一般标识三方应用ID）
	AdRequests        int64     `json:"ad_requests"`         // 到达服务器的请求数量
	MatchedAdRequests int64     `json:"matched_ad_requests"` // 匹配到的到达广告请求数量
	ShowCount         int64     `json:"show_count"`          // 展示数
	ClickCount        int64     `json:"click_count"`         // 点击数
	Earnings          float64   `json:"earnings"`            // 收入',
}

// NewRAC ReportAdsCollect 实例
func NewRAC(db *gorm.DB) *ReportAdsCollect {
	return &ReportAdsCollect{connectDb: connectDb{DB: db}}
}

func (m *ReportAdsCollect) TableName() string {
	return "report_ads_collects"
}

func (m *ReportAdsCollect) BatchInsert(collects []*ReportAdsCollect, collectActs []*ReportAdsCollectAct) (err error) {
	if len(collects) == 0 {
		return nil
	}
	updateColumns := []string{"ad_requests", "matched_ad_requests", "show_count", "click_count", "earnings"}

	err = m.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table(m.TableName()).Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns(updateColumns),
		}).CreateInBatches(collects, 500).Error; err != nil {
			return err
		}

		if err = NewRACC(tx).BatchInsert(collectActs); err != nil {
			return err
		}
		return nil
	})

	return err
}

// AnalysisComprehensive 综合报表投放数据部分
func (m *ReportAdsCollect) AnalysisComprehensive(appIds, dates, countries, selects, groups []string) (markets []*Ads, err error) {
	query := m.Table(m.TableName()).Select(selects).
		Where("stat_day between ? and ?", dates[0], dates[1]).
		Group("stat_day") // 变现数据以应用将数据匹配到投放数据上，所以应用必需分组
	if len(appIds) > 0 {
		query = query.Where("app_id in ?", appIds)
	}
	if len(countries) > 0 {
		query = query.Where("country in ?", countries)
	}

	for _, group := range groups {
		query = query.Group(group)
	}
	err = query.Find(&markets).Error
	return
}
