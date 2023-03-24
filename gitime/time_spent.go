package gitime

import (
	"fmt"
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
	minutes := 0.0
	minutes += ts.Minutes
	minutes += ts.Hours * 60.0
	minutes += ts.Days * 8.0 * 60.0
	minutes += ts.Weeks * 5.0 * 8.0 * 60.0
	minutes += ts.Months * 4.0 * 5.0 * 8.0 * 60.0

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
