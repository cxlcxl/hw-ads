package vars

import (
	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
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
	Casbin    *casbin.SyncedEnforcer
	HLog      *logrus.Logger
)

type LoginUser struct {
	UserId     int64  `json:"user_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	RoleId     int64  `json:"role_id"`
	IsInternal uint8  `json:"is_internal"`
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
