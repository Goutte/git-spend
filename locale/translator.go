package locale

import (
	"embed"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/goutte/git-spend/locale/guesser"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// defaultLanguage should be language.Esperanto 💡 ("eo")
var defaultLanguage = language.English
var domain = "strings"
var extension = "toml"

// localeFS points to an embedded filesystem of TOML translation files (eases binary distribution)
//
//go:embed *.toml
var localeFS embed.FS

// Localizer can be used to fetch localized messages
var Localizer *i18n.Localizer

// T fetches the translation of the specified key
func T(key string) string {
	localized, _ := Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})

	return localized
}

// Tf fetches the translation of the specified key and formats it like Sprintf
func Tf(key string, args ...any) string {
	localized, _ := Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})

	return fmt.Sprintf(localized, args...)
}

func loadTranslationFiles(bundle *i18n.Bundle, languages []string) {
	for _, lang := range languages {
		_, _ = bundle.LoadMessageFileFS(
			localeFS,
			fmt.Sprintf("%s.%s.%s", domain, lang, extension),
		)
	}
}

func init() {
	bundle := i18n.NewBundle(defaultLanguage)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	detectedLanguages := guesser.DetectLanguages(defaultLanguage)
	loadTranslationFiles(bundle, detectedLanguages)
	Localizer = i18n.NewLocalizer(bundle, detectedLanguages...)
}
