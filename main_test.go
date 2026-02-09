package main

import (
	"testing"
	"time"
)

var timezoneUtcPlusOne, _ = time.Parse("-07:00", "+01:00")

var timeStrings = []struct {
	in  string
	out time.Time
}{
	{"2026-02-09T07:36:57+01:00", time.Date(2026, 2, 9, 7, 36, 57, 0, timezoneUtcPlusOne.Location())},
	{"2026-02-09T07:36:57Z", time.Date(2026, 2, 9, 7, 36, 57, 0, time.UTC)},
}

func TestStringParsing(t *testing.T) {
	err := setupTimezone()

	if err != nil {
		t.Errorf("%v", err)
	}

	for _, tt := range timeStrings {
		t.Run(tt.in, func(t *testing.T) {
			out, err := parseCiCdTimeString(tt.in)
			if out != tt.out {
				t.Errorf("%v != %v", out, tt.out)
			} else if err != nil {
				t.Errorf("%v", err)
			}
		})
	}
}
