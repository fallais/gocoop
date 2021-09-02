package services

import (
	"github.com/fallais/gocoop/pkg/coop"
)

// ConditionUpdateRequest ...
type ConditionUpdateRequest struct {
	Mode  string
	Value string
}

// CoopUpdateRequest ...
type CoopUpdateRequest struct {
	OpeningCondition ConditionUpdateRequest
	ClosingCondition ConditionUpdateRequest
	Latitude         float64
	Longitude        float64
	Status           coop.Status
	IsAutomatic      bool
}
