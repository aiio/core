package valid

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gogf/gf/util/gconv"
)

var Trans ut.Translator

func init() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //中文翻译器
		uni := ut.New(zhT)
		Trans, _ = uni.GetTranslator("zh")

		err := zhTranslations.RegisterDefaultTranslations(v, Trans)
		if err != nil {
			panic(err)
		}
		err = CoreTranslations(v, Trans)
		if err != nil {

			panic(err)
		}
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		_ = v.RegisterValidation("in", func(fl validator.FieldLevel) bool {
			data := gconv.String(fl.Field().Interface())
			for _, s := range strings.Split(fl.Param(), "#") {
				if strings.EqualFold(s, data) {
					return true
				}
			}
			return false
		})
	}
}
