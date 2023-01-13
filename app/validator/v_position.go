package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VPositionQuery(ctx *gin.Context) {
	var params v_data.VPositionQuery
	bindData(ctx, &params, emptyValidator, (&handlers.Position{}).PositionQuery)
}

func (v BsValidator) VPositionPlacement(ctx *gin.Context) {
	var params v_data.VPositionPlacement
	bindData(ctx, &params, emptyValidator, (&handlers.Position{}).PositionPlacement)
}

func (v BsValidator) VPositionPrice(ctx *gin.Context) {
	var params v_data.VPositionPrice
	bindData(ctx, &params, emptyValidator, (&handlers.Position{}).PositionPrice)
}
