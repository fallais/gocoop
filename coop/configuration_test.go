package coop

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	t1 := "30"
	h, m, err := parseTime(t1)
	if err == nil {
		t.Errorf("Should raise an error")
		t.Fail()
	}

	t2 := "08h30"
	h, m, err = parseTime(t2)
	if err != nil {
		t.Errorf("Should not raise an error")
		t.Fail()
	}
	if h != 8 {
		t.Errorf("Hours should be 8")
		t.Fail()
	}
	if m != 30 {
		t.Errorf("Minutes should be 30")
		t.Fail()
	}
}
