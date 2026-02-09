package main

import (
	"testing"
	"time"
)

var timezoneUtcPlusOne, _ = time.Parse("-07:00", "+01:00")
var timezoneUtcMinusEight, _ = time.Parse("-07:00", "-08:00")

var timeStringsShouldParse = []struct {
	in  string
	out time.Time
}{
	{"2026-02-09T07:36:57+01:00", time.Date(2026, 2, 9, 7, 36, 57, 0, timezoneUtcPlusOne.Location())},
	{"2026-02-09T07:36:57Z", time.Date(2026, 2, 9, 7, 36, 57, 0, time.UTC)},
	{"2026-02-09T07:36:57+00:00", time.Date(2026, 2, 9, 7, 36, 57, 0, time.UTC)},
	{"2026-02-09T07:36:57-08:00", time.Date(2026, 2, 9, 7, 36, 57, 0, timezoneUtcMinusEight.Location())},
}

var timeStringsShouldFail = []string{
	"2026-02-09T07:36:57+0100",
	"20260209T07:36:57",
	"20260209",
}

func Test_parseCiCdTimeString_ShouldSucceedParsing(t *testing.T) {
	err := setupTimezone()

	if err != nil {
		t.Errorf("Failed setting up timezone: %v", err)
	}

	for _, tt := range timeStringsShouldParse {
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

func Test_parseCiCdTimeString_ShouldFailParsing(t *testing.T) {
	err := setupTimezone()

	if err != nil {
		t.Errorf("Failed setting up timezone: %v", err)
	}

	for _, tt := range timeStringsShouldFail {
		t.Run(tt, func(t *testing.T) {
			out, err := parseCiCdTimeString(tt)
			if err == nil {
				t.Errorf("Timestring shouldn't parse but succeeded: '%s', '%v'", tt, out)
			}
		})
	}
}
