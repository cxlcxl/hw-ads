package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Campaign struct {
	connectDb

	Id                        int64     `json:"id"`
	CampaignId                string    `json:"campaign_id"`                  // 计划 ID
	CampaignName              string    `json:"campaign_name"`                // 计划名称
	AccountId                 int64     `json:"account_id"`                   // 账户 ID
	AppId                     string    `json:"app_id"`                       // 账户 ID
	AdvertiserId              string    `json:"advertiser_id"`                // 广告主账户 ID
	OptStatus                 string    `json:"opt_status"`                   // 操作状态
	CampaignDailyBudgetStatus string    `json:"campaign_daily_budget_status"` // 计划日预算状态
	ProductType               string    `json:"product_type"`                 // 推广产品类型
	ShowStatus                string    `json:"show_status"`                  // 计划状态
	UserBalanceStatus         string    `json:"user_balance_status"`          // 账户余额状态
	FlowResource              string    `json:"flow_resource"`                // 投放网络
	SyncFlowResource          string    `json:"sync_flow_resource"`           // 同时同步投放搜索广告网络
	CampaignType              string    `json:"campaign_type"`                // 计划类型
	TodayDailyBudget          int64     `json:"today_daily_budget"`           // 当日计划日限额
	TomorrowDailyBudget       int64     `json:"tomorrow_daily_budget"`        // 次日计划日限额，不返回表示与当日计划日限额相同
	MarketingGoal             string    `json:"marketing_goal"`               // 营销目标
	IsCallback                int64     `json:"is_callback"`                  // 是否通过查询计划任务回调完整信息
	CreatedAt                 time.Time `json:"created_at"`                   // 添加时间
	UpdatedAt                 time.Time `json:"updated_at"`                   // 最后一次修改时间
	App                       BelongApp `json:"app" gorm:"-"`
}

type BelongApp struct {
	ProductId string `json:"product_id"`
	AppName   string `json:"app_name"`
	IconUrl   string `json:"icon_url"`
}

func NewCampaign(db *gorm.DB) *Campaign {
	return &Campaign{connectDb: connectDb{DB: db}}
}

func (m *Campaign) TableName() string {
	return "campaigns"
}

func (m *Campaign) CampaignList(appId, campaignId, campaignName, campaignType string, offset, limit int64) (campaigns []*Campaign, total int64, err error) {
	query := m.Table(m.TableName()).Where("app_id = ?", appId)
	if len(campaignId) > 0 {
		query = query.Where("campaign_id like ?", "%"+campaignId+"%")
	}
	if len(campaignName) > 0 {
		query = query.Where("campaign_name like ?", "%"+campaignName+"%")
	}
	if len(campaignType) > 0 {
		query = query.Where("campaign_type = ?", campaignType)
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	if total > 0 {
		err = query.Offset(int(offset)).Limit(int(limit)).Order("updated_at desc").Find(&campaigns).Error
	}
	return
}

func (m *Campaign) FindByCampaignId(campaignId string) (campaign *Campaign, err error) {
	err = m.Table(m.TableName()).Where("campaign_id = ?", campaignId).First(&campaign).Error
	return
}

func (m *Campaign) BatchInsert(campaigns []*Campaign) (err error) {
	rows := "campaign_id,app_id,campaign_name,account_id,advertiser_id,opt_status,campaign_daily_budget_status,product_type," +
		"show_status,user_balance_status,flow_resource,sync_flow_resource,campaign_type,today_daily_budget," +
		"tomorrow_daily_budget,marketing_goal,is_callback,created_at,updated_at"
	query := fmt.Sprintf("insert into %s (%s) values ", m.TableName(), rows)
	updateFields := []string{
		"opt_status", "campaign_daily_budget_status", "show_status", "user_balance_status", "flow_resource",
		"today_daily_budget", "tomorrow_daily_budget", "marketing_goal", "is_callback",
	}
	fieldSql := make([]string, 0)
	for _, field := range updateFields {
		fieldSql = append(fieldSql, fmt.Sprintf("%s=values(%s)", field, field))
	}
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	return m.Transaction(func(session *gorm.DB) error {
		for i := 0; i < len(campaigns); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			data := campaigns[i]
			values = append(values,
				data.CampaignId, "", data.CampaignName, data.AccountId, data.AdvertiserId, data.OptStatus,
				data.CampaignDailyBudgetStatus, data.ProductType, data.ShowStatus, data.UserBalanceStatus, data.FlowResource,
				data.SyncFlowResource, data.CampaignType, data.TodayDailyBudget, data.TomorrowDailyBudget,
				data.MarketingGoal, 1, data.CreatedAt, data.UpdatedAt,
			)
			// 达到了 300 条数据，或最后一条了
			if chunk == 300 || i == len(campaigns)-1 {
				// 写入库
				insertSQL := query + strings.Join(valueStatement, ",") + " on duplicate key update " + strings.Join(fieldSql, ",")
				if err = session.Exec(insertSQL, values...).Error; err != nil {
					return err
				}
				// 重置
				values, valueStatement = make([]interface{}, 0), make([]string, 0)
				chunk = 0
			}
			chunk++
		}
		return nil
	})
}
