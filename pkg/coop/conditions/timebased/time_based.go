package timebased

import (
	"fmt"
	"time"

	"github.com/fallais/gocoop/pkg/coop/conditions"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type timeBasedCondition struct {
	hours   int
	minutes int
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewTimeBasedCondition returns a new Condition.
func NewTimeBasedCondition(t string) (conditions.Condition, error) {
	h, m, err := parseTime(t)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the time for the opening condition : %s", err)
	}

	return &timeBasedCondition{
		hours:   h,
		minutes: m,
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// OpeningTime returns the time based on the conditions.
func (c *timeBasedCondition) OpeningTime() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), c.hours, c.minutes, 0, 0, time.Local)
}

// ClosingTime returns the time based on the conditions.
func (c *timeBasedCondition) ClosingTime() time.Time {
	return c.OpeningTime()
}

// NextOpeningTime returns the next time based on the conditions.
func (c *timeBasedCondition) NextOpeningTime() time.Time {
	todayOpeningTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), c.hours, c.minutes, 0, 0, time.Local)
	if time.Now().After(todayOpeningTime) {
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), c.hours, c.minutes, 0, 0, time.Local).AddDate(0, 0, 1)
	}

	return todayOpeningTime
}

// NextClosingTime returns the next time based on the conditions.
func (c *timeBasedCondition) NextClosingTime() time.Time {
	return c.OpeningTime()
}

// Mode returns the mode of the condition.
func (c *timeBasedCondition) Mode() string {
	return "time_based"
}

// Value returns the value of the condition.
func (c *timeBasedCondition) Value() string {
	return fmt.Sprintf("%02dh%02d", c.hours, c.minutes)
}
