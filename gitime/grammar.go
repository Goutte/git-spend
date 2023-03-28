package gitime

import "regexp"

var commandRegex = "^\\s*/spen[dt]\\s*"
var floatRegex = "[0-9]+[.]?[0-9]*|[0-9]*[.]?[0-9]+"

// no negative lookahead in regexp, so we hack around it (to ignore datetime suffix)
// there's also regexp2, but its API needs some more work at the time of this writing
var minutesRegex = "(?P<minutes>" + floatRegex + ")\\s*(?:minutes?|mins?|mi?)?([^-/0-9]|$)"
var hoursRegex = "(?P<hours>" + floatRegex + ")\\s*(?:hours?|ho?)\\s*"
var daysRegex = "(?P<days>" + floatRegex + ")\\s*(?:days?|da?)\\s*"
var weeksRegex = "(?P<weeks>" + floatRegex + ")\\s*(?:weeks?|we?)\\s*"
var monthsRegex = "(?P<months>" + floatRegex + ")\\s*(?:months?|mo)\\s*"
var miP = "(?:" + minutesRegex + ")?"
var hoP = "(?:" + hoursRegex + ")?"
var daP = "(?:" + daysRegex + ")?"
var weP = "(?:" + weeksRegex + ")?"
var moP = "(?:" + monthsRegex + ")?"

var spentAllRegex = regexp.MustCompile(commandRegex + moP + weP + daP + hoP + miP)
