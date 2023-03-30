package gitime

import (
	"fmt"
	"github.com/goutte/git-spend/locale"
	"math"
)

type TimeSpent struct {
	Months  float64
	Weeks   float64
	Days    float64
	Hours   float64
	Minutes float64
}

func (ts *TimeSpent) String() string {
	s := ""

	if ts.Months > 0.0 {
		s += ts.monthsToString()
	}
	if ts.Weeks > 0.0 {
		if s != "" {
			s += " "
		}
		s += ts.weeksToString()
	}
	if ts.Days > 0.0 {
		if s != "" {
			s += " "
		}
		s += ts.daysToString()
	}
	if ts.Hours > 0.0 {
		if s != "" {
			s += " "
		}
		s += ts.hoursToString()
	}
	if ts.Minutes >= 0.1 {
		if s != "" {
			s += " "
		}
		s += ts.minutesToString()
	}

	return s
}

func (ts *TimeSpent) ToMinutes() uint64 {
	minutes := ts.Minutes
	minutes += ts.Hours * MinutesInOneHour
	minutes += ts.Days * MinutesInOneDay
	minutes += ts.Weeks * MinutesInOneWeek
	minutes += ts.Months * MinutesInOneMonth

	return uint64(math.Round(minutes))
}

func (ts *TimeSpent) ToHours() uint64 {
	minutes := ts.ToMinutes()
	hours := math.Round(float64(minutes) / MinutesInOneHour)
	return uint64(hours)
}

func (ts *TimeSpent) ToDays() uint64 {
	minutes := ts.ToMinutes()
	hours := math.Round(float64(minutes) / MinutesInOneDay)
	return uint64(hours)
}

func (ts *TimeSpent) ToWeeks() uint64 {
	minutes := ts.ToMinutes()
	hours := math.Round(float64(minutes) / MinutesInOneWeek)
	return uint64(hours)
}

func (ts *TimeSpent) ToMonths() uint64 {
	minutes := ts.ToMinutes()
	hours := math.Round(float64(minutes) / MinutesInOneMonth)
	return uint64(hours)
}

func (ts *TimeSpent) Add(other *TimeSpent) *TimeSpent {
	ts.Minutes += other.Minutes
	ts.Hours += other.Hours
	ts.Days += other.Days
	ts.Weeks += other.Weeks
	ts.Months += other.Months

	return ts
}

func (ts *TimeSpent) Normalize() *TimeSpent {
	return ts.normalizeFractions().normalizeModuli()
}

func (ts *TimeSpent) normalizeFractions() *TimeSpent {
	var frac float64

	_, frac = math.Modf(ts.Months)
	if frac > 0.0 {
		ts.Months -= frac
		ts.Weeks += frac * WeeksInOneMonth
	}

	_, frac = math.Modf(ts.Weeks)
	if frac > 0.0 {
		ts.Weeks -= frac
		ts.Days += frac * DaysInOneWeek
	}

	_, frac = math.Modf(ts.Days)
	if frac > 0.0 {
		ts.Days -= frac
		ts.Hours += frac * HoursInOneDay
	}

	_, frac = math.Modf(ts.Hours)
	if frac > 0.0 {
		ts.Hours -= frac
		ts.Minutes += frac * MinutesInOneHour
	}

	return ts
}

func (ts *TimeSpent) normalizeModuli() *TimeSpent {

	if ts.Minutes >= MinutesInOneHour {
		remain := math.Mod(ts.Minutes, MinutesInOneHour)
		more := (ts.Minutes - remain) / MinutesInOneHour
		ts.Hours += more
		ts.Minutes = remain
	}

	if ts.Hours >= HoursInOneDay {
		remain := math.Mod(ts.Hours, HoursInOneDay)
		more := (ts.Hours - remain) / HoursInOneDay
		ts.Days += more
		ts.Hours = remain
	}

	if ts.Days >= DaysInOneWeek {
		remain := math.Mod(ts.Days, DaysInOneWeek)
		more := (ts.Days - remain) / DaysInOneWeek
		ts.Weeks += more
		ts.Days = remain
	}

	if ts.Weeks >= WeeksInOneMonth {
		remain := math.Mod(ts.Weeks, WeeksInOneMonth)
		more := (ts.Weeks - remain) / WeeksInOneMonth
		ts.Months += more
		ts.Weeks = remain
	}

	return ts
}

func (ts *TimeSpent) minutesToString() string {
	return formatUnitComponent(
		ts.Minutes,
		locale.T("UnitMinuteSingular"),
		locale.T("UnitMinutePlural"),
	)
}

func (ts *TimeSpent) hoursToString() string {
	return formatUnitComponent(
		ts.Hours,
		locale.T("UnitHourSingular"),
		locale.T("UnitHourPlural"),
	)
}

func (ts *TimeSpent) daysToString() string {
	return formatUnitComponent(
		ts.Days,
		locale.T("UnitDaySingular"),
		locale.T("UnitDayPlural"),
	)
}

func (ts *TimeSpent) weeksToString() string {
	return formatUnitComponent(
		ts.Weeks,
		locale.T("UnitWeekSingular"),
		locale.T("UnitWeekPlural"),
	)
}

func (ts *TimeSpent) monthsToString() string {
	return formatUnitComponent(
		ts.Months,
		locale.T("UnitMonthSingular"),
		locale.T("UnitMonthPlural"),
	)
}

func formatUnitComponent(value float64, singularUnit string, pluralUnit string) string {
	s := ""
	if value > 0.0 {
		var unit string
		if value >= 2.0 {
			unit = pluralUnit
		} else {
			unit = singularUnit
		}
		intPart, fracPart := math.Modf(value)
		if fracPart == 0.0 {
			s += fmt.Sprintf("%d %s", int64(intPart), unit)
		} else {
			s += fmt.Sprintf("%.1f %s", value, unit)
		}
	}
	return s
}
