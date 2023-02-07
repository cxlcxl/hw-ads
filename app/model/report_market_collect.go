package model

import (
	"bs.mobgi.cc/app/utils"
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
}

// ReportComprehensive 综合报表利用 JOIN 查询
func (m *ReportMarketCollect) ReportComprehensive(
	dates []string, actIds []int64, appIds, countries, marketSelects, adsSelects, groups, orders []string, offset, limit uint64,
) (markets []*ComprehensiveReport, total int64, err error) {
	// 自定义排序
	_os := make([]string, 0)
	if len(orders) > 0 {
		for _, order := range orders {
			_os = append(_os, order)
		}
	}

	_os = append(_os, "t0.stat_day desc")
	JoinOn := []string{"t0.stat_day = t1.stat_day"}
	// 统计行数字段「只取需要的字段」
	countSelects := []string{"stat_day"}
	// 变现表「包含账户维度需要查与账户关联的表」
	if utils.InArray("account_id", groups) {
		JoinOn = append(JoinOn, "t0.account_id = t1.account_id")
		_os = append(_os, "t0.account_id asc")
		countSelects = append(countSelects, "account_id")
	}

	if utils.InArray("app_id", groups) {
		JoinOn = append(JoinOn, "t0.app_id = t1.app_id")
		_os = append(_os, "t0.app_id asc")
		countSelects = append(countSelects, "app_id")
	}

	if utils.InArray("country", groups) {
		JoinOn = append(JoinOn, "t0.country = t1.country")
		_os = append(_os, "t0.country asc")
		countSelects = append(countSelects, "country")
	}

	// 变现表「包含账户维度需要查与账户关联的表」
	adsTable := NewRAC(nil).TableName()
	if utils.InArray("account_id", groups) {
		adsTable = NewRACC(nil).TableName()
	}
	var values []interface{}
	marketBase := m.comprehensiveBaseSQL(m.TableName(), actIds, appIds, countries, dates, groups)
	adsBase := m.comprehensiveBaseSQL(adsTable, actIds, appIds, countries, dates, groups)
	marketQuery, ms, err := marketBase(marketSelects)
	if err != nil {
		return
	}
	values = append(values, ms...)
	adsQuery, as, err := adsBase(adsSelects)
	if err != nil {
		return
	}
	values = append(values, as...)
	sql, _, err := squirrel.
		Select("t0.*,t1.`earnings`,t1.`ad_requests`,t1.`matched_ad_requests`,t1.`ad_show_count`,t1.`ad_click_count`").
		From(fmt.Sprintf("(%s) as t0", marketQuery)).
		LeftJoin(fmt.Sprintf("(%s) as t1 on %s", adsQuery, strings.Join(JoinOn, " and "))).
		OrderBy(_os...).Offset(offset).Limit(limit).ToSql()
	if err != nil {
		return
	}

	marketCount, _, err := marketBase(countSelects)
	if err != nil {
		return
	}
	adsCount, _, err := adsBase(countSelects)
	if err != nil {
		return
	}
	totalSql, _, err := squirrel.Select("count(1) as total").From(fmt.Sprintf("(%s) as t0", marketCount)).
		LeftJoin(fmt.Sprintf("(%s) as t1 on %s", adsCount, strings.Join(JoinOn, " and "))).ToSql()
	if err != nil {
		return
	}
	if err = m.Raw(totalSql, values...).Scan(&total).Error; err != nil || total == 0 {
		return
	}
	err = m.Debug().Raw(sql, values...).Find(&markets).Error
	return
}

func (m *ReportMarketCollect) comprehensiveBaseSQL(
	tableName string, actIds []int64, appIds, countries, dates, groups []string,
) func(selects []string) (string, []interface{}, error) {
	return func(s []string) (baseSQL string, values []interface{}, err error) {
		sql := squirrel.Select(s...).From(tableName).Where("stat_day between ? and ?", dates[0], dates[1]).GroupBy(groups...)

		// 组合维度筛选后的过滤条件
		if len(actIds) > 0 {
			in, vs := utils.WhereIn(actIds)
			sql = sql.Where("account_id in "+in, vs...)
		}
		if len(appIds) > 0 {
			in, vs := utils.WhereIn(appIds)
			sql = sql.Where("app_id in "+in, vs...)
		}
		if len(countries) > 0 {
			in, vs := utils.WhereIn(countries)
			sql = sql.Where("country in "+in, vs...)
		}

		baseSQL, vs, err := sql.ToSql()
		if err != nil {
			return
		}
		values = append(values, vs...)
		return
	}

}
