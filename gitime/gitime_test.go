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
type CollectTestExpected struct {
	Minutes uint64 `yaml:"minutes"`
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
			if got := CollectTimeSpent(tt.Message).ToMinutes(); got != tt.Expected.Minutes {
				t.Errorf("CollectTimeSpent(%s) = %v, want %v", tt.Message, got, tt.Expected.Minutes)
			}
		})
	}
}
