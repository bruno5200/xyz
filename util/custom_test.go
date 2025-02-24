package util_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	u "github.com/bruno5200/xyz/util"
)

func TestCustomTimeUnmarshalJSON(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		s    string
		want string
		test string
		pass bool
	}{
		{
			"2021-01-01T00:00:00",
			"01/01/2021 00:00:00",
			"valid date",
			true,
		},
		{
			"2021-01-01T00:00:00.9999999",
			"01/01/2021 00:00:00",
			"valid date",
			true,
		},
		{
			"2025-02-11T16:28:01.1331429",
			"02/11/2025 16:28:01",
			"valid date",
			true,
		},
		{
			"null",
			"01/01/0001 00:00:00",
			"null date",
			false,
		},
	}

	type Testable struct {
		CustomTime u.CustomTime `json:"custom_time"`
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.test, func(t *testing.T) {
			t.Parallel()
			var tt Testable
			if err := json.Unmarshal([]byte(fmt.Sprintf(`{"custom_time": "%s"}`, test.s)), &tt); (err != nil) == test.pass {
				t.Errorf("unexpected error: %v", err)
			}
			got := tt.CustomTime.Format("01/02/2006 15:04:05")
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestCustomTimeMarshalJSON(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		t    u.CustomTime
		want string
		test string
	}{
		{
			u.CustomTime{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			`{"custom_time":"2021-01-01T00:00:00"}`,
			"valid date",
		},
		{
			u.CustomTime{Time: time.Date(2021, 1, 1, 0, 0, 0, 9999999, time.UTC)},
			`{"custom_time":"2021-01-01T00:00:00.0099999"}`,
			"valid date",
		},
		{
			u.CustomTime{Time: time.Date(2025, 2, 11, 16, 28, 1, 1331429, time.UTC)},
			`{"custom_time":"2025-02-11T16:28:01.0013314"}`,
			"valid date",
		},
	}

	type Testable struct {
		CustomTime u.CustomTime `json:"custom_time"`
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.test, func(t *testing.T) {
			t.Parallel()
			tt := Testable{CustomTime: test.t}
			b, err := json.Marshal(tt)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			got := string(b)
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}
