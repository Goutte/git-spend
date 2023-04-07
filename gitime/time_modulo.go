package gitime

import (
	"github.com/spf13/viper"
)

/*

These are the time modulo constants that Gitlab uses.

These values may be set to another value at runtime using ENV variables, eg:

	GIT_SPEND_HOURS_IN_ONE_DAY=7 git-spend sum

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
	MinutesInOneHour = getConfigFloat([]string{"minutes_per_hour", "minutes_in_one_hour"}, DefaultMinutesInOneHour)
	HoursInOneDay = getConfigFloat([]string{"hours_per_day", "hours_in_one_day"}, DefaultHoursInOneDay)
	DaysInOneWeek = getConfigFloat([]string{"days_per_week", "days_in_one_week"}, DefaultDaysInOneWeek)
	WeeksInOneMonth = getConfigFloat([]string{"weeks_per_month", "weeks_in_one_month"}, DefaultWeeksInOneMonth)
	refreshCompoundConversions()
}

func getConfigFloat(keys []string, defaultValue float64) float64 {
	for _, key := range keys {
		val := viper.GetFloat64(key)
		if val != defaultValue {
			return val
		}
	}

	return defaultValue
}

func refreshCompoundConversions() {
	MinutesInOneDay = MinutesInOneHour * HoursInOneDay
	MinutesInOneWeek = MinutesInOneHour * HoursInOneDay * DaysInOneWeek
	MinutesInOneMonth = MinutesInOneHour * HoursInOneDay * DaysInOneWeek * WeeksInOneMonth
}

func init() {
	refreshCompoundConversions()
	viper.SetDefault("minutes_in_one_hour", DefaultMinutesInOneHour)
	viper.SetDefault("minutes_per_hour", DefaultMinutesInOneHour)
	viper.SetDefault("hours_in_one_day", DefaultHoursInOneDay)
	viper.SetDefault("hours_per_day", DefaultHoursInOneDay)
	viper.SetDefault("days_in_one_week", DefaultDaysInOneWeek)
	viper.SetDefault("days_per_week", DefaultDaysInOneWeek)
	viper.SetDefault("weeks_in_one_month", DefaultWeeksInOneMonth)
	viper.SetDefault("weeks_per_month", DefaultWeeksInOneMonth)
	//viper.RegisterAlias("minutes_per_hour", "minutes_in_one_hour")
	//viper.RegisterAlias("hours_per_day", "hours_in_one_day")
	//viper.RegisterAlias("days_per_week", "days_in_one_week")
	//viper.RegisterAlias("weeks_per_month", "weeks_in_one_month")
}
