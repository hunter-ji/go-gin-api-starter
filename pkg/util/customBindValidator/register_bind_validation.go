// @Title register_bind_validation.go
// @Description
// @Author Hunter 2024/9/3 18:20

package customBindValidator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Register() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("isPhoneNumber", isMobileNumber)
		if err != nil {
			return
		}

		err = v.RegisterValidation("isEmail", isEmail)
		if err != nil {
			return
		}

		err = v.RegisterValidation("isAccountName", isAccountName)
		if err != nil {
			return
		}
	}

	return
}
