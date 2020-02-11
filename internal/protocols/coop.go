package protocols

import (
	"github.com/fallais/gocoop/pkg/coop/conditions"
)

type Condition struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type CoopUpdateRequestController struct {
	OpeningCondition Condition `json:"opening_condition"`
	ClosingCondition Condition `json:"closing_condition"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	Status           string    `json:"status"`
	IsAutomatic      bool      `json:"is_automatic"`
}

type CoopUpdateRequestService struct {
	OpeningCondition conditions.Condition
	ClosingCondition conditions.Condition
	Latitude         float64
	Longitude        float64
	Status           string
	IsAutomatic      bool
}
