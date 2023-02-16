package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	connectDb

	Id         int64     `json:"id"`
	SsoUid     string    `json:"sso_uid"`
	Email      string    `json:"email"`                  // 登陆邮箱
	Username   string    `json:"username"`               // 用户名
	Mobile     string    `json:"mobile"`                 // 手机号
	State      uint8     `json:"state"`                  // 账号状态
	RoleId     int64     `json:"role_id"`                // 角色ID
	IsInternal uint8     `json:"is_internal"`            // 是否内部账号
	Secret     string    `json:"-" gorm:"column:secret"` // 密码加密符
	Pass       string    `json:"-" gorm:"column:pass"`   // 密码
	CreatedAt  time.Time `json:"created_at"`             // 添加时间
	UpdatedAt  time.Time `json:"updated_at"`             // 最后一次修改时间
}

func (m *User) TableName() string {
	return "users"
}

func NewUser(db *gorm.DB) *User {
	return &User{connectDb: connectDb{DB: db}}
}

func (m *User) FindUserByEmail(email string) (user *User, err error) {
	err = m.Table(m.TableName()).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (m *User) FindUserBySso(email, ssoUid string) (user *User, err error) {
	err = m.Table(m.TableName()).Where("email = ? or sso_uid = ?", email, ssoUid).Where("is_internal = 1").First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (m *User) FindUserById(id int64) (user *User, err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).First(&user).Error
	return
}

func (m *User) UpdateUser(d map[string]interface{}, id int64) (err error) {
	err = m.Table(m.TableName()).Where("id = ? ", id).Updates(d).Error
	return
}

func (m *User) List(username, email string, state uint8, offset, limit int64) (users []*User, total int64, err error) {
	tbl := m.Table(m.TableName()).Where("state = ?", state)
	if len(username) > 0 {
		tbl = tbl.Where("username like ?", "%"+username+"%")
	}
	if len(email) > 0 {
		tbl = tbl.Where("email like ?", "%"+email+"%")
	}
	err = tbl.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	err = tbl.Offset(int(offset)).Limit(int(limit)).Order("id desc").Find(&users).Error
	return
}

func (m *User) CreateUser(user *User) (err error) {
	err = m.Table(m.TableName()).Create(user).Error
	return
}

func (m *User) Destroy(id int64) error {
	err := m.Exec(fmt.Sprintf("delete from %s where `id` = ?", m.TableName()), id).Error
	return err
}

func (m *User) Update(id int64, d map[string]int64) error {
	err := m.Table(m.TableName()).Where("id = ?", id).UpdateColumns(d).Error
	return err
}
