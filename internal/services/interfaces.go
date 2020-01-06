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

// JwtService is the service for JWT.
type JwtService interface {
	Create(string) (map[string]string, error)
	Get(string) (string, error)
}
