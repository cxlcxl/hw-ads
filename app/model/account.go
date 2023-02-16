package model

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	connectDb

	Id           int64     `json:"id"`
	ParentId     int64     `json:"parent_id"`     // 所属上级服务商
	AdvertiserId string    `json:"advertiser_id"` // 广告主账户ID
	DeveloperId  string    `json:"developer_id"`  // 开发者ID
	AccountType  uint8     `json:"account_type"`  // 账户类型
	State        int64     `json:"state"`         // 状态
	AccountName  string    `json:"account_name"`  // 账户名
	ClientId     string    `json:"client_id"`     // 客户端ID
	IsAuth       int64     `json:"is_auth"`       // 是否已认证
	Secret       string    `json:"secret"`        // 密钥
	CreatedAt    time.Time `json:"created_at"`    // 添加时间
	UpdatedAt    time.Time `json:"updated_at"`    // 最后一次修改时间
}

type BelongAccount struct {
	Id          int64  `json:"id"`
	AccountType int64  `json:"account_type"` // 账户类型
	AccountName string `json:"account_name"` // 账户名
}

func (m *Account) TableName() string {
	return "accounts"
}

func NewAct(db *gorm.DB) *Account {
	return &Account{connectDb: connectDb{DB: db}}
}

func (m *Account) FindAccountById(id int64) (act *Account, err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).First(&act).Error
	return
}

func (m *Account) AccountList(ids []int64, accountType, state int64, accountName string, offset, limit int64) (ls []*Account, total int64, err error) {
	tbl := m.Table(m.TableName()).Where("state = ?", state).Order("id desc")
	if len(accountName) > 0 {
		tbl = tbl.Where("account_name like ?", "%"+accountName+"%")
	}
	if len(ids) > 0 {
		tbl = tbl.Where("id in ?", ids)
	}
	if accountType > 0 {
		tbl = tbl.Where("account_type = ?", accountType)
	}
	if err = tbl.Count(&total).Error; err != nil {
		return
	}
	if total > 0 {
		err = tbl.Offset(int(offset)).Limit(int(limit)).Find(&ls).Error
	}
	return
}

func (m *Account) RemoteAccounts(accountName string, isParent int64) (accounts []*Account, err error) {
	query := m.Table(m.TableName()).Where("state = 1")
	if isParent == 1 {
		query = query.Where("parent_id = 0")
	}
	if accountName != "" {
		query = query.Where("account_name like ?", "%"+accountName+"%")
	}
	err = query.Order("updated_at desc").Offset(0).Limit(20).Find(&accounts).Error
	return
}

func (m *Account) AllAccounts(ids []int64) (accounts []*BelongAccount, err error) {
	query := m.Table(m.TableName()).Where("state = 1")
	if len(ids) > 0 {
		query = query.Where("id in ?", ids)
	}
	err = query.Find(&accounts).Error
	return
}

func (m *Account) AccountCreate(act *Account, userAct UserAccount) (err error) {
	err = m.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table(m.TableName()).Create(act).Error; err != nil {
			return err
		}
		if err = tx.Table(NewToken(nil).TableName()).Create(&Token{
			AccountId: act.Id,
			ExpiredAt: act.CreatedAt,
			TokenType: "Bearer",
			Timestamp: Timestamp{CreatedAt: act.CreatedAt, UpdatedAt: act.CreatedAt},
		}).Error; err != nil {
			return err
		}
		userAct.AccountId = act.Id
		if err = tx.Table(NewUserAccount(nil).TableName()).Create(&userAct).Error; err != nil {
			return err
		}
		return nil
	})
	return
}

func (m *Account) AccountUpdate(d map[string]interface{}, id int64) (err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).Updates(d).Error
	return
}

func (m *Account) GetAccountsByIds(ids []int64) (accounts []*Account, err error) {
	err = m.Table(m.TableName()).Where("id in ?", ids).
		Select("id", "account_name", "state", "account_type", "developer_id", "advertiser_id").
		Find(&accounts).Error
	return
}

type ClientInfo struct {
	Id       int64  `json:"id"`
	ClientId string `json:"client_id"` // 客户端ID
	Secret   string `json:"secret"`    // 密钥
}

func (m *Account) ClientInfo(accountType int64) (info []*ClientInfo, err error) {
	err = m.Table(m.TableName()).Where("account_type = ?", accountType).
		Select("id", "client_id", "secret").
		Find(&info).Error
	return
}
