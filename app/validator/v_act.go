package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ParentCanNotSelf = errors.New("上级服务商不可是本账户")
)

func (v BsValidator) VAccountList(ctx *gin.Context) {
	var params v_data.VAccountList
	bindData(ctx, &params, (&handlers.Account{}).AccountList, fillUser)
}

func (v BsValidator) VAccountParents(ctx *gin.Context) {
	var params v_data.VAccountParents
	bindData(ctx, &params, (&handlers.Account{}).AccountParents)
}

func (v BsValidator) VAccountSearch(ctx *gin.Context) {
	var params v_data.VAccountSearch
	bindData(ctx, &params, (&handlers.Account{}).AccountSearch)
}

func (v BsValidator) VAccountDefault(ctx *gin.Context) {
	(&handlers.Account{}).AccountDefault(ctx)
}

func (v BsValidator) VAllAccounts(ctx *gin.Context) {
	(&handlers.Account{}).AllAccounts(ctx)
}

func (v BsValidator) VAccountCreate(ctx *gin.Context) {
	var params v_data.VAccountCreate
	bindData(ctx, &params, (&handlers.Account{}).AccountCreate)
}

func (v BsValidator) VAccountAuth(ctx *gin.Context) {
	var params v_data.VAccountAuth
	bindData(ctx, &params, (&handlers.Account{}).AccountAuthToken)
}

func (v BsValidator) VAccountRefreshToken(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Account{}).RefreshToken)
}

func (v BsValidator) VAccountUpdate(ctx *gin.Context) {
	var params v_data.VAccountUpdate
	bindData(ctx, &params, (&handlers.Account{}).AccountUpdate, func(_ *gin.Context, v interface{}) error {
		p := v.(*v_data.VAccountUpdate)
		if p.Id == p.ParentId {
			return ParentCanNotSelf
		}
		return nil
	})
}

func (v BsValidator) VAccountInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Account{}).AccountInfo)
}
