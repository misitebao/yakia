package i18n

import (
	"embed"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*
var LocaleFS embed.FS

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	bundle.LoadMessageFileFS(LocaleFS, "locales/en.toml")
	bundle.LoadMessageFileFS(LocaleFS, "locales/zh-Hans.toml")
}

func T(id string) string {
	localizer := i18n.NewLocalizer(bundle, "zh-Hans")

	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: id,
		// DefaultMessage: &i18n.Message{
		// 	ID: id,
		// },
	})
}
