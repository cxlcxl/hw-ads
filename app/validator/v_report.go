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
	case *v_data.VReportComprehensiveColumn:
		p.(*v_data.VReportComprehensiveColumn).User = u.(*vars.LoginUser)
	case *v_data.VReportComprehensive:
		p.(*v_data.VReportComprehensive).User = u.(*vars.LoginUser)
	default:
	}
	return nil
}

func (v BsValidator) VReportComprehensive(ctx *gin.Context) {
	var params v_data.VReportComprehensive
	bindData(ctx, &params, (&handlers.Report{}).Comprehensive, fillUser)
}

func (v BsValidator) VReportComprehensiveColumn(ctx *gin.Context) {
	var params v_data.VReportComprehensiveColumn
	bindData(ctx, &params, (&handlers.Report{}).ComprehensiveColumn, fillUser)
}
