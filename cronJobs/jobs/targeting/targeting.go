package targeting

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/targeting/logic"
	"fmt"
	"log"
	"time"
)

func Targeting() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= Targeting job start ==================")

	if err := logic.NewTargetingLogic().TargetingQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleTargeting, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= Targeting job end ==================")
	fmt.Println()
	fmt.Println()
}
