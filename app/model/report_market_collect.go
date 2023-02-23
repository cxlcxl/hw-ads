package model

import (
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
	"fmt"
	"github.com/Masterminds/squirrel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
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

type ComprehensiveReport struct {
	StatDay           time.Time `json:"stat_day"`
	Country           string    `json:"country"`
	AccountId         int64     `json:"account_id"`
	AppId             string    `json:"app_id"`
	AppName           string    `json:"app_name"`
	Cost              float64   `json:"cost"`
	ShowCount         int64     `json:"show_count"`
	ClickCount        int64     `json:"click_count"`
	DownloadCount     int64     `json:"download_count"`
	InstallCount      int64     `json:"install_count"`
	ActivateCount     int64     `json:"activate_count"`
	RetainCount       int64     `json:"retain_count"`
	ThreeRetainCount  int64     `json:"three_retain_count"`
	SevenRetainCount  int64     `json:"seven_retain_count"`
	AdRequests        int64     `json:"ad_requests"`
	MatchedAdRequests int64     `json:"matched_ad_requests"`
	AdShowCount       int64     `json:"ad_show_count"`
	AdClickCount      int64     `json:"ad_click_count"`
	Earnings          float64   `json:"earnings"`
	AreaId            int64     `json:"area_id"`
}

// ReportComprehensive 综合报表利用 JOIN 查询
func (m *ReportMarketCollect) ReportComprehensive(
	dates []string, actIds, areas []int64, appIds, countries []string,
	marketSelects, adsSelects, granularityColumns, groups, orders []string, offset, limit uint64, granularity string,
) (markets []*ComprehensiveReport, total int64, err error) {
	// 自定义排序
	_os := make([]string, 0)
	if len(orders) > 0 {
		for _, order := range orders {
			_os = append(_os, order)
		}
	}
	var JoinOn []string
	// 按日期粒度或者按整体没有任何维度筛选时
	if vars.ReportGranularityDate == granularity || len(groups) == 0 {
		_os = append(_os, "t0.stat_day desc")
		JoinOn = append(JoinOn, "t0.stat_day = t1.stat_day")
	}
	// 变现表「包含账户维度需要查与账户关联的表」
	if utils.InArray(vars.ReportDimensionAccount, groups) {
		JoinOn = append(JoinOn, "t0.account_id = t1.account_id")
		_os = append(_os, "t0.account_id asc")
	}
	if utils.InArray(vars.ReportDimensionApp, groups) {
		JoinOn = append(JoinOn, "t0.app_id = t1.app_id")
		_os = append(_os, "t0.app_id asc")
	}
	if utils.InArray(vars.ReportDimensionCountry, groups) {
		JoinOn = append(JoinOn, "t0.country = t1.country")
		_os = append(_os, "t0.country asc")
	}
	if utils.InArray(vars.ReportDimensionArea, groups) {
		JoinOn = append(JoinOn, "t0.area_id = t1.area_id")
		_os = append(_os, "t0.area_id asc")
	}

	// 变现表「包含账户维度需要查与账户关联的表」
	adsTable := NewRAC(nil).TableName()
	if utils.InArray(vars.ReportDimensionAccount, groups) {
		adsTable = NewRACA(nil).TableName()
	}
	// 生成投放与变现子查询的函数
	subQuery := func(table, _m string, selects []string) string {
		return m.ToSQL(func(tx *gorm.DB) *gorm.DB {
			query := tx.Table(table).Where("stat_day between ? and ?", dates[0], dates[1])
			if len(groups) > 0 {
				query = query.Group(strings.Join(groups, ","))
			}
			// 组合维度筛选后的过滤条件
			if len(actIds) > 0 {
				if _m == "ads" && !utils.InArray(vars.ReportDimensionAccount, groups) {
					in := fmt.Sprintf("app_id in (select app_id from `%s` where account_id in ?)", NewAppAct(nil).TableName())
					query = query.Where(in, actIds)
				} else {
					query = query.Where("account_id in ?", actIds)
				}
			}
			if len(appIds) > 0 {
				query = query.Where("app_id in ?", appIds)
			}
			if len(countries) > 0 {
				query = query.Where("country in ?", countries)
			}
			if utils.InArray(vars.ReportDimensionArea, groups) {
				if areaColumn := NewOverseasAreaRegion(m.DB).AreaColumnParse(areas); areaColumn != "" {
					if _m != "count" {
						selects = append(selects, areaColumn)
					} else {
						selects = []string{areaColumn}
					}
				}
				if len(areas) > 0 {
					in := fmt.Sprintf(
						"country in (select c_code from `%s` where area_id in ?)",
						NewOverseasAreaRegion(nil).TableName(),
					)
					query = query.Where(in, areas)
				}
			}

			return query.Select(selects).Find(nil)
		})
	}

	query := squirrel.
		Select(granularityColumns...).
		From(fmt.Sprintf("(%s) as t0", subQuery(m.TableName(), "market", marketSelects))).
		LeftJoin(fmt.Sprintf("(%s) as t1 on %s", subQuery(adsTable, "ads", adsSelects), strings.Join(JoinOn, " and "))).
		OrderBy(_os...)

	if limit > 0 {
		if err = m.Raw(subQuery(m.TableName(), "count", []string{"1"})).Count(&total).Error; err != nil || total == 0 {
			return nil, 0, err
		}
		query = query.Offset(offset).Limit(limit)
		if vars.ReportGranularityAll == granularity {
			query = query.GroupBy(groups...)
		}
	}
	var sql string
	if sql, _, err = query.ToSql(); err == nil {
		err = m.Raw(sql).Find(&markets).Error
	}

	return
}

type Summaries struct {
	Cost         float64 `json:"cost"`
	Earnings     float64 `json:"earnings"`
	Roi          float64 `json:"roi"`
	ECpm         float64 `json:"ecpm"`
	AdShowCount  int64   `json:"ad_show_count"`
	AdClickCount int64   `json:"ad_click_count"`
}

func (m *ReportMarketCollect) ComprehensiveSummaries(dates []string, actIds []int64, appIds, countries, marketSelects []string) (summaries Summaries) {
	query := m.Table(m.TableName()).Where("stat_day between ? and ?", dates[0], dates[1]).Select(marketSelects)
	// 变现表「包含账户维度需要查与账户关联的表」
	if len(actIds) > 0 {
		query = query.Where("account_id in ?", actIds)
	}
	if len(appIds) > 0 {
		query = query.Where("app_id in ?", appIds)
	}
	if len(countries) > 0 {
		query = query.Where("country in ?", countries)
	}
	query.Scan(&summaries)
	return
}
