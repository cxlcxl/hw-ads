package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ReportMarketSource struct {
	connectDb

	Id                   int64     `json:"id"`
	StatDay              time.Time `json:"stat_day"`               // 日: 日粒度，例如2021-09-08
	StatHour             uint8     `json:"stat_hour"`              // 时间: 小时
	Country              string    `json:"country"`                // 国家代码，使用华为开发者文档中的广告代码库
	AccountId            int64     `json:"account_id"`             // 应用所属账户ID
	AppId                string    `json:"app_id"`                 // 应用ID（此处一般标识三方应用ID）
	AppName              string    `json:"app_name"`               // 应用名称
	PkgName              string    `json:"pkg_name"`               // 投放的应用包名
	CampaignId           string    `json:"campaign_id"`            // 广告计划ID
	CampaignName         string    `json:"campaign_name"`          // 广告计划名称
	AdgroupId            string    `json:"adgroup_id"`             // 广告任务ID
	AdgroupName          string    `json:"adgroup_name"`           // 广告任务名称
	CreativeId           string    `json:"creative_id"`            // 广告创意ID
	CreativeName         string    `json:"creative_name"`          // 广告创意名称
	Cost                 float64   `json:"cost"`                   // 花费
	ShowCount            int64     `json:"show_count"`             // 展示数
	ClickCount           int64     `json:"click_count"`            // 点击数
	DownloadCount        int64     `json:"download_count"`         // 下载数
	InstallCount         int64     `json:"install_count"`          // 安装数
	ActivateCount        int64     `json:"activate_count"`         // 激活数
	RetainCount          int64     `json:"retain_count"`           // 留存数
	ThreeRetainCount     int64     `json:"three_retain_count"`     // 三日留存数
	SevenRetainCount     int64     `json:"seven_retain_count"`     // 七日留存数
	ClickThroughRate     float64   `json:"click_through_rate"`     // 点击率
	ClickDownloadRate    float64   `json:"click_download_rate"`    // 点击下载转化率
	DownloadActivateRate float64   `json:"download_activate_rate"` // 下载激活转化率
	Cpm                  float64   `json:"cpm"`                    // CPM(千人成本): 花费*1000/展示量
	Cpc                  float64   `json:"cpc"`                    // CPC: 花费/点击量
	Cpd                  float64   `json:"cpd"`                    // CPD: 花费/下载数
	Cpi                  float64   `json:"cpi"`                    // CPI: 花费/安装数
	Cpa                  float64   `json:"cpa"`                    // CPA: 花费/激活数
	SevenRetainCost      float64   `json:"seven_retain_cost"`      // 平均留存花费=花费/七日留存数
	RetainCost           float64   `json:"retain_cost"`            // 平均留存花费=花费/留存数
	ThreeRetainCost      float64   `json:"three_retain_cost"`      // 平均留存花费=花费/三日留存数
	Timestamp
}

// NewRMS ReportMarketSource 实例
func NewRMS(db *gorm.DB) *ReportMarketSource {
	return &ReportMarketSource{connectDb: connectDb{DB: db}}
}

func (m *ReportMarketSource) TableName() string {
	return "report_market_sources"
}

func (m *ReportMarketSource) BatchInsert(markets []*ReportMarketSource) (err error) {
	if len(markets) == 0 {
		return nil
	}
	updateColumns := []string{
		"cost", "show_count", "click_count", "download_count", "install_count", "activate_count", "retain_count",
		"click_through_rate", "click_download_rate", "download_activate_rate", "cpm", "cpd", "cpc", "cpi", "cpa", "retain_cost",
	}
	return m.Table(m.TableName()).Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(updateColumns),
	}).CreateInBatches(markets, 500).Error
}

func (m *ReportMarketSource) CollectSources(day string) (markets []*ReportMarketSource, err error) {
	columns := []string{
		"stat_day",
		"app_id",
		"app_name",
		"account_id",
		"country",
		"round(sum(cost), 5) as cost",
		"sum(show_count) as show_count",
		"sum(click_count) as click_count",
		"sum(download_count) as download_count",
		"sum(install_count) as install_count",
		"sum(activate_count) as activate_count",
		"sum(retain_count) as retain_count",
		"sum(three_retain_count) as three_retain_count",
		"sum(seven_retain_count) as seven_retain_count",
	}
	err = m.Table(m.TableName()).Where("stat_day = ?", day).Select(columns).
		Group("stat_day,app_id,country,account_id").Find(&markets).Error
	return
}
