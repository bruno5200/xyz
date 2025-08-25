package util_test

import (
	"errors"
	"testing"

	u "github.com/bruno5200/xyz/util"
)

var ErrTest = errors.New("test error")

func TestInfo(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name   string
		format string
		args   []any
	}{
		{
			"Simple message",
			"Hello, %s!",
			[]any{"World"},
		},
		{
			"Multiple args",
			"Number: %d, String: %s",
			[]any{42, "Test"},
		},
		{
			"Empty",
			"",
			[]any{},
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u.Info(test.format, test.args...)
		})
	}
}

func TestWarn(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name   string
		err    error
		format string
		args   []any
	}{
		{
			"Nil error",
			nil,
			"Warning: %s",
			[]any{"This is a warning"},
		},
		{
			"Non-nil error",
			ErrTest,
			"Warning: %s",
			[]any{"This is a warning"},
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u.Warn(test.err, test.format, test.args...)
		})
	}
}

func TestFail(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name   string
		err    error
		format string
		args   []any
	}{
		{
			"Nil error",
			nil,
			"Error: %s",
			[]any{"This is an error"},
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u.Fail(test.err, test.format, test.args...)
		})
	}
}