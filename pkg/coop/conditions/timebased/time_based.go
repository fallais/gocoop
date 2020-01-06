package timebased

import (
	"time"

	"gocoop/pkg/coop/conditions"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type timeBasedCondition struct {
	mode    string
	hours   int
	minutes int
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewTimeBasedCondition returns a new TimeBasedCondition.
func NewTimeBasedCondition(hours, minutes int) conditions.Condition {
	return &timeBasedCondition{
		hours:   hours,
		minutes: minutes,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetOpeningTime returns the time based on the conditions.
func (c *timeBasedCondition) GetOpeningTime() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), c.hours, c.minutes, 0, 0, time.Local)
}

// GetClosingTime returns the time based on the conditions.
func (c *timeBasedCondition) GetClosingTime() time.Time {
	return c.GetOpeningTime()
}
