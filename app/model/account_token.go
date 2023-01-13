package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	connectDb

	Id           int64     `db:"id"`
	AccountId    int64     `db:"account_id"`    // 账户ID
	AdvertiserId string    `db:"advertiser_id"` //
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	ExpiredAt    time.Time `db:"expired_at"` // access_token 过期时间
	TokenType    string    `db:"token_type"`
	Timestamp
}

func NewToken(db *gorm.DB) *Token {
	return &Token{connectDb: connectDb{DB: db}}
}

func (m *Token) TableName() string {
	return "account_tokens"
}

func (m *Token) FindByAccountId(accountId int64) (token *Token, err error) {
	err = m.Table(m.TableName()).Where("account_id = ?", accountId).First(&token).Error
	return
}

func (m *Token) TokenUpdate(d map[string]interface{}, id int64) (err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).Updates(d).Error
	return
}

func (m *Token) GetAccessTokenList() (tokens []*Token, err error) {
	err = m.Table(m.TableName()).Find(&tokens).Error
	return
}

func (m *Token) ReportAccessTokens(accountType int64) (tokens []*Token, err error) {
	sql := fmt.Sprintf(
		"select t0.* from %s t0 left join %s t1 on t0.account_id = t1.id where t1.account_type = ?",
		m.TableName(),
		NewAct(nil).TableName(),
	)
	err = m.Raw(sql, accountType).Find(&tokens).Error
	return
}
