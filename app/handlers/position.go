package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	servicemarketing "bs.mobgi.cc/app/service/marketing"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
)

type Position struct{}

func (h *Position) PositionQuery(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VPositionQuery)
	detail, err := servicemarketing.GetPositionDetail(params.Category, params.ProductType, params.AccountId)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, detail)
}

func (h *Position) PositionPlacement(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VPositionPlacement)
	placements, err := model.NewPositionPlacement(vars.DBMysql).PlacementsByCreativeSizeId(params.CreativeSizeId)
	if err != nil {
		response.Fail(ctx, "没有查到此版位对应的创意子类型："+err.Error())
		return
	}
	rs := make(map[string][]string)
	for _, placement := range placements {
		rs[placement.CreativeSizeSubType] = append(rs[placement.CreativeSizeSubType], placement.CreativeSize)
	}
	response.Success(ctx, gin.H{"sub_types": vars.PositionSubType, "placements": rs})
}

func (h *Position) PositionPrice(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VPositionPrice)
	f, err := model.NewPositionPrice(vars.DBMysql).FindFloorPrice(params.CreativeSizeId, params.PriceType)
	if err != nil {
		response.Fail(ctx, "底价查询失败："+err.Error())
		return
	}
	response.Success(ctx, f)
}
