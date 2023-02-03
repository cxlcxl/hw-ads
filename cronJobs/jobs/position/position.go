package position

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/position/logic"
	"fmt"
	"log"
	"time"
)

func Position() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= Position job start ==================")

	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModulePosition)
	}()
	if err := logic.NewPositionLogic().PositionQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModulePosition, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= Position job end ==================")
	fmt.Println()
	fmt.Println()
}

func PriceManual(_ time.Time, _ int64) {
	Price()
}
func Price() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= PositionPrice job start ==================")

	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModulePositionPrice)
	}()
	if err := logic.NewPositionPriceLogic().PriceQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModulePositionPrice, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= PositionPrice job end ==================")
	fmt.Println()
	fmt.Println()
}

func ElementManual(_ time.Time, _ int64) {
	Element()
}

// Element 版位元素
func Element() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= PositionElement job start ==================")

	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModulePositionElement)
	}()
	if err := logic.NewPositionElementLogic().ElementQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModulePositionElement, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= PositionElement job end ==================")
	fmt.Println()
	fmt.Println()
}
