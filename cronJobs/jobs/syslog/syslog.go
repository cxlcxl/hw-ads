package syslog

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/syslog/logic"
	"bs.mobgi.cc/library/hlog"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func SysLog() {
	hlog.NewLog(logrus.InfoLevel, "jobs-log").Log(logrus.Fields{}, "Log job start")
	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleLog)
	}()
	job, err := model.NewJob(vars.DBMysql).FindOneByApiModule(vars.ApiModuleLog)
	if err != nil {
		hlog.NewLog(logrus.ErrorLevel, "jobs-log-module").Log(logrus.Fields{}, err)
		return
	}
	jobDay := job.StatDay
	now := time.Now()
	for {
		pauseDay := now.AddDate(0, 0, -1*int(job.PauseRule))
		if jobDay.After(pauseDay) {
			break
		}
		d := jobDay.Format(vars.DateFormat)
		if err = logic.NewLogLogic(d).Parse(); err != nil {
			hlog.NewLog(logrus.ErrorLevel, "jobs-log").Log(logrus.Fields{
				"stat_day": d,
			}, err)
			return
		}
		if err = model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleLog, jobDay.Format(vars.DateFormat)); err != nil {
			hlog.NewLog(logrus.WarnLevel, "jobs-log-update-day").Log(logrus.Fields{}, err)
		}
		jobDay = jobDay.AddDate(0, 0, 1)
		time.Sleep(time.Millisecond * 500)
	}

	hlog.NewLog(logrus.InfoLevel, "jobs-log").Log(logrus.Fields{}, "Log job end")
	fmt.Println()
	fmt.Println()
}

func SysLogManual(d time.Time, _ int64) {
	// 最晚只能调度到前一天
	if d.After(time.Now().AddDate(0, 0, -1)) {
		fmt.Println("最多调度到当前时间的前一天")
		return
	}
	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleLog)
	}()
	if err := logic.NewLogLogic(d.Format(vars.DateFormat)).Parse(); err != nil {
		hlog.NewLog(logrus.ErrorLevel, "jobs-log").Log(logrus.Fields{}, err)
		return
	}
}
