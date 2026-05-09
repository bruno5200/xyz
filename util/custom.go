package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05.9999999"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {

	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		// Handle cases where the JSON value is not a string (e.g., null)
		// or if it's an invalid string literal.
		return fmt.Errorf("error unmarshalling CustomTime JSON: %w", err)
	}

	ct.Time, err = time.Parse(ctLayout, strings.TrimSpace(s)) // Trim space in case of leading/trailing whitespace

	if err != nil {
		return fmt.Errorf("error parsing time %q as %q: %v", s, ctLayout, err)
	}

	return
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format(ctLayout))
}
