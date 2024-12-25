package golangloader

import (
	"errors"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Locale struct {
	Translator          locales.Translator
	ValidateTranslation func(v *validator.Validate, trans ut.Translator) error
}

type MapLocales map[string]Locale

// returns a new *ut.UniversalTranslator instance
func (l MapLocales) NewUniversalTranslator(fallback string) (*ut.UniversalTranslator, error) {
	fallbackLocale, ok := l[fallback]
	if !ok {
		return nil, errors.New("fallback locale not found")
	}

	var supportedLocaleTranslators = make([]locales.Translator, 0, len(l)-1)
	for langCode, locale := range l {
		if langCode == fallback {
			continue
		}
		supportedLocaleTranslators = append(supportedLocaleTranslators, locale.Translator)
	}

	return ut.New(fallbackLocale.Translator, supportedLocaleTranslators...), nil
}

// registers the translations for all locales
func (l MapLocales) RegisterValidateTranslation(v *validator.Validate, uni *ut.UniversalTranslator) error {
	for _, locale := range l {
		trans, _ := uni.GetTranslator(locale.Translator.Locale())
		if err := locale.ValidateTranslation(v, trans); err != nil {
			return err
		}
	}
	return nil
}
