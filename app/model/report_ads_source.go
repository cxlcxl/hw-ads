package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ReportAdsSource struct {
	connectDb

	Id                  int64     `json:"id"`
	StatDay             time.Time `json:"stat_day"`                // 日: 日粒度，例如2021-09-08
	StatHour            uint8     `json:"stat_hour"`               // 时间: 小时
	Country             string    `json:"country"`                 // 国家代码，使用华为开发者文档中的广告代码库
	AccountId           int64     `json:"account_id"`              // 变现类型账户ID
	AppId               string    `json:"app_id"`                  // 应用ID（此处一般标识三方应用ID）
	AdType              string    `json:"ad_type"`                 // 广告类型，字符串类型，例如banner，native
	PlacementId         string    `json:"placement_id"`            // 广告位ID
	AdRequests          int64     `json:"ad_requests"`             // 到达服务器的请求数量
	MatchedAdRequests   int64     `json:"matched_ad_requests"`     // 匹配到的到达广告请求数量
	ShowCount           int64     `json:"show_count"`              // 展示数
	ClickCount          int64     `json:"click_count"`             // 点击数
	AdRequestsMatchRate float64   `json:"ad_requests_match_rate"`  //填充率
	AdRequestsShowRate  float64   `json:"ad_requests_show_rate"`   // 请求展示率',
	ClickThroughRate    float64   `json:"click_through_rate"`      // 点击率',
	Earnings            float64   `json:"earnings"`                // 收入',
	ECpm                float64   `json:"ecpm" gorm:"column:ecpm"` // ECPM',
	Timestamp
}

// NewRAS ReportAdsSource 实例
func NewRAS(db *gorm.DB) *ReportAdsSource {
	return &ReportAdsSource{connectDb: connectDb{DB: db}}
}

func (m *ReportAdsSource) TableName() string {
	return "report_ads_sources"
}

func (m *ReportAdsSource) BatchInsert(realizations []*ReportAdsSource) (err error) {
	if len(realizations) == 0 {
		return nil
	}
	updateColumns := []string{
		"ad_requests", "matched_ad_requests", "ad_requests_match_rate", "show_count", "ad_requests_show_rate",
		"click_count", "click_through_rate", "earnings", "ecpm",
	}

	err = m.Table(m.TableName()).Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(updateColumns),
	}).CreateInBatches(realizations, 300).Error
	return err
}

func (m *ReportAdsSource) CollectSources(day string) (realizations []*ReportAdsSource, err error) {
	columns := []string{
		"stat_day",
		"app_id",
		"account_id",
		"country",
		"round(sum(earnings), 5) as earnings",
		"sum(ad_requests) as ad_requests",
		"sum(matched_ad_requests) as matched_ad_requests",
		"sum(show_count) as show_count",
		"sum(click_count) as click_count",
	}
	err = m.Table(m.TableName()).Where("stat_day = ?", day).Select(columns).
		Group("stat_day,app_id,country").Find(&realizations).Error
	return
}

// AnalysisAds 变现报表数据
func (m *ReportAdsSource) AnalysisAds(
	accountIds []int64, appIds, dates, countries, selects, groups []string, offset, limit int,
) (ads []*Ads, total int64, err error) {
	query := m.Table(m.TableName()).Select(selects).
		Where("stat_day between ? and ?", dates[0], dates[1]).
		Group("stat_day") // 变现数据以应用将数据匹配到投放数据上，所以应用必需分组
	if len(appIds) > 0 {
		query = query.Where("app_id in ?", appIds)
	}
	if len(countries) > 0 {
		query = query.Where("country in ?", countries)
	}
	if len(accountIds) > 0 {
		query = query.Where("ads_account_id in ?", accountIds)
	}

	for _, group := range groups {
		query = query.Group(group)
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	err = query.Offset(offset).Limit(limit).Order("stat_day desc").Find(&ads).Error
	return
}
