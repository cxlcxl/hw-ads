package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Permission struct {
	connectDb
	Id       int64         `json:"id"`
	Per      string        `json:"permission" gorm:"column:permission"`
	PName    string        `json:"p_name"`
	Method   string        `json:"method"`
	Pid      int64         `json:"pid"`
	Children []*Permission `json:"children" gorm:"-"`
}

var (
//cachePermissionKey = "db:permissions"
)

func (m *Permission) TableName() string {
	return "role_permissions"
}

func NewPermission(db *gorm.DB) *Permission {
	return &Permission{connectDb: connectDb{DB: db}}
}

func (m *Permission) PermissionCreate(p *Permission) error {
	return m.Table(m.TableName()).Create(&p).Error
}

func (m *Permission) Permissions() (ps []*Permission, err error) {
	err = m.Table(m.TableName()).Order("id desc").Find(&ps).Error
	return
}

func (m *Permission) PermissionUpdate(id int64, v map[string]interface{}) (err error) {
	return m.Table(m.TableName()).Where("id = ?", id).Updates(v).Error
}

func (m *Permission) PermissionDestroy(id int64) (err error) {
	var permission Permission
	err = m.Table(m.TableName()).Where("id = ?", id).First(&permission).Error
	if err != nil {
		return
	}
	delPrSQL := fmt.Sprintf("delete from %s where v1 = ? and v2 = ?", NewPR(nil).TableName())
	return m.Transaction(func(tx *gorm.DB) error {
		if err = tx.Exec(fmt.Sprintf("delete from %s where id = ?", m.TableName()), id).Error; err != nil {
			return err
		}
		if err = tx.Exec(delPrSQL, permission.Per, permission.Method).Error; err != nil {
			return err
		}
		return nil
	})
}

func (m *Permission) FindPermissionsByPers(ps []string) (_ps []*Permission, err error) {
	if len(ps) > 0 {
		err = m.Table(m.TableName()).Select("permission", "method").Where("permission in ?", ps).Find(&_ps).Error
	}
	return
}
