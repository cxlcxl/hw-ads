package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Job struct {
	connectDb
	Id          int64     `json:"id"`
	StatDay     time.Time `json:"stat_day"`     // 数据日期
	ApiModule   string    `json:"api_module"`   // 数据模块
	TotalPage   int64     `json:"total_page"`   // 总页数
	TotalNum    int64     `json:"total_num"`    // 总数
	JobSchedule string    `json:"job_schedule"` // cron 调度
	PauseRule   int64     `json:"pause_rule"`   // 调度截止规则：0 调度到当天；-1 停止调度此任务；> 0 为当前日期减{pause_rule}天；
	Version     int64     `json:"version"`      // 版本：每次有规则或调度修改 +1
}

func (m *Job) TableName() string {
	return "jobs"
}

func NewJob(db *gorm.DB) *Job {
	return &Job{connectDb: connectDb{DB: db}}
}

func (m *Job) GetJobs() (jobs []*Job, err error) {
	err = m.Table(m.TableName()).Find(&jobs).Error
	return
}

func (m *Job) FindOneByApiModule(module string) (job *Job, err error) {
	err = m.Table(m.TableName()).Where("`api_module` = ?", module).First(&job).Error
	return
}

func (m *Job) UpdateJobDayByModule(module, day string) (err error) {
	query := fmt.Sprintf("update %s set stat_day = ? where `api_module` = ?", m.TableName())
	err = m.Exec(query, day, module).Error
	return
}
