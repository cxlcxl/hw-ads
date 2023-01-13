package refreshToken

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs"
	"fmt"
	"log"
	"time"
)

func RefreshToken() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= RefreshToken job start ==================")

	list, err := model.NewToken(vars.DBMysql).GetAccessTokenList()
	if err != nil {
		fmt.Println("刷新 Token 任务失败，查询 Token 数据失败：", err)
	}
	for _, tokens := range list {
		_, err = jobs.Refresh(tokens)
		if err != nil {
			log.Println("Token 刷新失败，账户 ID：", tokens.AccountId, err)
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
