package golangloader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	ut "github.com/go-playground/universal-translator"
)

// loads the translations from the .json files in the specified directory
func LoadTranslate(dir string, uni *ut.UniversalTranslator) error {
	return LoadTranslateFS(os.DirFS(dir), uni)
}

// loads the translations from the .json files in the specified directory
func LoadTranslateFS(fsys fs.FS, uni *ut.UniversalTranslator) error {
	dirs, err := fs.ReadDir(fsys, ".")
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

		subFS, err := fs.Sub(fsys, locale)
		if err != nil {
			return err
		}

		err = fs.WalkDir(subFS, ".", func(name string, info fs.DirEntry, err error) error {
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

			fh, err := subFS.Open(name)
			if err != nil {
				return err
			}

			content, err := io.ReadAll(fh)
			if err != nil {
				return err
			}

			data, err := getFlattenJson(content)
			if err != nil {
				return err
			}

			for key, value := range data {
				prefix := strings.ReplaceAll(filepath.Dir(name), "/", ".")
				if os.PathSeparator != '/' {
					prefix = strings.ReplaceAll(filepath.Dir(name), string(os.PathSeparator), ".")
				}

				var newKey string
				if prefix == "." {
					newKey = fmt.Sprintf("%s.%s", filepath.Base(info.Name())[:len(info.Name())-len(EXT)], key)
				} else {
					newKey = fmt.Sprintf("%s.%s.%s", prefix, filepath.Base(info.Name())[:len(info.Name())-len(EXT)], key)
				}

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
