package coop

import (
	"fmt"
	"strconv"
	"strings"
)

// Cond is a condition
type Cond struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

// Configuration is the configuration.
type Configuration struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Opening   Cond    `json:"opening"`
	Closing   Cond    `json:"closing"`
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
