package gitime

import (
	"fmt"
	"math"
)

const (
	WeeksInOneMonth   = 4.0
	DaysInOneWeek     = 5.0
	HoursInOneDay     = 8.0
	MinutesInOneHour  = 60.0
	MinutesInOneDay   = MinutesInOneHour * HoursInOneDay
	MinutesInOneWeek  = MinutesInOneHour * HoursInOneDay * DaysInOneWeek
	MinutesInOneMonth = MinutesInOneHour * HoursInOneDay * DaysInOneWeek * WeeksInOneMonth
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
		s += fmt.Sprintf("%.1f month", ts.Months)
		if ts.Months >= 2.0 {
			s += "s"
		}
	}
	if ts.Weeks > 0.0 {
		if s != "" {
			s += " "
		}
		s += fmt.Sprintf("%.1f week", ts.Weeks)
		if ts.Weeks >= 2.0 {
			s += "s"
		}
	}
	if ts.Days > 0.0 {
		if s != "" {
			s += " "
		}
		s += fmt.Sprintf("%.1f day", ts.Days)
		if ts.Days >= 2.0 {
			s += "s"
		}
	}
	if ts.Hours > 0.0 {
		if s != "" {
			s += " "
		}
		s += fmt.Sprintf("%.1f hour", ts.Hours)
		if ts.Hours >= 2.0 {
			s += "s"
		}
	}
	if ts.Minutes > 0.0 {
		if s != "" {
			s += " "
		}
		s += fmt.Sprintf("%.1f minute", ts.Minutes)
		if ts.Minutes >= 2.0 {
			s += "s"
		}
	}

	return s
}

func (ts *TimeSpent) ToMinutes() uint64 {
	minutes := ts.Minutes
	minutes += ts.Hours * MinutesInOneHour
	minutes += ts.Days * MinutesInOneDay
	minutes += ts.Weeks * MinutesInOneWeek
	minutes += ts.Months * MinutesInOneMonth

	return uint64(minutes)
}

func (ts *TimeSpent) Add(other *TimeSpent) *TimeSpent {
	ts.Months += other.Months
	ts.Weeks += other.Weeks
	ts.Days += other.Days
	ts.Hours += other.Hours
	ts.Minutes += other.Minutes

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
