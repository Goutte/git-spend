package gitime

import (
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

type TestData struct {
	Collect CollectTestData `yaml:"collect"`
}
type CollectTestData []CollectTestDatum
type CollectTestDatum struct {
	Name     string              `yaml:"rule"`
	Message  string              `yaml:"message"`
	Expected CollectTestExpected `yaml:"expected"`
}

// CollectTestExpected uses pointers, to handle missing values gracefully
type CollectTestExpected struct {
	Minutes   *uint64 `yaml:"minutes"`
	Hours     *uint64 `yaml:"hours"`
	Days      *uint64 `yaml:"days"`
	Weeks     *uint64 `yaml:"weeks"`
	Months    *uint64 `yaml:"months"`
	String    *string `yaml:"string"`
	StringRaw *string `yaml:"string_raw"`
}

func TestCollectTimeSpent(t *testing.T) {

	file := "gitime_test_data.yaml"
	yamlInput, err := os.ReadFile(file)
	require.NoError(t, err)

	testFile := TestData{}

	err = yaml.Unmarshal(yamlInput, &testFile)
	require.NoError(t, err)

	for _, tt := range testFile.Collect {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Expected.Minutes != nil {
				if got := CollectTimeSpent(tt.Message).ToMinutes(); got != *tt.Expected.Minutes {
					t.Errorf("CollectTimeSpent(%s).ToMinutes() = %v, want %v", tt.Message, got, *tt.Expected.Minutes)
				}
			}
			if tt.Expected.Hours != nil {
				if got := CollectTimeSpent(tt.Message).ToHours(); got != *tt.Expected.Hours {
					t.Errorf("CollectTimeSpent(%s).ToHours() = %v, want %v", tt.Message, got, *tt.Expected.Hours)
				}
			}
			if tt.Expected.Days != nil {
				if got := CollectTimeSpent(tt.Message).ToDays(); got != *tt.Expected.Days {
					t.Errorf("CollectTimeSpent(%s).ToDays() = %v, want %v", tt.Message, got, *tt.Expected.Days)
				}
			}
			if tt.Expected.Weeks != nil {
				if got := CollectTimeSpent(tt.Message).ToWeeks(); got != *tt.Expected.Weeks {
					t.Errorf("CollectTimeSpent(%s).ToWeeks() = %v, want %v", tt.Message, got, *tt.Expected.Weeks)
				}
			}
			if tt.Expected.Months != nil {
				if got := CollectTimeSpent(tt.Message).ToMonths(); got != *tt.Expected.Months {
					t.Errorf("CollectTimeSpent(%s).ToMonths() = %v, want %v", tt.Message, got, *tt.Expected.Months)
				}
			}
			if tt.Expected.String != nil {
				if got := CollectTimeSpent(tt.Message).Normalize().String(); got != *tt.Expected.String {
					t.Errorf("CollectTimeSpent(%s).Normalize().String() = %v, want %v", tt.Message, got, *tt.Expected.String)
				}
			}
			if tt.Expected.StringRaw != nil {
				if got := CollectTimeSpent(tt.Message).String(); got != *tt.Expected.StringRaw {
					t.Errorf("CollectTimeSpent(%s).String() = %v, want %v", tt.Message, got, *tt.Expected.StringRaw)
				}
			}
		})
	}
}
