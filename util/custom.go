package util

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05.9999999"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {

	s := string(b)

	s = s[1 : len(s)-1] // remove quotes

	ct.Time, err = time.Parse(ctLayout, s)

	if err != nil {
		return fmt.Errorf("error parsing time %q as %q: %v", s, ctLayout, err)
	}

	return
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format(ctLayout))
}
