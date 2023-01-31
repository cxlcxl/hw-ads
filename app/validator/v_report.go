package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VReportComprehensive(ctx *gin.Context) {
	var params v_data.VReportComprehensive
	bindData(ctx, &params, emptyValidator, (&handlers.Report{}).Comprehensive)
}
