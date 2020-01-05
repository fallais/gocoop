package services

import (
	"gocoop/pkg/coop"
)

//------------------------------------------------------------------------------
// Interfaces
//------------------------------------------------------------------------------

// CoopService is the interface
type CoopService interface {
	Status() coop.Status
	Open() error
	Close() error
}
