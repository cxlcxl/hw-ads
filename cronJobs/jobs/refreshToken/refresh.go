package refreshToken

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs"
	"bs.mobgi.cc/library/hlog"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func RefreshToken() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= RefreshToken job start ==================")

	defer func() {
		_ = model.NewJob(vars.DBMysql).UpdateLastSchedule(vars.ApiModuleRefreshToken)
	}()
	list, err := model.NewToken(vars.DBMysql).GetAccessTokenList()
	if err != nil {
		fmt.Println("刷新 Token 任务失败，查询 Token 数据失败：", err)
	}
	for _, tokens := range list {
		_, err = jobs.Refresh(tokens)
		if err != nil {
			hlog.NewLog(logrus.ErrorLevel, "jobs-refreshToken").Log(logrus.Fields{"account_id": tokens.AccountId}, err)
			continue
		}
	}

	if err = model.NewJob(vars.DBMysql).UpdateJobDayByModule(vars.ApiModuleRefreshToken, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= RefreshToken job end ==================")
	fmt.Println()
	fmt.Println()
}

func RefreshTokenManual(_ time.Time, _ int64) {
	RefreshToken()
}
