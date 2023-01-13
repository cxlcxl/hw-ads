package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Position struct {
	connectDb

	Id                         int64     `json:"id"`
	AccountId                  int64     `json:"account_id"`
	AdvertiserId               string    `json:"advertiser_id"`             // 广告主账户ID
	CreativeSizeId             string    `json:"creative_size_id"`          // 版位ID
	CreativeSizeNameDsp        string    `json:"creative_size_name_dsp"`    // 版位名称
	CreativeSizeDescription    string    `json:"creative_size_description"` // 版位描述
	Category                   string    `json:"category"`                  // 版位所属分类
	SupportProductType         string    `json:"support_product_type"`      // 支持的推广产品
	SupportObjectiveType       string    `json:"support_objective_type"`
	IsSupportTimePeriod        string    `json:"is_support_time_period"`        // 是否支持选择投放时段
	IsSupportMultipleCreatives string    `json:"is_support_multiple_creatives"` // 是否支持多创意
	SupportPriceType           string    `json:"support_price_type"`            // 付费方式
	LastPullTime               time.Time `json:"last_pull_time"`                // 最后拉取时间
}

func NewPosition(db *gorm.DB) *Position {
	return &Position{connectDb: connectDb{DB: db}}
}

func (m *Position) TableName() string {
	return "positions"
}

func (m *Position) PositionList(category, productType string, accountId int64) (positions []*Position, err error) {
	err = m.Table(m.TableName()).Where("category = ? and account_id = ? and FIND_IN_SET(?, support_product_type)", category, accountId, productType).Find(&positions).Error
	return
}

func (m *Position) BatchInsert(positions []*Position, samples []*PositionSample, placements []*PositionPlacement) (err error) {
	if len(positions) == 0 {
		return
	}
	creativeSizeIds := make([]string, 0)
	for _, position := range positions {
		creativeSizeIds = append(creativeSizeIds, position.CreativeSizeId)
	}
	deleteSql := fmt.Sprintf("delete from %s where creative_size_id in ?", m.TableName())
	rows := "account_id,advertiser_id,creative_size_id,creative_size_name_dsp,creative_size_description,category," +
		"support_product_type,support_objective_type,is_support_time_period,is_support_multiple_creatives,support_price_type,last_pull_time"
	query := fmt.Sprintf("insert into %s (%s) values ", m.TableName(), rows)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	return m.Transaction(func(session *gorm.DB) error {
		if err = session.Exec(deleteSql, creativeSizeIds).Error; err != nil {
			return err
		}
		for i := 0; i < len(positions); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			values = append(values,
				positions[i].AccountId, positions[i].AdvertiserId, positions[i].CreativeSizeId, positions[i].CreativeSizeNameDsp,
				positions[i].CreativeSizeDescription, positions[i].Category, positions[i].SupportProductType,
				positions[i].SupportObjectiveType, positions[i].IsSupportTimePeriod, positions[i].IsSupportMultipleCreatives,
				positions[i].SupportPriceType, positions[i].LastPullTime,
			)
			// 达到了 300 条数据，或最后一条了
			if chunk == 300 || i == len(positions)-1 {
				// 写入库
				insertSQL := query + strings.Join(valueStatement, ",")
				if err = session.Exec(insertSQL, values...).Error; err != nil {
					return err
				}
				// 重置
				values, valueStatement = make([]interface{}, 0), make([]string, 0)
				chunk = 0
			}
			chunk++
		}
		if err = NewPositionSample(session).BatchInsert(samples, creativeSizeIds); err != nil {
			return err
		}
		if err = NewPositionPlacement(session).BatchInsert(placements, creativeSizeIds); err != nil {
			return err
		}
		return nil
	})
}

func (m *Position) GetPositionByAccountIds(accountIds []int64) (positions []*Position, err error) {
	if len(accountIds) == 0 {
		return nil, err
	}
	fields := []string{
		"id",
		"account_id",
		"advertiser_id",
		"creative_size_id",
		"'' as creative_size_name_dsp",
		"'' as creative_size_description",
		"'' as category",
		"'' as support_product_type",
		"'' as support_objective_type",
		"'' as is_support_time_period",
		"'' as is_support_multiple_creatives",
		"support_price_type",
		"last_pull_time",
	}
	err = m.Table(m.TableName()).Where("account_id in ?", accountIds).Select(fields).Find(&positions).Error
	return
}
