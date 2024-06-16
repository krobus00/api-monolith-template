package util

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"gorm.io/gorm"
)

var (
	once                sync.Once
	UniversalTranslator *ut.UniversalTranslator
)

func AddValidation(db *gorm.DB) {
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

	// register custom validation
	v.RegisterValidation("unique_db", uniqueValidator(db))
	v.RegisterTranslation("unique_db", trans, func(ut ut.Translator) error {
		return ut.Add("unique_db", "{0} already taken", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique_db", fe.Field())
		return t
	})

	// set tag name from json tag
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

}

// uniqueValidator validate that new value is not exist on db / unique
// example usage: unique_db=users:email
func uniqueValidator(db *gorm.DB) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		param := fl.Param()
		params := strings.Split(param, ":")
		if len(params) != 2 {
			return false
		}

		tableName := params[0]
		columnName := params[1]
		fieldValue := fl.Field().String()
		var count int64

		query := fmt.Sprintf("%s = ?", columnName)
		err := db.Table(tableName).Where(query, fieldValue).Where("deleted_at IS NULL").Count(&count).Error
		if err != nil {
			return false
		}

		return count == 0
	}
}
