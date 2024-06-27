package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		cronString string
		expected   CronFields
	}{
		{
			"*/15 0 1,15 * 1-5 /usr/bin/find",
			CronFields{
				Minute:     []string{"0", "15", "30", "45"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
				DayOfWeek:  []string{"1", "2", "3", "4", "5"},
				Command:    "/usr/bin/find",
			},
		},
		{
			"0 0 1 1 * /bin/echo",
			CronFields{
				Minute:     []string{"0"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1"},
				Month:      []string{"1"},
				DayOfWeek:  []string{"0", "1", "2", "3", "4", "5", "6"},
				Command:    "/bin/echo",
			},
		},
	}

	for _, test := range tests {
		result, err := parser.Parse(test.cronString)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Parse(%q) = %+v, want %+v", test.cronString, result, test.expected)
		}
	}
}

func TestParseInvalid(t *testing.T) {
	parser := NewParser()

	invalidCronStrings := []string{
		"",
		"*/15 0 1,15 *",
		"invalid cron string",
	}

	for _, cronString := range invalidCronStrings {
		_, err := parser.Parse(cronString)
		if err == nil {
			fmt.Printf("expected error for invalid cron string %q, but got none", cronString)
		}
	}
}
