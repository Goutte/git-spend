package guesser

import (
	"fmt"
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

func GuessLocaleFromEnv() (string, error) {
	for _, envKey := range envVariablesHoldingLocale {
		lang := os.Getenv(envKey)
		if lang != "" {
			return lang, nil
		}
	}

	return "", fmt.Errorf("cannot guess locale")
}
