package middleware

import (
	"bs.mobgi.cc/app/vars"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var logExcludeApis = []string{
	"api/upload", "api/login", "api/logout", "api/profile", "api/settings/version", "api/settings/log",
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 会影响文件上传参数接收，文件上传不写日志
		if vars.YmlConfig.GetBool("Debug") || inExclude(c.Request.RequestURI) {
			return
		}

		beginTime := time.Now()
		c.Next()
		endTime := time.Now()
		user, _ := c.Get(vars.LoginUserKey)

		vars.HLog.WithFields(logrus.Fields{
			"module": "access-log",
			"uid":    user.(*vars.LoginUser).UserId,
			"uname":  user.(*vars.LoginUser).Username,
			"uri":    c.Request.RequestURI,
			"ip":     c.ClientIP(),
			"begin":  beginTime.Format(vars.DateTimeFormat),
			"end":    endTime.Format(vars.DateTimeFormat),
			"log_id": time.Now().UnixNano(),
		}).Info("access log")
	}
}

func inExclude(uri string) bool {
	for _, i2 := range logExcludeApis {
		if strings.Contains(uri, i2) {
			return true
		}
	}
	return false
}
