package services

import (
	"time"

	"gocoop-api/shared"

	"github.com/cpucycle/astrotime"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type weatherService struct{}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewWeatherService returns a new WeatherService
func NewWeatherService() WeatherService {
	return &weatherService{}
}

//------------------------------------------------------------------------------
// Services
//------------------------------------------------------------------------------

// GetSunrise of today
func (service *weatherService) GetSunrise(date time.Time) time.Time {
	return astrotime.NextSunrise(date, shared.Config.Latitude, shared.Config.Longitude)
}

// GetSunset of today
func (service *weatherService) GetSunset(date time.Time) time.Time {
	return astrotime.NextSunset(date, shared.Config.Latitude, shared.Config.Longitude)
}

// CalcSunrise of today
func (service *weatherService) CalcSunrise(date time.Time) time.Time {
	return astrotime.CalcSunrise(date, shared.Config.Latitude, shared.Config.Longitude)
}

// CalcSunset of today
func (service *weatherService) CalcSunset(date time.Time) time.Time {
	return astrotime.CalcSunset(date, shared.Config.Latitude, shared.Config.Longitude)
}
