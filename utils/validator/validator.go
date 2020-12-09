package validator

import (
	"fmt"
	"ginblog/utils/errormessage"
	"github.com/go-playground/validator/v10"
	untrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/locales/zh_Hans_CN"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

func Validatroy(date interface{})(string,int) {
	validate := validator.New()
	uni := untrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans-CN")
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)

	}
	err = validate.Struct(date)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans),errormessage.ERROR
		}

	}
	return "",errormessage.SUCCESS
}
