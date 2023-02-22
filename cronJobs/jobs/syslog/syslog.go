package syslog

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/syslog/logic"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func SysLog() {
	fmt.Println("================= Log job start ==================")
	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleLog)
	}()
	job, err := model.NewJob(vars.DBMysql).FindOneByApiModule(vars.ApiModuleLog)
	if err != nil {
		log.Fatal("调度模块查询错误：", err)
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
			vars.HLog.WithFields(logrus.Fields{
				"module": "jobs-log",
				"log_id": time.Now().UnixNano(),
			}).Error(err)
			return
		}
		if err = model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleLog, jobDay.Format(vars.DateFormat)); err != nil {
			fmt.Println("数据库调度时间修改失败: ", err)
		}
		jobDay = jobDay.AddDate(0, 0, 1)
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("================= Ads job end ==================")
	fmt.Println()
	fmt.Println()
}

func SysLogManual(d time.Time, _ int64) {
	// 最晚只能调度到前一天
	if d.After(time.Now().AddDate(0, 0, -1)) {
		fmt.Println("最多调度到当前时间的前一天")
		return
	}
	if err := logic.NewLogLogic(d.Format(vars.DateFormat)).Parse(); err != nil {
		vars.HLog.WithFields(logrus.Fields{
			"module": "jobs-log",
			"log_id": time.Now().UnixNano(),
		}).Error(err)
		return
	}
}
