package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/job_data/scripts"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
)

type Settings struct{}

func (h *Settings) VersionInfo(ctx *gin.Context) {
	response.Success(ctx, vars.VersionInfo)
}

func (h *Settings) Cron(ctx *gin.Context) {
	jobs, err := model.NewJob(vars.DBMysql).GetJobs()
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	handlerJobs := make([]string, 0)
	for s, _ := range scripts.ManualScheduleJobs {
		handlerJobs = append(handlerJobs, s)
	}
	response.Success(ctx, gin.H{
		"list":         jobs,
		"pause_rules":  vars.JobPauseRule,
		"api_modules":  vars.ApiModules,
		"handler_jobs": handlerJobs,
		"handler_pause_rules": map[int]string{
			0:  "只调度一天",
			1:  "到 1 天前停止",
			2:  "到 2 天前停止",
			3:  "到 3 天前停止",
			4:  "到 4 天前停止",
			5:  "到 5 天前停止",
			6:  "到 6 天前停止",
			7:  "到 7 天前停止",
			99: "使用配置的规则",
		},
	})
}

func (h *Settings) CronInfo(ctx *gin.Context, v string) {
	id, _ := strconv.ParseInt(v, 0, 64)
	if id <= 0 {
		response.Fail(ctx, "请求失败，参数有误")
		return
	}
	job, err := model.NewJob(vars.DBMysql).FindOneById(id)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	response.Success(ctx, job)
}

func (h *Settings) CronUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VSettingsCronUpdate)
	val := map[string]interface{}{
		"job_schedule": params.JobSchedule,
		"order_by":     params.OrderBy,
		"pause_rule":   params.PauseRule,
		"remark":       params.Remark,
		"stat_day":     params.StatDay,
		"version":      params.Version + 1,
	}
	if err := model.NewJob(vars.DBMysql).UpdateJob(params.Id, val); err != nil {
		response.Fail(ctx, "操作失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (h *Settings) CronSchedule(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VSettingsCronSchedule)
	if job, ok := scripts.ManualScheduleJobs[params.ApiModule]; !ok {
		response.Fail(ctx, "调度模块不存在")
		return
	} else {
		t, err := time.Parse(vars.DateFormat, params.StatDay)
		if err != nil {
			response.Fail(ctx, "调度日期参数格式错误")
			return
		}
		go job(t, params.PauseRule)
	}
	response.Success(ctx, nil)
}

func (h *Settings) Configs(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VSettingsConfigs)
	offset := utils.GetPages(params.Page, params.PageSize)
	configs, total, err := model.NewSysConfig(vars.DBMysql).List(params.Key, params.Desc, params.State, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{"total": total, "list": configs})
}

func (h *Settings) Config(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	config, err := model.NewSysConfig(vars.DBMysql).FindOneById(id)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, config)
}

func (h *Settings) ConfigCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VSettingsConfigCreate)
	err := model.NewSysConfig(vars.DBMysql).CreateConfig(model.SysConfig{
		Key:    params.Key,
		Val:    params.Val,
		Desc:   params.Desc,
		State:  1,
		Bak1:   params.Bak1,
		Bak2:   params.Bak2,
		Remark: params.Remark,
	})
	if err != nil {
		response.Fail(ctx, "创建失败: "+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (h *Settings) ConfigUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VSettingsConfigUpdate)
	err := model.NewSysConfig(vars.DBMysql).UpdateConfig(params.Id, map[string]interface{}{
		"_k":     params.Key,
		"_v":     params.Val,
		"_desc":  params.Desc,
		"state":  params.State,
		"bak1":   params.Bak1,
		"bak2":   params.Bak2,
		"remark": params.Remark,
	})
	if err != nil {
		response.Fail(ctx, "修改失败: "+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (h *Settings) ToolLog(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VSettingsLog)
	t, err := time.Parse(vars.DateFormat, params.D)
	if err != nil {
		response.Fail(ctx, "时间格式有误")
		return
	}
	// 每 24 小时切割，数量相当于天数
	d := vars.YmlConfig.GetInt("Logs.MaxBackups")
	if t.Before(time.Now().AddDate(0, 0, -1*d)) {
		response.Fail(ctx, fmt.Sprintf("最多只能查询 %d 天的日志数据", d))
		return
	}
	f := vars.YmlConfig.GetString("Logs.SysLogName")
	filename := fmt.Sprintf("%s.%s", f, t.Format("20060102"))
	if _, err = os.Stat(filename); err != nil {
		response.Fail(ctx, "请求日期的日志不存在")
		return
	}
	key := fmt.Sprintf("log:%s:%s", t.Format("20060102"), utils.GenerateSecret(15))
	if vars.DBRedis.SetString(key, filename, 30) {
		response.Success(ctx, key)
	} else {
		response.Fail(ctx, "日子文件请求失败")
	}
}

func (h *Settings) LogDownload(ctx *gin.Context, key string) {
	v := vars.DBRedis.GetString(key)
	if v == "" {
		response.Fail(ctx, "请求失败，密钥错误")
		return
	}
	keys := strings.Split(key, ":")
	ctx.Header("Content-Type", "application/text")
	ctx.Header("Content-Disposition", "attachment; filename=syslog.log"+keys[1])
	ctx.File(v)
}
