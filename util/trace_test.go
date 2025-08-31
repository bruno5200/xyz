package util_test

import (
	"testing"

	u "github.com/bruno5200/xyz/util"
)

func TestGenTraceId(t *testing.T) {
	id := u.GenTraceId()
	if len(id) != 36 {
		t.Errorf("Expected length 36, got %d", len(id))
	}
}
