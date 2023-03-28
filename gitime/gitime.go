package gitime

import (
	"regexp"
	"strconv"
	"strings"
)

// Keep these sorted by decreasing priority, since first match breaks.
var expressions = []*regexp.Regexp{
	spentAllRegex,
}

// CollectTimeSpent returns the TimeSpent that was collected from the message
// It reads the Gitlab /spend or /spent commands.
// Available time units: https://docs.gitlab.com/ee/user/project/time_tracking.html#available-time-units
// If no time unit is specified, minutes are assumed.
func CollectTimeSpent(message string) *TimeSpent {
	ts := &TimeSpent{}
	lines := strings.Split(message, "\n")

	for _, line := range lines {
		lineTs := extractTimeSpentFromLine(strings.TrimSpace(line))
		if lineTs == nil {
			continue
		}

		ts.Add(lineTs)
	}

	return ts
}

func extractTimeSpentFromLine(line string) *TimeSpent {
	for _, expression := range expressions {
		ts := extractTimeSpentUsingRegexp(line, expression)
		if ts != nil {
			return ts
		}
	}

	return nil
}

func extractTimeSpentUsingRegexp(line string, r *regexp.Regexp) *TimeSpent {
	matches := r.FindStringSubmatch(line)
	if len(matches) == 0 {
		return nil
	}

	months := extractTimeComponent(matches, r, "months")
	weeks := extractTimeComponent(matches, r, "weeks")
	days := extractTimeComponent(matches, r, "days")
	hours := extractTimeComponent(matches, r, "hours")
	minutes := extractTimeComponent(matches, r, "minutes")

	return &TimeSpent{
		Months:  months,
		Weeks:   weeks,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
	}
}

func extractTimeComponent(matches []string, r *regexp.Regexp, component string) float64 {
	componentIndex := r.SubexpIndex(component)
	componentString := "0"
	if componentIndex != -1 {
		if matches[componentIndex] != "" {
			componentString = matches[componentIndex]
		}
	}
	componentFloat, _ := strconv.ParseFloat(componentString, 64)

	return componentFloat
}
