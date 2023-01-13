package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type PositionElement struct {
	connectDb

	Id              int64  `json:"id"`
	CreativeSizeId  string `json:"creative_size_id"`   // 版位ID
	SubType         string `json:"sub_type"`           // 版位子形式
	GroupNumber     int64  `json:"group_number"`       // 版位子形式分组
	ElementId       string `json:"element_id"`         // 版位元素id
	ElementName     string `json:"element_name"`       // 版位元素类型
	ElementTitle    string `json:"element_title"`      // 版位元素名称
	ElementCaption  string `json:"element_caption"`    // 版位元素描述
	Width           int64  `json:"width"`              // 图片宽
	Height          int64  `json:"height"`             // 图片高
	MinWidth        int64  `json:"min_width"`          // 视频最小宽度
	MinHeight       int64  `json:"min_height"`         // 视频最小高度
	MinLength       int64  `json:"min_length"`         // 最小输入长度
	MaxLength       int64  `json:"max_length"`         // 文案、摘要、品牌名称，都是指中文长度，英文长度
	FileSizeKbLimit int64  `json:"file_size_kb_limit"` // 文件大小上限，单位KB
	GifSizeKbLimit  int64  `json:"gif_size_kb_limit"`  // Gif文件大小上限，单位KB
	FileFormat      string `json:"file_format"`        // 文件类型
	Pattern         string `json:"pattern"`            // 输入校验规则
	Duration        string `json:"duration"`           // 视频时长
	MinOccurs       string `json:"min_occurs"`         // 最小出现次数，为0表示元素为可选
	MaxOccurs       string `json:"max_occurs"`         // 最大出现次数
}

func NewPositionElement(db *gorm.DB) *PositionElement {
	return &PositionElement{connectDb: connectDb{DB: db}}
}

func (m *PositionElement) TableName() string {
	return "position_elements"
}

func (m *PositionElement) BatchInsert(elements []*PositionElement) (err error) {
	rows := "creative_size_id,sub_type,group_number,element_id,element_name,element_title,element_caption,width,height," +
		"min_width,min_height,min_length,max_length,file_size_kb_limit,gif_size_kb_limit,file_format,pattern,duration," +
		"min_occurs,max_occurs"
	query := fmt.Sprintf("insert ignore into %s (%s) values ", m.TableName(), rows)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	return m.Transaction(func(session *gorm.DB) error {
		if err = session.Exec("TRUNCATE " + m.TableName()).Error; err != nil {
			return err
		}
		for i := 0; i < len(elements); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			values = append(values,
				elements[i].CreativeSizeId, elements[i].SubType, elements[i].GroupNumber, elements[i].ElementId, elements[i].ElementName,
				elements[i].ElementTitle, elements[i].ElementCaption, elements[i].Width, elements[i].Height,
				elements[i].MinWidth, elements[i].MinHeight, elements[i].MinLength, elements[i].MaxLength,
				elements[i].FileSizeKbLimit, elements[i].GifSizeKbLimit, elements[i].FileFormat, elements[i].Pattern,
				elements[i].Duration, elements[i].MinOccurs, elements[i].MaxOccurs,
			)
			// 达到了 200 条数据，或最后一条了
			if chunk == 200 || i == len(elements)-1 {
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
