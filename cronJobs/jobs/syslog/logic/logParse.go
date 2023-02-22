package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type LogLogic struct {
	statDay string
	mu      sync.Mutex
}

func NewLogLogic(day string) *LogLogic {
	return &LogLogic{
		statDay: day,
		mu:      sync.Mutex{},
	}
}

func (l *LogLogic) Parse() error {
	filename := vars.YmlConfig.GetString("Logs.SysLogName")
	d := fmt.Sprintf("%s%s%s", l.statDay[:4], l.statDay[5:7], l.statDay[8:])
	f, err := os.Open(filename + "." + d)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := bufio.NewReader(f)

	sysLogs := make([]*model.SysLog, 0)
	i := 0
	for {
		i++
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			// 文件读取结束
			if len(sysLogs) > 0 {
				go l.saveLogData(sysLogs)
			}
			break
		}
		if sl, err := l.handleLog(line); err != nil {
			vars.HLog.WithFields(logrus.Fields{
				"module": "jobs-log-lineErr",
				"log_id": time.Now().UnixNano(),
				"day":    l.statDay,
			}).Error(string(line))
		} else {
			sysLogs = append(sysLogs, sl)
			if i >= 300 {
				go l.saveLogData(sysLogs)
				i = 0
			}
		}
	}
	return nil
}

// 处理读取出来的日志信息
func (l *LogLogic) handleLog(line []byte) (sysLog *model.SysLog, err error) {
	if err = json.Unmarshal(line, &sysLog); err != nil {
		return
	}
	logMessage := strings.TrimSpace(string(line))
	sysLog.Info = logMessage
	d, _ := time.Parse(vars.DateFormat, l.statDay)
	sysLog.StatDay = d
	return
}

func (l *LogLogic) saveLogData(sysLogs []*model.SysLog) {
	if err := model.NewLog(vars.DBMysql).BatchInsertLog(sysLogs); err != nil {
		vars.HLog.WithFields(logrus.Fields{
			"module": "jobs-log-saveErr",
			"log_id": time.Now().UnixNano(),
			"day":    l.statDay,
		}).Error(err)
	}
	return
}
