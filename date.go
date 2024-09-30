package gohelpers

import (
	"encoding/json"
	"time"

	"github.com/araddon/dateparse"
)

type APIDate struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (ad *APIDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return err
	}
	ad.Time = t
	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (ad APIDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(ad.Time.Format(time.RFC3339))
}
