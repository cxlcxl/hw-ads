package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type PositionPlacement struct {
	connectDb

	Id                         int64  `json:"id"`
	CreativeSizeId             string `json:"creative_size_id"`              // 版位ID
	PlacementSizeId            string `json:"placement_size_id"`             // 规格ID
	CreativeSize               string `json:"creative_size"`                 // 尺寸
	CreativeSizeSubType        string `json:"creative_size_sub_type"`        // 版位子形式
	IsSupportMultipleCreatives string `json:"is_support_multiple_creatives"` // 是否支持多创意
}

func NewPositionPlacement(db *gorm.DB) *PositionPlacement {
	return &PositionPlacement{connectDb: connectDb{DB: db}}
}

func (m *PositionPlacement) TableName() string {
	return "position_placements"
}

func (m *PositionPlacement) PlacementList(creativeSizeIds []string) (placements []*PositionPlacement, err error) {
	err = m.Table(m.TableName()).Where("creative_size_id in ?", creativeSizeIds).Find(&placements).Error
	return
}

func (m *PositionPlacement) PlacementsByCreativeSizeId(creativeSizeId string) (placements []*PositionPlacement, err error) {
	err = m.Table(m.TableName()).Select("creative_size,creative_size_sub_type").
		Where("creative_size_id = ?", creativeSizeId).
		Group("creative_size_sub_type,creative_size").
		Order("creative_size_sub_type asc,creative_size asc").Find(&placements).Error
	return
}

func (m *PositionPlacement) BatchInsert(placements []*PositionPlacement, creativeSizeIds []string) (err error) {
	deleteSql := fmt.Sprintf("delete from %s where creative_size_id in ?", m.TableName())
	if err = m.Exec(deleteSql, creativeSizeIds).Error; err != nil {
		return err
	}
	query := fmt.Sprintf("insert into %s (creative_size_id, placement_size_id, creative_size, creative_size_sub_type, is_support_multiple_creatives) values ", m.TableName())
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	for i := 0; i < len(placements); i++ {
		valueStatement = append(valueStatement, "(?, ?, ?, ?, ?)")
		values = append(values,
			placements[i].CreativeSizeId,
			placements[i].PlacementSizeId,
			placements[i].CreativeSize,
			placements[i].CreativeSizeSubType,
			placements[i].IsSupportMultipleCreatives,
		)
		// 达到了 500 条数据，或最后一条了
		if chunk == 500 || i == len(placements)-1 {
			// 写入库
			insertSQL := query + strings.Join(valueStatement, ",")
			if err = m.Exec(insertSQL, values...).Error; err != nil {
				return err
			}
			// 重置
			values, valueStatement = make([]interface{}, 0), make([]string, 0)
			chunk = 0
		}
		chunk++
	}
	return nil
}
