package error_catch

import (
	"bs.mobgi.cc/app/validator"
	"fmt"
)

// ValidateErr Key: 'LoginData.Email' Error:Field validation for 'Email' failed on the 'email' tag"
func ValidateErr(err error, prefix string) string {
	return fmt.Sprintf("%s：%s", prefix, validator.Translate(err))
}
