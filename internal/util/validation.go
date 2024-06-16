package util

import (
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	once                sync.Once
	UniversalTranslator *ut.UniversalTranslator
)

func AddValidation() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	once.Do(func() {
		en := en.New()
		UniversalTranslator = ut.New(en, en)
	})

	trans, _ := UniversalTranslator.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

}
