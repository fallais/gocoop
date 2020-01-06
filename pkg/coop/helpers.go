package coop

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (coop *Coop) shouldBeClosed(date time.Time) bool {
	// Get the closing time from the conditions
	closingTime := coop.closingCondition.GetTime()

	// Get the opening time from the conditions
	openingTime := coop.openingCondition.GetTime()

	// Check the time
	if date.Before(openingTime) || date.After(closingTime) {
		return true
	}

	return false
}

func (coop *Coop) shouldBeOpened(date time.Time) bool {
	// Get the closing time from the conditions
	closingTime := coop.closingCondition.GetTime()

	// Get the opening time from the conditions
	openingTime := coop.openingCondition.GetTime()

	// Check the time
	if date.After(openingTime) && date.Before(closingTime) {
		return true
	}

	return false
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
