package model

import (
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
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

// CollectAdsApps 变现数据过滤查【有投放的应用才整理】
func (m *AppAccount) CollectAdsApps() (apps []*AppAccount, err error) {
	err = m.Table(m.TableName()).
		Where("account_type = ?", vars.AccountTypeMarket).
		Select("app_id,account_id").Find(&apps).Error
	return
}

func (m *AppAccount) FindAppIdsByAccountIds(accountIds []int64) (appIds []string) {
	m.Table(m.TableName()).Where("account_id in ?", accountIds).Select("app_id").Find(&appIds)
	return
}

func (m *AppAccount) FindAccountIdsByAppIds(appIds []string) (v []*AppAccount, err error) {
	m.Table(m.TableName()).Where("app_id in ? and account_type = ?", appIds, vars.AccountTypeMarket).
		Select("app_id,account_id").Find(&v)
	return
}

func (m *AppAccount) FlushInfo() (err error) {
	err = m.Exec(fmt.Sprintf("delete from `%s`", m.TableName())).Error
	return
}
