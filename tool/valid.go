package tool

import (
	"github.com/go-playground/validator/v10"
)

const maxLen = 15

//限制密码和用户名的验证器
var LengthOk validator.Func = func(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if len(data) > maxLen {
		return false
	}
	return true
}
