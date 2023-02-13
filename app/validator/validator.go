package validator

import (
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	appCodeRule = `^[a-z0-9A-Z\_]{1,50}$`
	ticketRule  = `^[a-z0-9]+$`
	uuidRule    = `^[a-z0-9\-]{36}$`
)

type BsValidator struct{}

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("pass", pass)
	}
}

func pass(fl validator.FieldLevel) bool {
	//if fl.Field().String() == "invalid" {
	//	return false
	//}
	return true
}

func emptyValidator(_ *gin.Context, _ interface{}) error {
	return nil
}

// ctx 上下文
// v   要绑定的数据
// h   绑定完成后调用的方法
// f   自定义扩展验证规则
func bindData(ctx *gin.Context, v interface{}, h func(*gin.Context, interface{}), fs ...func(*gin.Context, interface{}) error) {
	//if err := ctx.ShouldBindBodyWith(v, binding.JSON); err != nil {
	if err := ctx.ShouldBind(v); err != nil {
		response.Fail(ctx, "验证失败："+Translate(err))
		return
	}
	for _, f := range fs {
		if err := f(ctx, v); err != nil {
			response.Fail(ctx, "验证失败："+err.Error())
			return
		}
	}

	h(ctx, v)
}

func bindRouteData(ctx *gin.Context, key string, h func(c *gin.Context, t string)) {
	h(ctx, ctx.Param(key))
}

func fillUser(ctx *gin.Context, p interface{}) error {
	u, _ := ctx.Get(vars.LoginUserKey)
	switch p.(type) {
	case *v_data.VReportComprehensive:
		p.(*v_data.VReportComprehensive).User = u.(*vars.LoginUser)
	case *v_data.VReportColumn:
		p.(*v_data.VReportColumn).User = u.(*vars.LoginUser)
	case *v_data.VReportAds:
		p.(*v_data.VReportAds).User = u.(*vars.LoginUser)
	case *v_data.VSelfUpdate:
		p.(*v_data.VSelfUpdate).User = u.(*vars.LoginUser)
	default:
	}
	return nil
}
