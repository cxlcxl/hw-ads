package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
)

type Common struct{}

func (h *Common) Regions(ctx *gin.Context) {
	regions, err := model.NewOverseasArea(vars.DBMysql).GetRegions()
	if err != nil {
		response.Fail(ctx, "地区请求失败")
		return
	}
	response.Success(ctx, treeRegion(regions, 0))
}

func treeRegion(regions []*model.Region, pid int64) (rs []*model.Region) {
	for _, region := range regions {
		if region.Pid == pid {
			region.Children = treeRegion(regions, region.Id)
			rs = append(rs, region)
		}
	}
	return
}
