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
