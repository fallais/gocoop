package coop

import (
	"github.com/fallais/gocoop/pkg/coop/conditions"
	"github.com/fallais/gocoop/pkg/door"
)

// Options are the options for the coop.
type Options struct {
	OpeningCondition conditions.Condition
	ClosingCondition conditions.Condition
	Door             door.Door
	IsAutomatic      bool
	Latitude         float64
	Longitude        float64
}
