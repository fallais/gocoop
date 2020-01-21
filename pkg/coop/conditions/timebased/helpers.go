package timebased

import (
	"fmt"
	"regexp"
	"strconv"
)

// parseTime returns the hours and minutes with given string.
func parseTime(t string) (int, int, error) {
	r := regexp.MustCompile(`^(\d{2})h(\d{2})$`)

	if !r.MatchString(t) {
		return 0, 0, fmt.Errorf("time format is incorrect: %s", t)
	}

	// Split the string
	hoursAndMinutes := r.FindStringSubmatch(t)

	// Hours
	hours, _ := strconv.Atoi(hoursAndMinutes[1])
	if hours > 23 || hours < 0 {
		return 0, 0, fmt.Errorf("incorrect hours: %d", hours)
	}

	// Minutes
	minutes, _ := strconv.Atoi(hoursAndMinutes[2])
	if minutes > 59 || minutes < 0 {
		return 0, 0, fmt.Errorf("incorrect minutes: %d", hours)
	}

	return hours, minutes, nil
}
