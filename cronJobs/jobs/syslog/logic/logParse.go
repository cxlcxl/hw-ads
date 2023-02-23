package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/hlog"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type LogLogic struct {
	statDay string
	mu      sync.Mutex
	errs    []*LogInfo
}

type LogInfo struct {
	StatDay time.Time `json:"-"`
	Module  string    `json:"module"`
	Msg     string    `json:"msg"`
	Level   string    `json:"level"`
	LogId   int       `json:"log_id"`
	Info    string    `json:"info"`
}

func NewLogLogic(day string) *LogLogic {
	return &LogLogic{
		statDay: day,
		mu:      sync.Mutex{},
		errs:    make([]*LogInfo, 0),
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
	i, k := 0, 1
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			// 文件读取结束
			if len(sysLogs) > 0 {
				l.saveLogData(sysLogs)
			}
			break
		}
		sl, err := l.handleLog(line)
		if err != nil {
			// 错误级别的日志再记录，其他的丢弃
			if sl.Level == logrus.ErrorLevel.String() {
				hlog.NewLog(logrus.ErrorLevel, "jobs-log-line").Log(logrus.Fields{
					"day":      l.statDay,
					"err_line": k,
				}, err)
			}
		} else {
			sysLogs = append(sysLogs, &model.SysLog{
				StatDay: sl.StatDay,
				Module:  sl.Module,
				Msg:     sl.Msg,
				Info:    sl.Info,
				Level:   sl.Level,
				LogId:   strconv.Itoa(sl.LogId),
			})
			if sl.Level == logrus.ErrorLevel.String() {
				l.errs = append(l.errs, &LogInfo{
					StatDay: sl.StatDay,
					Module:  sl.Module,
					Msg:     sl.Msg,
					Level:   sl.Level,
					LogId:   sl.LogId,
				})
			}
			if i >= 300 {
				l.saveLogData(sysLogs)
				i = 0
			} else {
				i++
			}
		}
		k++
	}

	if len(l.errs) > 0 {
		// TODO 发现错误，告警通知
	}
	return nil
}

// 处理读取出来的日志信息
func (l *LogLogic) handleLog(line []byte) (sysLog *LogInfo, err error) {
	if err = json.Unmarshal(line, &sysLog); err != nil {
		return
	}
	logMessage := strings.TrimSpace(string(line))
	sysLog.Info = logMessage
	t, _ := time.Parse(vars.DateFormat, l.statDay)
	sysLog.StatDay = t
	return
}

func (l *LogLogic) saveLogData(sysLogs []*model.SysLog) {
	if err := model.NewLog(vars.DBMysql).BatchInsertLog(sysLogs); err != nil {
		hlog.NewLog(logrus.ErrorLevel, "jobs-log-save").Log(logrus.Fields{
			"day": l.statDay,
		}, err)
	}
	return
}
