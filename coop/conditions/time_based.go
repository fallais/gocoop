package conditions

import (
	"time"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type timeBasedCondition struct {
	mode     string
	hours    int
	minutes  int
	location *time.Location
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewTimeBasedCondition returns a new TimeBasedCondition.
func NewTimeBasedCondition(hours, minutes int, location *time.Location) Condition {
	return &timeBasedCondition{
		hours:    hours,
		minutes:  minutes,
		location: location,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetTime returns the time based on the conditions.
func (c *timeBasedCondition) GetTime() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), c.hours, c.minutes, 0, 0, c.location)
}

// GetNextTime returns the next time based on the conditions (the day after).
func (c *timeBasedCondition) GetNextTime() time.Time {
	return c.GetTime().AddDate(0, 0, 1)
}
