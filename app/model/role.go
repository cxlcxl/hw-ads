package model

import (
	"gorm.io/gorm"
)

type Role struct {
	connectDb

	Id       int64  `json:"id"`
	RoleName string `json:"role_name"`
	State    uint8  `json:"state"` // 1正常0停用
	Sys      uint8  `json:"sys"`   // 角色所属系统
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

func (m *Role) FindRoleById(id int64) (role *Role, err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).First(&role).Error
	return
}

func (m *Role) UpdateRole(d map[string]interface{}, id int64, prs []*PR) error {
	return m.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(m.TableName()).Where("id = ?", id).Updates(d).Error; err != nil {
			return err
		}
		if err := NewPR(tx).PermissionDelete(id); err != nil {
			return err
		}
		if err := NewPR(tx).PermissionCreate(prs); err != nil {
			return err
		}
		return nil
	})
}
