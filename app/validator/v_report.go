package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
)

func fillUser(ctx *gin.Context, p interface{}) error {
	u, _ := ctx.Get(vars.LoginUserKey)
	switch p.(type) {
	case *v_data.VReportComprehensive:
		p.(*v_data.VReportComprehensive).User = u.(*vars.LoginUser)
	case *v_data.VReportColumn:
		p.(*v_data.VReportColumn).User = u.(*vars.LoginUser)
	case *v_data.VReportAds:
		p.(*v_data.VReportAds).User = u.(*vars.LoginUser)
	default:
	}
	return nil
}

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
