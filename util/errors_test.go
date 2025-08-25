package util_test

import (
	"fmt"
	"testing"

	u "github.com/bruno5200/xyz/util"
)

func TestAnnotate(t *testing.T) {

	//table driven tests
	tests := []struct {
		name   string
		err    error
		format string
		args   []any
		want   string
	}{
		{
			name:   "nil error",
			err:    nil,
			format: "someFunction failed with input %d",
			args:   []any{42},
			want:   "",
		},
		{
			name:   "non-nil error",
			err:    fmt.Errorf("original error"),
			format: "someFunction failed with input %d",
			args:   []any{42},
			want:   "someFunction failed with input 42: original error",
		},
		{
			name:   "non-nil error with multiple args",
			err:    fmt.Errorf("another error"),
			format: "operation %s failed with code %d",
			args:   []any{"testOp", 500},
			want:   "operation testOp failed with code 500: another error",
		},
		{
			name:   "non-nil error with no args",
			err:    fmt.Errorf("error without args"),
			format: "simple error",
			args:   []any{},
			want:   "simple error: error without args",
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := u.Annotate(test.err, test.format, test.args...)
			if (got == nil && test.want != "") || (got != nil && got.Error() != test.want) {
				t.Errorf("Annotate() = %v, want %v", got, test.want)
			}
		})
	}
}
