package main

import (
	"testing"
	"time"
)

var timezoneCet, _ = time.LoadLocation("CET")

var timeStrings = []struct {
	in  string
	out time.Time
}{
	{"2026-02-09T07:36:57+01:00", time.Date(2026, 2, 9, 7, 36, 57, 0, timezoneCet)},
	{"2026-02-09T07:36:57Z", time.Date(2026, 2, 9, 7, 36, 57, 0, time.UTC)},
}

func TestStringParsing(t *testing.T) {
	for _, tt := range timeStrings {
		t.Run(tt.in, func(t *testing.T) {
			out, err := parseCiCdTimeString(tt.in)
			if out.String() != tt.out.String() {
				t.Errorf("%s != %s", out.String(), tt.out.String())
			} else if err != nil {
				t.Errorf("%v", err)
			}
		})
	}
}
