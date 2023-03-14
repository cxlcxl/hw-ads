package ads

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/ads/logic"
	"bs.mobgi.cc/library/hlog"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func ReportAds() {
	hlog.NewLog(logrus.InfoLevel, "jobs-ads").Log(logrus.Fields{}, "Ads job start")
	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleAds)
	}()
	job, err := model.NewJob(vars.DBMysql).FindOneByApiModule(vars.ApiModuleAds)
	if err != nil {
		hlog.NewLog(logrus.WarnLevel, "jobs-ads-module").Log(logrus.Fields{}, "Ads job start")
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
		if err = logic.NewAdsQueryLogic(d, 0).AdsQuery(); err != nil {
			hlog.NewLog(logrus.ErrorLevel, "jobs-ads-schedule").Log(logrus.Fields{
				"stat_day": d,
			}, err)
			return
		}
		if err = logic.NewAdsCollectLogic(d, 0).AdsCollect(); err != nil {
			hlog.NewLog(logrus.ErrorLevel, "jobs-ads-collect").Log(logrus.Fields{
				"stat_day": d,
			}, err)
			return
		}
		if err = model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleAds, jobDay.Format(vars.DateFormat)); err != nil {
			hlog.NewLog(logrus.WarnLevel, "jobs-ads-update-day").Log(logrus.Fields{}, err)
		}
		jobDay = jobDay.AddDate(0, 0, 1)
		time.Sleep(time.Millisecond * 500)
	}

	hlog.NewLog(logrus.InfoLevel, "jobs-ads").Log(logrus.Fields{}, "Ads job end")
	fmt.Println()
	fmt.Println()
}

func ReportAdsManual(day time.Time, pauseRule, accountId int64) {
	// 手动调度不改脚本自动调度的日期
	fmt.Println("================= Ads job start ==================")
	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleAds)
	}()
	now := time.Now()
	if pauseRule == 0 {
		if err := adsDoSchedule(day.Format(vars.DateFormat), accountId); err != nil {
			fmt.Println("调度失败：", err)
		}
	} else if pauseRule == 99 {
		job, err := model.NewJob(vars.DBMysql).FindOneByApiModule(vars.ApiModuleAds)
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
			if err = logic.NewAdsQueryLogic(d, accountId).AdsQuery(); err != nil {
				log.Fatal(err)
				return
			}
			if err = logic.NewAdsCollectLogic(d, accountId).AdsCollect(); err != nil {
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
			if err := logic.NewAdsQueryLogic(d, accountId).AdsQuery(); err != nil {
				log.Fatal(err)
				return
			}
			if err := logic.NewAdsCollectLogic(d, accountId).AdsCollect(); err != nil {
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

	fmt.Println("================= Ads job end ==================")
	fmt.Println()
	fmt.Println()
}

func adsDoSchedule(d string, accountId int64) (err error) {
	if err = logic.NewAdsQueryLogic(d, accountId).AdsQuery(); err != nil {
		log.Fatal(err)
		return
	}
	if err = logic.NewAdsCollectLogic(d, accountId).AdsCollect(); err != nil {
		log.Fatal(err)
		return
	}
	return nil
}

func ReportAdsCollectManual(day time.Time, _, accountId int64) {
	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleAdsCollect)
	}()
	if err := logic.NewAdsCollectLogic(day.Format(vars.DateFormat), accountId).AdsCollect(); err != nil {
		log.Fatal(err)
	}
	return
}
