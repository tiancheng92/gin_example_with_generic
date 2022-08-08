package validator

import (
	"gin_example_with_generic/config"
	"gin_example_with_generic/pkg/ecode"
	"gin_example_with_generic/pkg/errors"
	"gin_example_with_generic/pkg/log"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	jaTranslations "github.com/go-playground/validator/v10/translations/ja"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/tiancheng92/gf"
	"reflect"
	"strings"
	"time"
)

var translator ut.Translator

type validatorInfo struct {
	Tag           string
	TagDetail     string
	Msg           string
	TranslateFunc func(ut.Translator, validator.FieldError) string
	ValidatorFunc func(validator.FieldLevel) bool
}

func dateValidatorFunc(fieldLevel validator.FieldLevel) bool {
	date := fieldLevel.Field().String()
	if date == "" {
		return true
	}
	if _, err := time.ParseInLocation("2006-01", date, time.Local); err != nil {
		return false
	}
	return true

}

func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, true); err != nil {
			return err
		}
		return nil
	}
}

func translateFunc(translator ut.Translator, fieldError validator.FieldError) string {
	msg, err := translator.T(fieldError.Tag(), fieldError.Field())
	if err != nil {
		log.Errorf("register validation failed: %#v", err)
	}
	return msg
}

func translateFuncWithParam(translator ut.Translator, fieldError validator.FieldError) string {
	tag := fieldError.Tag()
	kind := fieldError.Kind()
	if kind == reflect.Ptr {
		kind = fieldError.Type().Elem().Kind()
	}

	switch kind {
	case reflect.String:
		tag += "-string"
	case reflect.Slice, reflect.Map, reflect.Array:
		tag += "-item"
	default:
		tag += "-number"
	}
	msg, err := translator.T(tag, fieldError.Field(), fieldError.Param())
	if err != nil {
		log.Errorf("register validation failed: %#v", err)
	}
	return msg
}

func initValidator() {
	locale := config.GetConf().I18n
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatal("binding validate engine failed")
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		describe := fld.Tag.Get(gf.StringJoin("describe_", locale))
		if describe != "" {
			return describe
		}
		return fld.Name
	})

	translator, _ = ut.New(en.New(), zh.New(), ja.New()).GetTranslator(locale)

	var validatorInfoList []validatorInfo
	switch locale {
	case "en":
		if err := enTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
			log.Fatalf("register default translations failed: %+v", err)
		}

		validatorInfoList = []validatorInfo{
			{Tag: "date", TagDetail: "date", Msg: "{0} invalid format (yyyy-mm)", TranslateFunc: translateFunc, ValidatorFunc: dateValidatorFunc},
		}
	case "zh":
		if err := zhTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
			log.Fatalf("register default translations failed: %+v", err)
		}

		validatorInfoList = []validatorInfo{
			{Tag: "required", TagDetail: "required", Msg: "{0}为必填项", TranslateFunc: translateFunc},
			{Tag: "email", TagDetail: "email", Msg: "{0}格式错误", TranslateFunc: translateFunc},
			{Tag: "max", TagDetail: "max-string", Msg: "{0}不能大于{1}个字符", TranslateFunc: translateFuncWithParam},
			{Tag: "max", TagDetail: "max-item", Msg: "{0}最多包含{1}个对象", TranslateFunc: translateFuncWithParam},
			{Tag: "max", TagDetail: "max-number", Msg: "{0}最大值为{1}", TranslateFunc: translateFuncWithParam},
			{Tag: "min", TagDetail: "min-string", Msg: "{0}不能小于{1}个字符", TranslateFunc: translateFuncWithParam},
			{Tag: "min", TagDetail: "min-item", Msg: "{0}最少包含{1}个对象", TranslateFunc: translateFuncWithParam},
			{Tag: "min", TagDetail: "min-number", Msg: "{0}最小值为{1}", TranslateFunc: translateFuncWithParam},
			{Tag: "date", TagDetail: "date", Msg: "{0}格式错误(yyyy-mm)", TranslateFunc: translateFunc, ValidatorFunc: dateValidatorFunc},
		}
	case "ja":
		if err := jaTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
			log.Fatalf("register default translations failed: %+v", err)
		}

		validatorInfoList = []validatorInfo{
			{Tag: "date", TagDetail: "date", Msg: "{0}無効フォーマット(yyyy-mm)", TranslateFunc: translateFunc, ValidatorFunc: dateValidatorFunc},
		}
	}

	for i := range validatorInfoList {
		if validatorInfoList[i].ValidatorFunc != nil {
			if err := validate.RegisterValidation(validatorInfoList[i].Tag, validatorInfoList[i].ValidatorFunc); err != nil {
				log.Fatal("register validation failed: %v", err)
			}
		}

		if validatorInfoList[i].TranslateFunc != nil {
			if err := validate.RegisterTranslation(validatorInfoList[i].Tag, translator, registerTranslator(validatorInfoList[i].TagDetail, validatorInfoList[i].Msg), validatorInfoList[i].TranslateFunc); err != nil {
				log.Fatal("register validation translation failed: %v", err)
			}
		}
	}
}

func Init() {
	initValidator()

	go func() {
		for {
			select {
			case <-config.HotUpdateForValidator:
				initValidator()
				log.Info("Validator 热更新完成。")
			}
		}
	}()
}

func HandleValidationErr(err error) error {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		errList := make([]string, 0, len(validationErr))
		for _, v := range validationErr.Translate(translator) {
			errList = append(errList, v)
		}

		return errors.WithCode(ecode.ErrParam, strings.Join(errList, ";"))
	} else {
		return errors.WithCode(ecode.ErrParam, err)
	}
}
