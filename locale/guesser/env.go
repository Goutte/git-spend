package guesser

import (
	"golang.org/x/text/language"
	"os"
)

// envVariablesHoldingLocale is sorted by decreasing priority (breaks on first found)
// These environment variables are expected to hold a parsable locale (fr_FR, es, en-US, …)
// ADR: https://www.gnu.org/software/gettext/manual/html_node/Locale-Environment-Variables.html
var envVariablesHoldingLocale = []string{
	"GIT_SPEND_LANGUAGE",
	"LANGUAGE",
	"LC_ALL",
	"LANG",
}

func DetectLanguages(defaultLanguage language.Tag) []string {
	var detectedLangs []string
	for _, envKey := range envVariablesHoldingLocale {
		lang := os.Getenv(envKey)
		if lang != "" {
			detectedLang := language.Make(lang)
			appendLang(&detectedLangs, detectedLang)
		}
	}
	appendLang(&detectedLangs, defaultLanguage)

	return detectedLangs
}

func appendLang(langs *[]string, lang language.Tag) {
	langString := lang.String()
	*langs = append(*langs, langString)

	langBase, confidentInBase := lang.Base()
	if confidentInBase != language.No {
		*langs = append(*langs, langBase.String())
		*langs = append(*langs, langBase.ISO3())
	}
}
