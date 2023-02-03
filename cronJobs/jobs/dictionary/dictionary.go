package dictionary

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/dictionary/logic"
	"fmt"
	"log"
	"time"
)

func Dictionary() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= Dictionary job start ==================")

	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleDictionary)
	}()
	if err := logic.NewDictionaryQueryLogic().DictionaryQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleDictionary, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= Dictionary job end ==================")
	fmt.Println()
	fmt.Println()
}
