package conditions

import (
	"time"

	"github.com/cpucycle/astrotime"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// A sun based condition is based on the sunrise or the sunset, with an offset.
type sunBasedCondition struct {
	mode      string
	offset    time.Duration
	location  *time.Location
	latitude  float64
	longitude float64
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewSunBasedCondition returns a new SunBasedCondition.
func NewSunBasedCondition(offset time.Duration, latitude, longitude float64, location *time.Location) Condition {
	return &sunBasedCondition{
		offset:    offset,
		latitude:  latitude,
		longitude: longitude,
		location:  location,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetTime
func (c *sunBasedCondition) GetTime() time.Time {
	return astrotime.CalcSunrise(time.Now().In(c.location), c.latitude, c.longitude).Add(c.offset)
}

// GetNextTime
func (c *sunBasedCondition) GetNextTime() time.Time {
	return c.GetTime().AddDate(0, 0, 1)
}
