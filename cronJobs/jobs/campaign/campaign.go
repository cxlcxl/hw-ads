package campaign

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/campaign/logic"
	"fmt"
	"log"
	"time"
)

func Campaign() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= campaign job start ==================")

	defer func() {
		fmt.Println("================= campaign job end ==================")
		fmt.Println()
		fmt.Println()
	}()

	job, err := model.NewJob(vars.DBMysql).FindOneByApiModule(vars.ApiModuleCampaign)
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
		if err = logic.NewCampaignQueryLogic(d).CampaignQuery(); err != nil {
			log.Fatal(err)
			return
		}
		jobDay = jobDay.AddDate(0, 0, 1)
		if err = model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleCampaign, jobDay.Format(vars.DateFormat)); err != nil {
			fmt.Println("数据库调度时间修改失败: ", err)
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func CampaignManual(day time.Time, pauseRule int64) {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= campaign job start ==================")
	defer func() {
		fmt.Println("================= campaign job end ==================")
		fmt.Println()
		fmt.Println()
	}()

	now := time.Now()
	if pauseRule == 0 {
		if err := campaignDoSchedule(day.Format(vars.DateFormat), day); err != nil {
			fmt.Println("调度失败：", err)
		}
	} else if pauseRule == 99 {
		job, err := model.NewJob(vars.DBMysql).FindOneByApiModule(vars.ApiModuleCampaign)
		if err != nil {
			log.Fatal("调度模块查询错误：", err)
			return
		}
		if job.PauseRule == -1 {
			fmt.Println("原规则此模块已停止调度：", err)
			return
		}
		pauseDay := now.AddDate(0, 0, -1*int(job.PauseRule))
		for {
			if day.After(pauseDay) {
				break
			}
			d := day.Format(vars.DateFormat)
			if err = logic.NewCampaignQueryLogic(d).CampaignQuery(); err != nil {
				log.Fatal(err)
				return
			}
			day = day.AddDate(0, 0, 1)
			fmt.Println(d, day.Format(vars.DateFormat))
			if err = model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleCampaign, day.Format(vars.DateFormat)); err != nil {
				fmt.Println("数据库调度时间修改失败: ", err)
			}
			time.Sleep(time.Millisecond * 500)
		}
	} else if pauseRule >= 1 && pauseRule <= 31 {
		pauseDay := now.AddDate(0, 0, -1*int(pauseRule))
		for {
			if day.After(pauseDay) {
				break
			}
			d := day.Format(vars.DateFormat)
			if err := logic.NewCampaignQueryLogic(d).CampaignQuery(); err != nil {
				log.Fatal(err)
				return
			}
			day = day.AddDate(0, 0, 1)
			fmt.Println(d, day.Format(vars.DateFormat))
			if err := model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleCampaign, day.Format(vars.DateFormat)); err != nil {
				fmt.Println("数据库调度时间修改失败: ", err)
			}
			time.Sleep(time.Millisecond * 500)
		}
	} else {
		log.Fatal("规则有误，可使用 -h 查看")
	}
}

func campaignDoSchedule(d string, jobDay time.Time) (err error) {
	if err = logic.NewCampaignQueryLogic(d).CampaignQuery(); err != nil {
		log.Fatal(err)
		return
	}
	jobDay = jobDay.AddDate(0, 0, 1)
	if err = model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleCampaign, jobDay.Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}
	return err
}
