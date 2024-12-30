package golangloader

import (
	"embed"
	"io/fs"
	"testing"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

//go:embed _test/*
var embedExample embed.FS

func TestLoadTranslateFS(t *testing.T) {
	dir, err := fs.Sub(embedExample, "_test/lang")
	if err != nil {
		t.Fatal(err)
	}

	uni := ut.New(en.New())
	err = LoadTranslateFS(dir, uni)
	if err != nil {
		t.Fatal(err)
	}

	trans := uni.GetFallback()
	t.Run("TestLoadTranslateFS", func(t *testing.T) {
		expected := "Username"
		value, err := trans.T("validation.attributes.Username")
		if err != nil {
			t.Fatal(err)
		}

		if value != expected {
			t.Errorf("Expected %s, got %s", expected, value)
		}
	})

	t.Run("TestLoadTranslateFS", func(t *testing.T) {
		expected := "bar"
		value, err := trans.T("nested1.nested2.foo.hello.world")
		if err != nil {
			t.Fatal(err)
		}

		if value != expected {
			t.Errorf("Expected %s, got %s", expected, value)
		}
	})
}
