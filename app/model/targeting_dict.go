package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type TargetingDict struct {
	connectDb

	DictKey    string `json:"dict_key"` // 属于什么字典
	Id         string `json:"id"`       // 字典的元素 ID
	Pid        string `json:"pid"`      // 父节点元素ID
	Label      string `json:"label"`    // 显示的内容
	Value      string `json:"value"`    // 元素的值
	Code       string `json:"code"`
	Seq        string `json:"seq"`
	DataStruct int64  `json:"data_struct"`
}

func NewTargetingDict(db *gorm.DB) *TargetingDict {
	return &TargetingDict{connectDb: connectDb{DB: db}}
}

func (m *TargetingDict) TableName() string {
	return "targeting_dictionaries"
}

func (m *TargetingDict) FindDictionaries(keys, values []string) (dictionaries []*TargetingDict, err error) {
	sb := m.Table(m.TableName())
	if len(keys) > 0 {
		sb = sb.Where("dict_key in ?", keys)
	}
	if len(values) > 0 {
		sb = sb.Where("value in ?", values)
	}
	err = sb.Order("id asc").Find(&dictionaries).Error
	return
}

func (m *TargetingDict) BatchInsert(dictionaries []*TargetingDict) (err error) {
	query := fmt.Sprintf("insert into %s (dict_key,id,pid,label,value,code,seq,data_struct) values ", m.TableName())
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	err = m.Transaction(func(session *gorm.DB) error {
		if err = session.Exec("TRUNCATE " + m.TableName()).Error; err != nil {
			return err
		}
		for i := 0; i < len(dictionaries); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?)")
			values = append(values,
				dictionaries[i].DictKey,
				dictionaries[i].Id,
				dictionaries[i].Pid,
				dictionaries[i].Label,
				dictionaries[i].Value,
				dictionaries[i].Code,
				dictionaries[i].Seq,
				dictionaries[i].DataStruct,
			)
			// 达到了 300 条数据，或最后一条了
			if chunk == 300 || i == len(dictionaries)-1 {
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
	return err
}
