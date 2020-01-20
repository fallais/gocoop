package coop

import (
	"time"
)

func (coop *Coop) shouldBeClosed(date time.Time) bool {
	// Check if the date if before the opening time (it is the morning)
	openingTime := coop.openingCondition.GetOpeningTime()
	if date.Before(openingTime) {
		return true
	}

	// Check if the date is after the closing time (it is the evening)
	closingTime := coop.closingCondition.GetClosingTime()
	if date.After(closingTime) {
		return true
	}

	return false
}

func (coop *Coop) shouldBeOpened(date time.Time) bool {
	// Check if the date is before the opening time
	openingTime := coop.openingCondition.GetOpeningTime()
	if date.Before(openingTime) {
		return false
	}

	// Check if the date is after the closing time
	closingTime := coop.closingCondition.GetClosingTime()
	if date.After(closingTime) {
		return false
	}

	return true
}
