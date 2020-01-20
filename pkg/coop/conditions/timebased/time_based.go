package timebased

import (
	"fmt"
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

// GetOpeningTime returns the time based on the conditions.
func (c *timeBasedCondition) GetOpeningTime() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), c.hours, c.minutes, 0, 0, time.Local)
}

// GetClosingTime returns the time based on the conditions.
func (c *timeBasedCondition) GetClosingTime() time.Time {
	return c.GetOpeningTime()
}
