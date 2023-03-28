package reader

import "time"

func parseTimePerhaps(input string) *time.Time {
	if input == "" {
		return nil
	}

	layouts := []string{
		time.RFC3339,
		time.DateTime,
		time.DateOnly,
		time.RFC822,
		time.RFC850,
	}

	for _, layout := range layouts {
		parse, err := time.Parse(layout, input)
		if err == nil {
			return &parse
		}
	}

	return nil
}
