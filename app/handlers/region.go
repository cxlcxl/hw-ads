package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"time"
)

type Region struct{}

func treeRegion(regions []*model.Region, pid int64) (rs []*model.Region) {
	for _, region := range regions {
		if region.Pid == pid {
			region.Children = treeRegion(regions, region.Id)
			rs = append(rs, region)
		}
	}
	return
}

func (h *Region) Regions(ctx *gin.Context) {
	regions, err := model.NewOverseasArea(vars.DBMysql).GetRegions()
	if err != nil {
		response.Fail(ctx, "地区请求失败")
		return
	}
	response.Success(ctx, treeRegion(regions, 0))
}

func (h *Region) Areas(ctx *gin.Context) {
	areas, err := model.NewOverseasArea(vars.DBMysql).Areas()
	if err != nil {
		response.Fail(ctx, "地区请求失败")
		return
	}
	response.Success(ctx, areas)
}

func (h *Region) Countries(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VCountries)
	offset := utils.GetPages(params.Page, params.PageSize)
	countries, total, err := model.NewOverseasRegion(vars.DBMysql).FindCountries(params.AreaId, params.K, int(offset), int(params.PageSize))
	if err != nil {
		response.Fail(ctx, "地区请求失败")
		return
	}
	response.Success(ctx, gin.H{"total": total, "list": countries})
}

func (h *Region) RegionCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VRegionCreate)
	var err error
	if params.T == "area" {
		err = model.NewOverseasArea(vars.DBMysql).AreaCreate(&model.OverseasArea{Name: params.AreaName})
	} else {
		err = model.NewOverseasRegion(vars.DBMysql).RegionCreate(&model.OverseasRegion{
			CId:         utils.MD5(time.Now().GoString()),
			CCode:       params.CCode,
			CName:       params.CName,
			Level:       0,
			ContinentId: 0,
		}, params.AreaIds)
	}
	if err != nil {
		response.Fail(ctx, "地区请求失败")
		return
	}
	response.Success(ctx, nil)
}

func (h *Region) RegionAreaSet(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VRegionAreaSet)
	var areaCountries []*model.OverseasAreaRegion
	for _, id := range params.AreaIds {
		areaCountries = append(areaCountries, &model.OverseasAreaRegion{AreaId: id, CCode: params.CCode})
	}
	err := model.NewOverseasAreaRegion(vars.DBMysql).AreaSet(areaCountries, params.CCode)
	if err != nil {
		response.Fail(ctx, "地区设置失败")
		return
	}
	response.Success(ctx, nil)
}
