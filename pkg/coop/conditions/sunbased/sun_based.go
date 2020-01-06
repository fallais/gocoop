package sunbased

import (
	"time"

	"gocoop/pkg/coop/conditions"

	"github.com/cpucycle/astrotime"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// A sun based condition is based on the sunrise or the sunset, with an offset.
type sunBasedCondition struct {
	offset    time.Duration
	latitude  float64
	longitude float64
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewSunBasedCondition returns a new SunBasedCondition.
func NewSunBasedCondition(offset time.Duration, latitude, longitude float64) conditions.Condition {
	return &sunBasedCondition{
		offset:    offset,
		latitude:  latitude,
		longitude: longitude,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetOpeningTime
func (c *sunBasedCondition) GetOpeningTime() time.Time {
	return astrotime.CalcSunrise(time.Now(), c.latitude, c.longitude).Add(c.offset)
}

// GetClosingTime
func (c *sunBasedCondition) GetClosingTime() time.Time {
	return astrotime.CalcSunset(time.Now(), c.latitude, c.longitude).Add(c.offset)
}
