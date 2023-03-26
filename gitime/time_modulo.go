package gitime

import (
	"github.com/spf13/viper"
)

/*

These are the time modulo constants that Gitlab uses.

These values may be set to another value at runtime using ENV variables, eg:

	GITIME_HOURS_IN_ONE_DAY=7 gitime sum

*/

const (
	DefaultWeeksInOneMonth  = 4.0
	DefaultDaysInOneWeek    = 5.0
	DefaultHoursInOneDay    = 8.0
	DefaultMinutesInOneHour = 60.0
)

var (
	WeeksInOneMonth  = DefaultWeeksInOneMonth
	DaysInOneWeek    = DefaultDaysInOneWeek
	HoursInOneDay    = DefaultHoursInOneDay
	MinutesInOneHour = DefaultMinutesInOneHour
)

var (
	MinutesInOneDay   float64
	MinutesInOneWeek  float64
	MinutesInOneMonth float64
)

// UpdateTimeModuloConfiguration must be ran AFTER viper has loaded the config file and env
func UpdateTimeModuloConfiguration() {
	MinutesInOneHour = viper.GetFloat64("minutes_in_one_hour")
	HoursInOneDay = viper.GetFloat64("hours_in_one_day")
	DaysInOneWeek = viper.GetFloat64("days_in_one_week")
	WeeksInOneMonth = viper.GetFloat64("weeks_in_one_month")
	refreshCompoundConversions()
}

func refreshCompoundConversions() {
	MinutesInOneDay = MinutesInOneHour * HoursInOneDay
	MinutesInOneWeek = MinutesInOneHour * HoursInOneDay * DaysInOneWeek
	MinutesInOneMonth = MinutesInOneHour * HoursInOneDay * DaysInOneWeek * WeeksInOneMonth
}

func init() {
	refreshCompoundConversions()
	viper.SetDefault("minutes_in_one_hour", DefaultMinutesInOneHour)
	viper.SetDefault("hours_in_one_day", DefaultHoursInOneDay)
	viper.SetDefault("days_in_one_week", DefaultDaysInOneWeek)
	viper.SetDefault("weeks_in_one_month", DefaultWeeksInOneMonth)
}
