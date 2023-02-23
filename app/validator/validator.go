package validator

import (
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/vars"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

const (
	passRule = `^[a-zA-Z]+[a-zA-Z0-9\.\@\#\$\%\&\*\!\?\,]{5,17}$`
)

type BsValidator struct{}

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("pass", pass)
		_ = v.RegisterValidation("dimensions", dimensions)
	}
}

func pass(fl validator.FieldLevel) bool {
	_pass := fl.Field().String()
	if ok, err := regexp.MatchString(passRule, _pass); err != nil {
		return false
	} else {
		return ok
	}
}

func dimensions(fl validator.FieldLevel) bool {
	_dimensions := fl.Field().Interface().([]string)
	for _, dimension := range _dimensions {
		if _, ok := vars.ReportDimensions[dimension]; !ok {
			return false
		}
	}
	return true
}

// ctx 上下文
// v   要绑定的数据
// h   绑定完成后调用的方法
// f   自定义扩展验证规则
func bindData(ctx *gin.Context, v interface{}, h func(*gin.Context, interface{}), fs ...func(*gin.Context, interface{}) error) {
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
	if _, ok := reflect.TypeOf(p).Elem().FieldByName("User"); !ok {
		return errors.New("用户信息绑定失败，请检查是否包含 User 结构体")
	} else {
		u, _ := ctx.Get(vars.LoginUserKey)
		reflect.ValueOf(p).Elem().FieldByName("User").Set(reflect.ValueOf(u))
	}
	return nil
}
