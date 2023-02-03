package app

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/app/logic"
	"fmt"
	"log"
	"time"
)

func ReportApp() {
	fmt.Println("================= App job start ==================")
	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleApp)
	}()
	if err := logic.NewAppQueryLogic().AppQuery(); err != nil {
		log.Fatal(err)
		return
	}
	if err := model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleApp, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= App job end ==================")
	fmt.Println()
	fmt.Println()
}

func ReportAppManual(_ time.Time, _ int64) {
	ReportApp()
}
