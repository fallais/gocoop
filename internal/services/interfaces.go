package services

import (
	"github.com/fallais/gocoop/pkg/coop"
)

//------------------------------------------------------------------------------
// Interfaces
//------------------------------------------------------------------------------

// CoopService is the interface
type CoopService interface {
	Get() *coop.Coop
	Update(CoopUpdateRequest) error
	Open() error
	Close() error
}
