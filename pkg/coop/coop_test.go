package coop

import (
	"testing"
	"time"

	"gocoop/pkg/coop/conditions"
)

const latitude = 43.525776
const longitude = 1.327727

func TestShouldBeClosed(t *testing.T) {
	openingCondition := conditions.NewTimeBasedCondition(8, 30)
	closingCondition := conditions.NewTimeBasedCondition(18, 30)

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
	openingCondition := conditions.NewTimeBasedCondition(8, 30)
	closingCondition := conditions.NewTimeBasedCondition(18, 30)

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
