package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VTargetingList(ctx *gin.Context) {
	var params v_data.VTargetingList
	bindData(ctx, &params, emptyValidator, (&handlers.Targeting{}).List)
}

func (v BsValidator) VTargetingCreate(ctx *gin.Context) {
	var params v_data.VTargetingCreate
	bindData(ctx, &params, emptyValidator, (&handlers.Targeting{}).Create)
}
