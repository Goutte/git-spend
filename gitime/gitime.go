package gitime

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var sp = "^/spen[dt]\\s+"
var fl = "[0-9]+[.]?[0-9]*|[0-9]*[.]?[0-9]+"
var mi = "(?P<minutes>" + fl + ")\\s*(mi?|mins?|minutes?)?\\s*"
var ho = "(?P<hours>" + fl + ")\\s*(ho?|hours?)\\s*"
var da = "(?P<days>" + fl + ")\\s*(da?|days?)\\s*"
var we = "(?P<weeks>" + fl + ")\\s*(we?|weeks?)\\s*"
var mo = "(?P<months>" + fl + ")\\s*(mo|months?)\\s*"
var miP = "(" + mi + ")?"
var hoP = "(" + ho + ")?"
var daP = "(" + da + ")?"
var weP = "(" + we + ")?"
var moP = "(" + mo + ")?"

var spentAllRegex = regexp.MustCompile(sp + moP + weP + daP + hoP + miP)

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
	componentFloat, err := strconv.ParseFloat(componentString, 64)
	if err != nil {
		// this should never happen unless we fiddle with and break our regexes
		fmt.Println("cannot parse", component, componentString, r.String())
		return 0
	}

	return componentFloat
}
