package model

import (
	"gorm.io/gorm"
)

type UserToken struct {
	connectDb
	Id         int64  `json:"id"`
	Uid        string `json:"uid"`
	Token      string `json:"token"`
	ActionName string `json:"action_name"`
	Scopes     string `json:"scopes"`
	ClientIp   string `json:"client_ip"`
	ExpiredAt  int64  `json:"expired_at"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func (m *UserToken) TableName() string {
	return "user_tokens"
}

func NewUserToken(db *gorm.DB) *UserToken {
	return &UserToken{connectDb: connectDb{DB: db}}
}
