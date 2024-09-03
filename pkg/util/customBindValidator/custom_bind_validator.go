// @Title custom_bind_validator.go
// @Description
// @Author Hunter 2024/9/3 18:18

package customBindValidator

import (
	"github.com/go-playground/validator/v10"
	"go-gin-api-starter/pkg/util/formatValidator"
)

var isMobileNumber validator.Func = func(fl validator.FieldLevel) bool {
	text := fl.Field().String()
	if err := formatValidator.ValidateMobileNumber(text); err != nil {
		return false
	}
	return true
}

var isEmail validator.Func = func(fl validator.FieldLevel) bool {
	text := fl.Field().String()
	if err := formatValidator.ValidateEmail(text); err != nil {
		return false
	}
	return true
}

var isAccountName validator.Func = func(fl validator.FieldLevel) bool {
	text := fl.Field().String()
	if err := formatValidator.ValidateAccountName(text); err != nil {
		return false
	}
	return true
}
