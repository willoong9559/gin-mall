package localization

import (
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/willoong9559/gin-mall/conf"
)

var trans ut.Translator

type ValidationErrors struct {
	validator.ValidationErrors
}

func GetValidationErrors(err error) (ValidationErrors, bool) {
	errs, ok := err.(validator.ValidationErrors)
	return ValidationErrors{errs}, ok
}

func (errs ValidationErrors) Translate() map[string]string {
	ValidationErrs := errs.ValidationErrors.Translate(trans)
	return removeTopStruct(ValidationErrs)
}

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.LastIndex(field, ".")+1:]] = err
	}
	return rsp
}

func LoadValidTrans() {
	// get language conf
	lang := conf.ValidatorLang
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取自定义tag的自定义方法
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zh := zh.New()
		en := en.New()
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(en, zh, en)
		trans, ok = uni.GetTranslator(lang)
		log.Println("lang=", lang)
		if !ok {
			log.Printf("get uni.GetTranslator(%s) failed", lang)
			return
		}
		switch lang {
		case "en":
			en_translations.RegisterDefaultTranslations(validate, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(validate, trans)
		default:
			en_translations.RegisterDefaultTranslations(validate, trans)
		}
	}
}
