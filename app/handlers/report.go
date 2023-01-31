package handlers

import (
	"bs.mobgi.cc/app/response"
	servicereport "bs.mobgi.cc/app/service/report"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

type Report struct {
}

func (h *Report) Comprehensive(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VReportComprehensive)
	list, total, err := servicereport.ReportComprehensive(params)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	response.Success(ctx, gin.H{
		"total":   total,
		"list":    list,
		"columns": servicereport.ReportComprehensiveColumns(params.ShowColumns, params.Dimensions),
	})
}
