package vars

import "errors"

var (
	ErrLoginUserLocked          = errors.New("当前账号已锁定，不可登陆")
	ErrLoginUnknownEmail        = errors.New("用户不存在")
	ErrLoginPass                = errors.New("密码错误")
	ErrLoginTicketExipred       = errors.New("登陆凭证不存在")
	ErrLoginTicketBuildFailed   = errors.New("登陆凭证生成失败")
	ErrLoginTicketSessionFailed = errors.New("登陆凭证存储失败")
)
