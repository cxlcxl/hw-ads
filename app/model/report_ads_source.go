package model

import (
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ReportAdsSource struct {
	connectDb

	Id                int64     `json:"id"`
	StatDay           time.Time `json:"stat_day"`            // 日: 日粒度，例如2021-09-08
	StatHour          uint8     `json:"stat_hour"`           // 时间: 小时
	Country           string    `json:"country"`             // 国家代码，使用华为开发者文档中的广告代码库
	AccountId         int64     `json:"account_id"`          // 变现类型账户ID
	AppId             string    `json:"app_id"`              // 应用ID（此处一般标识三方应用ID）
	AdType            string    `json:"ad_type"`             // 广告类型，字符串类型，例如banner，native
	PlacementId       string    `json:"placement_id"`        // 广告位ID
	AdRequests        int64     `json:"ad_requests"`         // 到达服务器的请求数量
	MatchedAdRequests int64     `json:"matched_ad_requests"` // 匹配到的到达广告请求数量
	ShowCount         int64     `json:"show_count"`          // 展示数
	ClickCount        int64     `json:"click_count"`         // 点击数
	Earnings          float64   `json:"earnings"`            // 收入',
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
		"ad_requests", "matched_ad_requests", "show_count", "click_count", "earnings",
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

type ReportAds struct {
	Ads
	AreaId int64 `json:"area_id"`
}

// AnalysisAds 变现报表数据
func (m *ReportAdsSource) AnalysisAds(
	accountIds, areas []int64, appIds, dates, countries, selects, groups []string, offset, limit int,
) (ads []*ReportAds, total int64, err error) {
	query := func(selects []string, _m string) string {
		return m.ToSQL(func(tx *gorm.DB) *gorm.DB {
			if _m == "count" {
				selects = []string{"1"}
			}
			query := tx.Table(m.TableName()).
				Where("stat_day between ? and ?", dates[0], dates[1]).
				Group("stat_day")

			if len(appIds) > 0 {
				query = query.Where("app_id in ?", appIds)
			}
			if len(countries) > 0 {
				query = query.Where("country in ?", countries)
			}
			if len(accountIds) > 0 {
				query = query.Where("account_id in ?", accountIds)
			}

			if utils.InArray(vars.ReportDimensionArea, groups) {
				if areaColumn := NewOverseasAreaRegion(m.DB).AreaColumnParse(areas); areaColumn != "" {
					if _m != "count" {
						selects = append(selects, areaColumn)
					} else {
						selects = []string{areaColumn}
					}
				}
			}
			if len(areas) > 0 {
				in := fmt.Sprintf(
					"country in (select c_code from `%s` where area_id in ?)",
					NewOverseasAreaRegion(nil).TableName(),
				)
				query = query.Where(in, areas)
			}

			for _, group := range groups {
				query = query.Group(group)
			}
			return query.Select(selects).Find(nil)
		})
	}

	if err = m.Raw(query(selects, "count")).Count(&total).Error; err != nil {
		return
	}
	err = m.Raw(query(selects, "ads")+" order by stat_day desc limit ? offset ?", limit, offset).Find(&ads).Error
	return
}
