package model

import (
	"gorm.io/gorm"
	"time"
)

type SysLog struct {
	connectDb
	Id      int64     `json:"id"`
	StatDay time.Time `json:"stat_day"`
	Module  string    `json:"module"`
	Msg     string    `json:"msg"`
	Info    string    `json:"info"`
	Level   string    `json:"level"`
	LogId   string    `json:"log_id"`
}

func (m *SysLog) TableName() string {
	return "sys_logs"
}

func NewLog(db *gorm.DB) *SysLog {
	return &SysLog{connectDb: connectDb{DB: db}}
}

func (m *SysLog) BatchInsertLog(sysLogs []*SysLog) error {
	if len(sysLogs) == 0 {
		return nil
	}
	return m.Table(m.TableName()).CreateInBatches(sysLogs, 150).Error // IGNORE
}
