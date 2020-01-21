package sunbased

import (
	"fmt"
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

// NewSunBasedCondition returns a new Condition.
func NewSunBasedCondition(o string, latitude, longitude float64) (conditions.Condition, error) {
	// Parse the duration
	offset, err := time.ParseDuration(o)
	if err != nil {
		return nil, fmt.Errorf("Error when parsing the duration for the closing condition : %s", err)
	}

	return &sunBasedCondition{
		offset:    offset,
		latitude:  latitude,
		longitude: longitude,
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// OpeningTime
func (c *sunBasedCondition) OpeningTime() time.Time {
	return astrotime.CalcSunrise(time.Now(), c.latitude, c.longitude).Add(c.offset)
}

// ClosingTime
func (c *sunBasedCondition) ClosingTime() time.Time {
	return astrotime.CalcSunset(time.Now(), c.latitude, c.longitude).Add(c.offset)
}

// Mode returns the mode of the condition.
func (c *sunBasedCondition) Mode() string {
	return "sun_based"
}

// Value returns the value of the condition.
func (c *sunBasedCondition) Value() string {
	return c.offset.String()
}
