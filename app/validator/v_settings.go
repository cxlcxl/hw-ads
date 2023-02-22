package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (v BsValidator) VSettingsCronInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Settings{}).CronInfo)
}

func (v BsValidator) VSettingsConfigs(ctx *gin.Context) {
	var p v_data.VSettingsConfigs
	bindData(ctx, &p, (&handlers.Settings{}).Configs)
}

func (v BsValidator) VSettingsConfigCreate(ctx *gin.Context) {
	var p v_data.VSettingsConfigCreate
	bindData(ctx, &p, (&handlers.Settings{}).ConfigCreate)
}

func (v BsValidator) VSettingsConfig(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Settings{}).Config)
}

func (v BsValidator) VSettingsConfigUpdate(ctx *gin.Context) {
	var p v_data.VSettingsConfigUpdate
	bindData(ctx, &p, (&handlers.Settings{}).ConfigUpdate)
}

func (v BsValidator) VSettingsCronUpdate(ctx *gin.Context) {
	var p v_data.VSettingsCronUpdate
	bindData(ctx, &p, (&handlers.Settings{}).CronUpdate, bindId)
}

func (v BsValidator) VSettingsCronSchedule(ctx *gin.Context) {
	var p v_data.VSettingsCronSchedule
	bindData(ctx, &p, (&handlers.Settings{}).CronSchedule)
}

func bindId(ctx *gin.Context, p interface{}) error {
	v, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		return err
	}
	if v <= 0 {
		return errors.New("参数错误")
	}
	p.(*v_data.VSettingsCronUpdate).Id = v
	return nil
}

func (v BsValidator) VSettingsLog(ctx *gin.Context) {
	var p v_data.VSettingsLog
	bindData(ctx, &p, (&handlers.Settings{}).ToolLog)
}

func (v BsValidator) VSettingsLogDownload(ctx *gin.Context) {
	bindRouteData(ctx, "key", (&handlers.Settings{}).LogDownload)
}
