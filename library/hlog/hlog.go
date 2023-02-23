package hlog

import (
	"bs.mobgi.cc/app/vars"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

// NewHLog 初始化系统日志
func NewHLog() {
	vars.HLog = logrus.New()
	//vars.HLog.SetReportCaller(true) // 添加调用的函数和文件

	logName := vars.YmlConfig.GetString("Logs.SysLogName")
	maxRemainCnt := vars.YmlConfig.GetInt("Logs.MaxBackups")
	writer, err := rotatelogs.New(
		logName+".%Y%m%d",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),

		// WithRotationTime设置日志分割的时间
		rotatelogs.WithRotationTime(time.Hour*24),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(uint(maxRemainCnt)),
	)
	if err != nil {
		log.Fatal("日志初始化失败：" + err.Error())
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{TimestampFormat: vars.DateTimeFormat})

	vars.HLog.AddHook(lfsHook)
}

type HLog struct {
	l      *logrus.Logger
	level  logrus.Level
	module string
}

func NewLog(level logrus.Level, m string) *HLog {
	return &HLog{
		l:      vars.HLog,
		level:  level,
		module: m,
	}
}

func (hl *HLog) Log(fields map[string]interface{}, v interface{}) {
	fields["log_id"] = time.Now().UnixMicro()
	fields["module"] = hl.module
	switch hl.level {
	case logrus.WarnLevel:
		hl.l.WithFields(fields).Warn(v)
		break
	case logrus.ErrorLevel:
		hl.l.WithFields(fields).Error(v)
		break
	case logrus.FatalLevel:
		hl.l.WithFields(fields).Fatal(v)
		break
	default:
		hl.l.WithFields(fields).Info(v)
	}
	return
}
