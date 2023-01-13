package model

import (
	"bs.mobgi.cc/app/utils"
	"fmt"
	"gorm.io/gorm"
)

type AppAccount struct {
	connectDb

	Id          int64  `json:"id"`
	AccountId   int64  `json:"account_id"`
	AppId       string `json:"app_id"`
	AccountType uint8  `json:"account_type"`
}

func (m *AppAccount) TableName() string {
	return "app_accounts"
}

func NewAppAct(db *gorm.DB) *AppAccount {
	return &AppAccount{connectDb: connectDb{DB: db}}
}

func (m *AppAccount) BatchInsert(d []*AppAccount) (err error) {
	if len(d) == 0 {
		return nil
	}
	sql := fmt.Sprintf("INSERT IGNORE INTO `%s`(`account_id`, `app_id`, `account_type`) VALUES", m.TableName())
	format := make([]string, len(d))
	values := make([]interface{}, 0)
	for i := 0; i < len(d); i++ {
		format[i] = "(?,?,?)"
		values = append(values, d[i].AccountId, d[i].AppId, d[i].AccountType)
	}
	err = m.Exec(sql+utils.BufferConcat(format, ","), values...).Error
	return err
}
