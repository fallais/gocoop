package coop

import (
	"fmt"
	"strconv"
	"strings"
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

// parseTime returns the hours and minutes with given string.
func parseTime(t string) (int, int, error) {
	// Split the string
	hoursAndMinutes := strings.Split(t, "h")

	// Check the length
	if len(hoursAndMinutes) != 2 {
		return 0, 0, fmt.Errorf("time format is incorrect")
	}

	// Hours
	hours, err := strconv.Atoi(hoursAndMinutes[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error while parsing hours : %s", err)
	}

	// Minutes
	minutes, err := strconv.Atoi(hoursAndMinutes[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error while parsing minutes : %s", err)
	}

	return hours, minutes, nil
}
