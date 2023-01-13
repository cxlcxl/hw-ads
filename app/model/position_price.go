package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type PositionPrice struct {
	connectDb

	Id             int64   `json:"id"`
	CreativeSizeId string  `json:"creative_size_id"` // 版位ID
	PriceType      string  `json:"price_type"`       // 付费方式
	BasePrice      float64 `json:"base_price"`       // 底价
}

func NewPositionPrice(db *gorm.DB) *PositionPrice {
	return &PositionPrice{connectDb: connectDb{DB: db}}
}

func (m *PositionPrice) TableName() string {
	return "position_base_prices"
}

func (m *PositionPrice) FindFloorPrice(creativeSizeId, priceType string) (f float64, err error) {
	err = m.Table(m.TableName()).Where("creative_size_id = ? and price_type = ?", creativeSizeId, priceType).Select("base_price").Scan(&f).Error
	return
}

func (m *PositionPrice) BatchInsert(basePrices []*PositionPrice) (err error) {
	query := fmt.Sprintf("insert into %s (creative_size_id,price_type,base_price) values ", m.TableName())
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	return m.Transaction(func(session *gorm.DB) error {
		if err = session.Exec("TRUNCATE " + m.TableName()).Error; err != nil {
			return err
		}
		for i := 0; i < len(basePrices); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?)")
			values = append(values, basePrices[i].CreativeSizeId, basePrices[i].PriceType, basePrices[i].BasePrice)
			// 达到了 300 条数据，或最后一条了
			if chunk == 800 || i == len(basePrices)-1 {
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
		return nil
	})
}
