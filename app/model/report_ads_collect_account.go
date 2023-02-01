package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ReportAdsCollectAct struct {
	connectDb

	Id int64 `json:"id"`
	Comprehensive
}

type Comprehensive struct {
	StatDay             time.Time `json:"stat_day"`                // 日: 日粒度，例如2021-09-08
	Country             string    `json:"country"`                 // 国家代码，使用华为开发者文档中的广告代码库
	AccountId           int64     `json:"account_id"`              // 投放账户ID
	AdsAccountId        int64     `json:"ads_account_id"`          // 变现类型账户ID
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
	updateColumns := []string{
		"ad_requests", "matched_ad_requests", "ad_requests_match_rate", "show_count", "ad_requests_show_rate",
		"click_count", "click_through_rate", "earnings", "ecpm",
	}

	err = m.Table(m.TableName()).Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(updateColumns),
	}).CreateInBatches(realizations, 500).Error
	return err
}

// AnalysisComprehensive 综合报表投放数据部分「维度只有国家地区的情况」
func (m *ReportAdsCollectAct) AnalysisComprehensive(
	actIds []int64, appIds, dates, countries, selects, groups []string,
) (ads []*Comprehensive, err error) {
	query := m.Debug().Table(m.TableName()).Select(selects).
		Where("stat_day between ? and ?", dates[0], dates[1]).
		Group("stat_day,account_id")
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
		query = query.Group(group)
	}
	err = query.Find(&ads).Error
	return
}
