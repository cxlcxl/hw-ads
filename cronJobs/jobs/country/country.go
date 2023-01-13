package country

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/country/logic"
	"fmt"
	"log"
	"time"
)

func Country() {
	fmt.Println("================= country job start ==================")

	job, err := model.NewJob(vars.DBMysql).FindOneByApiModule(vars.ApiModuleCountry)
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
		fmt.Println("schedule day: ", d)
		if err = logic.NewCountryQueryLogic(d).CountryQuery(); err != nil {
			log.Fatal(err)
			return
		}
		jobDay = jobDay.AddDate(0, 0, 1)
		if err = model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleCountry, jobDay.Format(vars.DateFormat)); err != nil {
			fmt.Println("数据库调度时间修改失败: ", err)
		}
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("================= country job end ==================")
	fmt.Println()
	fmt.Println()
}

func CountryManual(day time.Time, pauseRule int64) {
	// 手动调度不改脚本自动调度的日期
	fmt.Println("================= country job start ==================")

	now := time.Now()
	if pauseRule == 0 {
		if err := countryDoSchedule(day.Format(vars.DateFormat)); err != nil {
			fmt.Println("调度失败：", err)
		}
	} else if pauseRule == 99 {
		job, err := model.NewJob(vars.DBMysql).FindOneByApiModule(vars.ApiModuleCountry)
		if err != nil {
			log.Fatal("调度模块查询错误：", err)
			return
		}
		jobDay := job.StatDay
		for {
			pauseDay := now.AddDate(0, 0, -1*int(job.PauseRule))
			if jobDay.After(pauseDay) {
				break
			}
			d := jobDay.Format(vars.DateFormat)
			fmt.Println("schedule day: ", d)
			if err = logic.NewCountryQueryLogic(d).CountryQuery(); err != nil {
				log.Fatal(err)
				return
			}
			jobDay = jobDay.AddDate(0, 0, 1)

			time.Sleep(time.Millisecond * 500)
		}
	} else if pauseRule >= 1 && pauseRule <= 31 {
		pauseDay := now.AddDate(0, 0, -1*int(pauseRule))
		for {
			if day.After(pauseDay) {
				break
			}
			d := day.Format(vars.DateFormat)
			if err := logic.NewCountryQueryLogic(d).CountryQuery(); err != nil {
				log.Fatal(err)
				return
			}
			day = day.AddDate(0, 0, 1)
			fmt.Println(d, day.Format(vars.DateFormat))

			time.Sleep(time.Millisecond * 500)
		}
	} else {
		fmt.Println("规则有误，可使用 -h 查看")
	}

	fmt.Println("================= country job end ==================")
	fmt.Println()
	fmt.Println()
}

func countryDoSchedule(d string) (err error) {
	if err = logic.NewCountryQueryLogic(d).CountryQuery(); err != nil {
		log.Fatal(err)
		return
	}
	return nil
}
