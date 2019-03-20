package services

import (
	"time"

	"gocoop-api/raspberry/door"
)

//------------------------------------------------------------------------------
// Interfaces
//------------------------------------------------------------------------------

// WeatherService is the interface
type WeatherService interface {
	GetSunrise(time.Time) time.Time
	GetSunset(time.Time) time.Time
	CalcSunrise(time.Time) time.Time
	CalcSunset(time.Time) time.Time
}

// DoorService is the interface
type DoorService interface {
	Status() door.Status
	Open() error
	Close() error
	Stop() error
}
