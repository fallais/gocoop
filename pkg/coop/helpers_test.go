package coop

import (
	"testing"
	"time"

	"gocoop/pkg/coop/conditions/timebased"
)

const latitude = 43.525776
const longitude = 1.327727

func TestShouldBeClosed(t *testing.T) {
	openingCondition := timebased.NewTimeBasedCondition(8, 30)
	closingCondition := timebased.NewTimeBasedCondition(18, 30)

	c := &Coop{
		openingCondition: openingCondition,
		closingCondition: closingCondition,
		latitude:         latitude,
		longitude:        longitude,
	}

	if c.shouldBeClosed(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 9, 0, 0, 0, time.Local)) {
		t.Errorf("Should not be closed")
		t.Fail()
	}

	if !c.shouldBeClosed(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.Local)) {
		t.Errorf("Should be closed")
		t.Fail()
	}

	if !c.shouldBeClosed(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 19, 0, 0, 0, time.Local)) {
		t.Errorf("Should be closed")
		t.Fail()
	}
}

func TestShouldBeOpened(t *testing.T) {
	openingCondition := timebased.NewTimeBasedCondition(8, 30)
	closingCondition := timebased.NewTimeBasedCondition(18, 30)

	c := &Coop{
		openingCondition: openingCondition,
		closingCondition: closingCondition,
		latitude:         latitude,
		longitude:        longitude,
	}

	if !c.shouldBeOpened(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 14, 0, 0, 0, time.Local)) {
		t.Errorf("Should be opened")
		t.Fail()
	}

	if c.shouldBeOpened(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.Local)) {
		t.Errorf("Should not be opened")
		t.Fail()
	}

	if c.shouldBeOpened(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 19, 0, 0, 0, time.Local)) {
		t.Errorf("Should not be opened")
		t.Fail()
	}
}

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

	t3 := "DDhEE"
	h, m, err = parseTime(t3)
	if err == nil {
		t.Errorf("Should raise an error")
		t.Fail()
	}
}
