package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	connectDb

	Id        int64     `json:"id"`
	RoleName  string    `json:"role_name"`
	State     int64     `json:"state"`      // 1正常0停用
	Sys       int64     `json:"sys"`        // 角色所属系统
	CreatedAt time.Time `json:"created_at"` // 添加时间
	UpdatedAt time.Time `json:"updated_at"` // 最后一次修改时间
}

func (m *Role) TableName() string {
	return "roles"
}

func NewRole(db *gorm.DB) *Role {
	return &Role{connectDb: connectDb{DB: db}}
}

func (m *Role) List(roleName string, state uint8) (roles []*Role, err error) {
	tbl := m.Table(m.TableName()).Where("state = ?", state)
	if len(roleName) > 0 {
		tbl = tbl.Where("role_name like ?", "%"+roleName+"%")
	}
	err = tbl.Order("id desc").Find(&roles).Error
	return
}

func (m *Role) CreateRole(role *Role) error {
	return m.Table(m.TableName()).Create(role).Error
}

func (m *Role) UpdateRole(d map[string]interface{}, id int64) error {
	return m.Table(m.TableName()).Where("id = ?", id).Updates(d).Error
}
