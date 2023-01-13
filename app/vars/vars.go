package vars

import (
	"log"
	"os"
	"strings"

	"bs.mobgi.cc/library/config_interface"
	myRedis "bs.mobgi.cc/library/redis"
	"gorm.io/gorm"
)

var (
	BasePath  string
	DBMysql   *gorm.DB
	YmlConfig config_interface.YamlConfigInterface
	DBRedis   *myRedis.DBRedis
)

type LoginUser struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	RoleId   int64  `json:"roleId"`
}

type TicketInfo struct {
	AppCode   string `json:"app_code"`
	ExpiredAt int64  `json:"expired_at"`
	Uid       string `json:"uid"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Mobile    string `json:"mobile"`
	RoleId    int64  `json:"role_id"`
	Pid       string `json:"pid"`
}
type SsoLoginResData struct {
	SsoUid   string `json:"uid"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Token    string `json:"token"`
	Username string `json:"username"`
}
type AccountTokenInfo struct {
	AccountId    int64  `json:"account_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int64  `json:"expired_at"`
	TokenType    string `json:"token_type"`
}

func init() {
	if dir, err := os.Getwd(); err != nil {
		log.Fatal("文件目录获取失败")
		return
	} else {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(dir, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = dir
		}
	}
}
