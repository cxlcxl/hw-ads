package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	servicereport "bs.mobgi.cc/app/service/report"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	comprehensiveColumnKeyPrefix = "comprehensive"
	adsColumnKeyPrefix           = "ads"
)

type Report struct{}

func (h *Report) Comprehensive(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VReportComprehensive)
	if !utils.InArray("country", params.Dimensions) {
		params.Countries = make([][]string, 0)
	}
	if !utils.InArray("account_id", params.Dimensions) {
		params.AccountIds = make([]int64, 0)
	}
	if !utils.InArray("app_id", params.Dimensions) {
		params.AppIds = make([]string, 0)
	}
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
		"total":     total,
		"list":      list,
		"columns":   servicereport.ReportComprehensiveColumns(params.ShowColumns, params.Dimensions),
		"summaries": servicereport.ReportComprehensiveSummaries(params),
	})
}

func (h *Report) Ads(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VReportAds)
	if !utils.InArray("country", params.Dimensions) {
		params.Countries = make([][]string, 0)
	}
	if !utils.InArray("account_id", params.Dimensions) {
		params.AccountIds = make([]int64, 0)
	}
	if !utils.InArray("app_id", params.Dimensions) {
		params.AppIds = make([]string, 0)
	}
	list, total, err := servicereport.ReportAds(params)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	if len(params.ShowColumns) == 0 {
		key := fmt.Sprintf("%s_%d", adsColumnKeyPrefix, params.User.UserId)
		if column, err := model.NewReportColumn(vars.DBMysql).GetColumn(key); err == nil && column.Columns != "" {
			params.ShowColumns = strings.Split(column.Columns, ",")
		}
	}
	response.Success(ctx, gin.H{
		"total":   total,
		"list":    list,
		"columns": servicereport.ReportAdsColumns(params.ShowColumns, params.Dimensions),
	})
}

func (h *Report) Column(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VReportColumn)
	key := fmt.Sprintf("%s_%d", params.Module, params.User.UserId)
	val := map[string]interface{}{"columns": strings.Join(params.Columns, ",")}
	if err := model.NewReportColumn(vars.DBMysql).UpdateColumn(key, val); err != nil {
		response.Fail(ctx, "设置失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}
