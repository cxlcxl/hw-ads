package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"errors"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VCountries(ctx *gin.Context) {
	var params v_data.VCountries
	bindData(ctx, &params, (&handlers.Region{}).Countries)
}

func (v BsValidator) VRegionCreate(ctx *gin.Context) {
	var params v_data.VRegionCreate
	bindData(ctx, &params, (&handlers.Region{}).RegionCreate, validRegion)
}
func (v BsValidator) VRegionAreaSet(ctx *gin.Context) {
	var params v_data.VRegionAreaSet
	bindData(ctx, &params, (&handlers.Region{}).RegionAreaSet)
}

func validRegion(_ *gin.Context, p interface{}) error {
	params := p.(*v_data.VRegionCreate)
	if params.T == "area" {
		if params.AreaName == "" {
			return errors.New("地区名称必填")
		}
		return nil
	} else if params.T == "country" {
		if params.CName == "" {
			return errors.New("国家名称必填")
		}
		if params.CCode == "" {
			return errors.New("国家代码必填")
		}
		if params.AreaId <= 0 {
			return errors.New("所属区域必填")
		}
		return nil
	} else {
		return errors.New("非法请求")
	}
}
