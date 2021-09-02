package coop

import (
	"testing"
	"time"

	"github.com/fallais/gocoop/pkg/coop/conditions/timebased"
)

const latitude = 43.525776
const longitude = 1.327727

func TestShouldBeClosed(t *testing.T) {
	openingCondition, _ := timebased.NewTimeBasedCondition("08h30")
	closingCondition, _ := timebased.NewTimeBasedCondition("18h30")

	c := &Coop{
		OpeningCondition: openingCondition,
		ClosingCondition: closingCondition,
		Latitude:         latitude,
		Longitude:        longitude,
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
	openingCondition, _ := timebased.NewTimeBasedCondition("08h30")
	closingCondition, _ := timebased.NewTimeBasedCondition("18h30")

	c := &Coop{
		OpeningCondition: openingCondition,
		ClosingCondition: closingCondition,
		Latitude:         latitude,
		Longitude:        longitude,
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
