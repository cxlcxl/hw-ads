package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	servicereport "bs.mobgi.cc/app/service/report"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	comprehensiveColumnKeyPrefix = "comprehensive"
)

type Report struct{}

func (h *Report) Comprehensive(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VReportComprehensive)
	list, total, err := servicereport.ReportComprehensive(params)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	if len(params.ShowColumns) == 0 {
		key := fmt.Sprintf("%s_%d", comprehensiveColumnKeyPrefix, params.User.UserId)
		if column, err := model.NewReportColumn(vars.DBMysql).GetColumn(key); err == nil && column.Columns != "" {
			params.ShowColumns = strings.Split(column.Columns, ",")
		}
	}
	response.Success(ctx, gin.H{
		"total":   total,
		"list":    list,
		"columns": servicereport.ReportComprehensiveColumns(params.ShowColumns, params.Dimensions),
	})
}

func (h *Report) ComprehensiveColumn(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VReportComprehensiveColumn)
	key := fmt.Sprintf("%s_%d", comprehensiveColumnKeyPrefix, params.User.UserId)
	val := map[string]interface{}{"columns": strings.Join(params.Columns, ",")}
	if err := model.NewReportColumn(vars.DBMysql).UpdateColumn(key, val); err != nil {
		response.Fail(ctx, "设置失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}
