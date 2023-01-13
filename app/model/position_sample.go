package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type PositionSample struct {
	connectDb

	Id                 int64  `json:"id"`
	CreativeSizeId     string `json:"creative_size_id"`     // 版位ID
	CreativeSizeSample string `json:"creative_size_sample"` // 预览图地址
	PreviewTitle       string `json:"preview_title"`        // 预览图标题
}

func NewPositionSample(db *gorm.DB) *PositionSample {
	return &PositionSample{connectDb: connectDb{DB: db}}
}

func (m *PositionSample) TableName() string {
	return "position_samples"
}

func (m *PositionSample) SampleList(creativeSizeIds []string) (samples []*PositionSample, err error) {
	err = m.Table(m.TableName()).Where("creative_size_id in ?", creativeSizeIds).Find(&samples).Error
	return
}

func (m *PositionSample) BatchInsert(samples []*PositionSample, creativeSizeIds []string) (err error) {
	deleteSql := fmt.Sprintf("delete from %s where creative_size_id in ?", m.TableName())
	if err = m.Exec(deleteSql, creativeSizeIds).Error; err != nil {
		return err
	}
	query := fmt.Sprintf("insert into %s (creative_size_id, creative_size_sample, preview_title) values ", m.TableName())
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	for i := 0; i < len(samples); i++ {
		valueStatement = append(valueStatement, "(?, ?, ?)")
		values = append(values, samples[i].CreativeSizeId, samples[i].CreativeSizeSample, samples[i].PreviewTitle)
		// 达到了 500 条数据，或最后一条了
		if chunk == 500 || i == len(samples)-1 {
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
