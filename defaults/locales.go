package defaults

import (
	"github.com/go-playground/locales/ar"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/locales/fa"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/id"
	"github.com/go-playground/locales/it"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/lv"
	"github.com/go-playground/locales/nl"
	"github.com/go-playground/locales/pl"
	"github.com/go-playground/locales/pt"
	"github.com/go-playground/locales/pt_BR"
	"github.com/go-playground/locales/ru"
	"github.com/go-playground/locales/tr"
	"github.com/go-playground/locales/uk"
	"github.com/go-playground/locales/vi"
	"github.com/go-playground/locales/zh"
	zh_Hant_TW "github.com/go-playground/locales/zh_Hant_TW"
	"github.com/wolftotem4/golangloader"
	ar_translations "github.com/wolftotem4/golangloader/validate/translations/ar"
	en_translations "github.com/wolftotem4/golangloader/validate/translations/en"
	es_translations "github.com/wolftotem4/golangloader/validate/translations/es"
	fa_translations "github.com/wolftotem4/golangloader/validate/translations/fa"
	fr_translations "github.com/wolftotem4/golangloader/validate/translations/fr"
	id_translations "github.com/wolftotem4/golangloader/validate/translations/id"
	it_translations "github.com/wolftotem4/golangloader/validate/translations/it"
	ja_translations "github.com/wolftotem4/golangloader/validate/translations/ja"
	lv_translations "github.com/wolftotem4/golangloader/validate/translations/lv"
	nl_translations "github.com/wolftotem4/golangloader/validate/translations/nl"
	pl_translations "github.com/wolftotem4/golangloader/validate/translations/pl"
	pt_translations "github.com/wolftotem4/golangloader/validate/translations/pt"
	pt_BR_translations "github.com/wolftotem4/golangloader/validate/translations/pt_BR"
	ru_translations "github.com/wolftotem4/golangloader/validate/translations/ru"
	tr_translations "github.com/wolftotem4/golangloader/validate/translations/tr"
	uk_translations "github.com/wolftotem4/golangloader/validate/translations/uk"
	vi_translations "github.com/wolftotem4/golangloader/validate/translations/vi"
	zh_translations "github.com/wolftotem4/golangloader/validate/translations/zh"
	zh_Hant_TW_translations "github.com/wolftotem4/golangloader/validate/translations/zh_Hant_TW"
)

var Locales = golangloader.MapLocales{
	"ar": {
		Translator:          ar.New(),
		ValidateTranslation: ar_translations.RegisterDefaultTranslations,
	},
	"en": {
		Translator:          en.New(),
		ValidateTranslation: en_translations.RegisterDefaultTranslations,
	},
	"es": {
		Translator:          es.New(),
		ValidateTranslation: es_translations.RegisterDefaultTranslations,
	},
	"fa": {
		Translator:          fa.New(),
		ValidateTranslation: fa_translations.RegisterDefaultTranslations,
	},
	"fr": {
		Translator:          fr.New(),
		ValidateTranslation: fr_translations.RegisterDefaultTranslations,
	},
	"id": {
		Translator:          id.New(),
		ValidateTranslation: id_translations.RegisterDefaultTranslations,
	},
	"it": {
		Translator:          it.New(),
		ValidateTranslation: it_translations.RegisterDefaultTranslations,
	},
	"ja": {
		Translator:          ja.New(),
		ValidateTranslation: ja_translations.RegisterDefaultTranslations,
	},
	"lv": {
		Translator:          lv.New(),
		ValidateTranslation: lv_translations.RegisterDefaultTranslations,
	},
	"nl": {
		Translator:          nl.New(),
		ValidateTranslation: nl_translations.RegisterDefaultTranslations,
	},
	"pl": {
		Translator:          pl.New(),
		ValidateTranslation: pl_translations.RegisterDefaultTranslations,
	},
	"pt": {
		Translator:          pt.New(),
		ValidateTranslation: pt_translations.RegisterDefaultTranslations,
	},
	"pt_BR": {
		Translator:          pt_BR.New(),
		ValidateTranslation: pt_BR_translations.RegisterDefaultTranslations,
	},
	"ru": {
		Translator:          ru.New(),
		ValidateTranslation: ru_translations.RegisterDefaultTranslations,
	},
	"tr": {
		Translator:          tr.New(),
		ValidateTranslation: tr_translations.RegisterDefaultTranslations,
	},
	"uk": {
		Translator:          uk.New(),
		ValidateTranslation: uk_translations.RegisterDefaultTranslations,
	},
	"vi": {
		Translator:          vi.New(),
		ValidateTranslation: vi_translations.RegisterDefaultTranslations,
	},
	"zh": {
		Translator:          zh.New(),
		ValidateTranslation: zh_translations.RegisterDefaultTranslations,
	},
	"zh_Hant_TW": {
		Translator:          zh_Hant_TW.New(),
		ValidateTranslation: zh_Hant_TW_translations.RegisterDefaultTranslations,
	},
}
