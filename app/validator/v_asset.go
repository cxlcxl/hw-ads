package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VAssetList(ctx *gin.Context) {
	var params v_data.VAssetList
	bindData(ctx, &params, emptyValidator, (&handlers.Asset{}).AssetList)
}

func (v BsValidator) VAssetSync(ctx *gin.Context) {
	var params v_data.VAssetSync
	bindData(ctx, &params, emptyValidator, (&handlers.Asset{}).AssetSync)
}
