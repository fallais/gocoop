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
	GetSunrise(float64, float64) time.Time
	GetSunset(float64, float64) time.Time
	CalcSunrise(time.Time, float64, float64) time.Time
	CalcSunset(time.Time, float64, float64) time.Time
}

// DoorService is the interface
type DoorService interface {
	Status() door.Status
	Open() error
	Close() error
	Stop() error
}
