package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VReportComprehensive(ctx *gin.Context) {
	var params v_data.VReportComprehensive
	bindData(ctx, &params, (&handlers.Report{}).Comprehensive, fillUser)
}

func (v BsValidator) VReportAds(ctx *gin.Context) {
	var params v_data.VReportAds
	bindData(ctx, &params, (&handlers.Report{}).Ads, fillUser)
}

func (v BsValidator) VReportColumn(ctx *gin.Context) {
	var params v_data.VReportColumn
	bindData(ctx, &params, (&handlers.Report{}).Column, fillUser)
}
