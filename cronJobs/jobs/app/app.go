package app

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/app/logic"
	"bs.mobgi.cc/library/hlog"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func ReportApp() {
	hlog.NewLog(logrus.InfoLevel, "jobs-app").Log(logrus.Fields{}, "App job start")
	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleApp)
		hlog.NewLog(logrus.InfoLevel, "jobs-app").Log(logrus.Fields{}, "App job end")
	}()
	if err := logic.NewAppQueryLogic().AppQuery(); err != nil {
		hlog.NewLog(logrus.ErrorLevel, "jobs-app-schedule").Log(logrus.Fields{}, err)
		return
	}
	if err := model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleApp, time.Now().Format(vars.DateFormat)); err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-app-update-day").Log(logrus.Fields{}, err)
		return
	}

	fmt.Println()
	fmt.Println()
}

func ReportAppManual(_ time.Time, _ int64) {
	ReportApp()
}
