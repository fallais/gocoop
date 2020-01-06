package services

import (
	"gocoop/pkg/coop"
)

//------------------------------------------------------------------------------
// Interfaces
//------------------------------------------------------------------------------

// CoopService is the interface
type CoopService interface {
	GetStatus() coop.Status
	UpdateStatus(string) error
	Open() error
	Close() error
}

// JwtService is the service for JWT.
type JwtService interface {
	Create(string) (map[string]string, error)
	Get(string) (string, error)
}
