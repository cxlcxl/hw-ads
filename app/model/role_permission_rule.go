package model

import (
	"bs.mobgi.cc/app/cache"
	"fmt"
	"gorm.io/gorm"
)

type PR struct {
	connectDb

	Id    int64  `json:"id"`
	PType string `json:"ptype" gorm:"column:ptype"`
	V0    string `json:"v0"`
	V1    string `json:"v1"`
	V2    string `json:"v2"`
	V3    string `json:"v3"`
	V4    string `json:"v4"`
	V5    string `json:"v5"`
}

var (
	cachePermissionKey = "db:permissions"
)

func (m *PR) TableName() string {
	return "role_permission_rules"
}

func NewPR(db *gorm.DB) *PR {
	return &PR{connectDb: connectDb{DB: db}}
}

func (m *PR) GetRolePermissions(roleId int64) (permissions []string, err error) {
	err = cache.New(m.DB).SetExpire(300).QueryRow(cachePermissionKey, &permissions, roleId, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Select("v1").Where("v0 = ?", id).Scan(v).Error
	})
	return
}

func (m *PR) PermissionDelete(roleId int64) error {
	return m.Exec(fmt.Sprintf("delete from %s where v0 = ?", m.TableName()), roleId).Error
}

func (m *PR) PermissionCreate(prs []*PR) error {
	return m.Table(m.TableName()).Create(prs).Error
}
