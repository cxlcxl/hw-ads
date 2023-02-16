package model

import (
	"bs.mobgi.cc/app/cache"
	"fmt"
	"gorm.io/gorm"
)

type UserAccount struct {
	connectDb

	Id          int64 `json:"id"`
	AccountId   int64 `json:"account_id"` // 账户ID
	UserId      int64 `json:"user_id"`    //
	AccountType uint8 `json:"account_type"`
}

var (
	userActCacheKey = "db:user_account"
)

func NewUserAccount(db *gorm.DB) *UserAccount {
	return &UserAccount{connectDb: connectDb{DB: db}}
}

func (m *UserAccount) TableName() string {
	return "user_accounts"
}

func (m *UserAccount) FindAccountsByUserId(uid int64) (ua []*UserAccount, err error) {
	err = cache.New(m.DB).SetExpire(1800).QueryRow(userActCacheKey, &ua, uid, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Where("user_id = ?", id).Find(v).Error
	})
	return
}

func (m *UserAccount) SaveUserAccount(uid int64, ua []*UserAccount) error {
	return m.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf("delete from %s where user_id = ?", m.TableName()), uid).Error; err != nil {
			return err
		}
		if len(ua) > 0 {
			if err := tx.Table(m.TableName()).CreateInBatches(ua, 100).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
