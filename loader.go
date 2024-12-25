package golangloader

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	ut "github.com/go-playground/universal-translator"
)

// loads the translations from the .json files in the specified directory
func LoadTranslate(dir string, uni *ut.UniversalTranslator) error {
	dirs, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, subdir := range dirs {
		locale := subdir.Name()
		if strings.HasPrefix(locale, ".") {
			continue
		}

		trans, ok := uni.GetTranslator(locale)
		if !ok && trans.Locale() != locale {
			continue
		}

		err := filepath.Walk(filepath.Join("lang", locale), func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			const EXT = ".json"
			if filepath.Ext(info.Name()) != EXT {
				return nil
			}

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			data, err := getFlattenJson(content)
			if err != nil {
				return err
			}

			for key, value := range data {
				newKey := fmt.Sprintf("%s.%s", filepath.Base(info.Name())[:len(info.Name())-len(EXT)], key)
				trans.Add(newKey, value, true)
			}

			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func getFlattenJson(content []byte) (map[string]string, error) {
	var data map[string]json.RawMessage
	err := json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	var flatten = make(map[string]string)
	for key, raw := range data {
		var value string
		err := json.Unmarshal(raw, &value)
		if err != nil {
			children, err := getFlattenJson(raw)
			if err != nil {
				return nil, err
			}

			for childKey, childValue := range children {
				flatten[fmt.Sprintf("%s.%s", key, childKey)] = childValue
			}
		} else {
			flatten[key] = value
		}
	}

	return flatten, nil
}
