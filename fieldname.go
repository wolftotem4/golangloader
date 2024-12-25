package golangloader

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
)

// FieldNameKey is the key used to get the field name translation
var FieldNameKey = "validation.attributes.%s"

// FieldName returns the translated field name
func FieldName(trans ut.Translator, fieldName string) string {
	value, err := trans.T(fmt.Sprintf(FieldNameKey, fieldName))
	if err == nil {
		return value
	}
	return fieldName
}
