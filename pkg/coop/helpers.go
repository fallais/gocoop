package coop

import (
	"time"
)

func (coop *Coop) shouldBeClosed(date time.Time) bool {
	// Check if the date if before the opening time (it is the morning)
	openingTime := coop.openingCondition.OpeningTime()
	if date.Before(openingTime) {
		return true
	}

	// Check if the date is after the closing time (it is the evening)
	closingTime := coop.closingCondition.ClosingTime()
	if date.After(closingTime) {
		return true
	}

	return false
}

func (coop *Coop) shouldBeOpened(date time.Time) bool {
	// Check if the date is before the opening time
	openingTime := coop.openingCondition.OpeningTime()
	if date.Before(openingTime) {
		return false
	}

	// Check if the date is after the closing time
	closingTime := coop.closingCondition.ClosingTime()
	if date.After(closingTime) {
		return false
	}

	return true
}
