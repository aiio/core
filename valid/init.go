package valid

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("in", In)
	}
}

// In 使用例子  in=registry-login
var In validator.Func = func(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if ok {
		split := strings.Split(fl.Param(), "-")
		log.Println(data, split)
		for _, s := range split {
			if strings.EqualFold(s, data) {
				return true
			}
		}
	}
	return false
}
