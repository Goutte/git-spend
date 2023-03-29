package locale

import (
	"embed"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/goutte/git-spend/locale/guesser"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

// defaultLanguage should be language.Esperanto 💡 ("eo")
var defaultLanguage = language.English

// localeFS points to an embedded filesystem of TOML files (eases binary distribution)
//
//go:embed *.toml
var localeFS embed.FS

// Localizer can be used to fetch localized messages
var Localizer *i18n.Localizer

func T(key string) string {
	localized, _ := Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: key,
		//DefaultMessage: &i18n.Message{
		//	ID: "CommandRootDescription",
		//	Other: `Administri tempo-spurado "/spend …" direktivojn en commit-mesaĝoj.`,
		//},
	})

	return localized
}

func loadFirstMessageFileFound(bundle *i18n.Bundle, localeTag language.Tag, domain string, extension string) error {
	base, _ := localeTag.Base()
	files := []string{
		fmt.Sprintf("%s.%s.%s", domain, localeTag, extension),
		fmt.Sprintf("%s.%s.%s", domain, base, extension),
	}

	for _, fileName := range files {
		_, err := bundle.LoadMessageFileFS(localeFS, fileName)
		if err != nil {
			continue
		}

		return nil
	}

	return fmt.Errorf("cannot find any message file")
}

func init() {
	locale, err := guesser.GuessLocaleFromEnv()
	if err != nil {
		locale = defaultLanguage.String()
	}

	localeTag := language.Make(locale)
	bundle := i18n.NewBundle(defaultLanguage)

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	err = loadFirstMessageFileFound(bundle, localeTag, "strings", "toml")
	if err != nil {
		log.Fatalln("language not available:", err)
	}

	base, _ := localeTag.Base()
	Localizer = i18n.NewLocalizer(bundle, localeTag.String(), base.String())
}
