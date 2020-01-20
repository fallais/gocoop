package coop

import (
	"gocoop/pkg/coop/conditions"
	"gocoop/pkg/coop/door"
)

// Options are the options for the coop.
type Options struct {
	OpeningCondition conditions.Condition
	ClosingCondition conditions.Condition
	Door             *door.Door
	Latitude         float64
	Longitude        float64
}
