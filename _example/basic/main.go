package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/wolftotem4/golangloader"
	"github.com/wolftotem4/golangloader/defaults"
)

func main() {
	// Assign the fallback locale
	fallback := "en"

	// Create a new translator and load the default translations
	ut, err := defaults.Locales.NewUniversalTranslator(fallback)
	if err != nil {
		panic(err)
	}

	// Load more translations from .json files
	err = golangloader.LoadTranslate("lang", ut)
	if err != nil {
		panic(err)
	}

	// Uncomment this line to change the field name key
	// golangloader.FieldNameKey = "validation.attributes.%s"

	// Register the translator
	validate := validator.New()
	err = defaults.Locales.RegisterValidateTranslation(validate, ut)
	if err != nil {
		panic(err)
	}

	// Validate required fields
	err = validate.Struct(struct {
		Username string `validate:"required"`
		Password string `validate:"required"`
	}{
		Username: "",
		Password: "",
	})
	validateErrors := err.(validator.ValidationErrors)

	/* Output:

	Username: Username is a required field
	Password: Password is a required field

	*/
	{
		trans, _ := ut.GetTranslator("en")
		fmt.Println(validateErrors.Translate(trans))
	}

	/* Output:

	Username: 會員名稱為必填欄位
	Password: 密碼為必填欄位

	*/
	{
		trans, _ := ut.GetTranslator("zh_Hant_TW")
		fmt.Println(validateErrors.Translate(trans))
	}
}
