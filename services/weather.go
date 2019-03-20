package services

import (
	"time"

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
func (service *weatherService) GetSunrise(latitude, longitude float64) time.Time {
	return astrotime.NextSunrise(time.Now(), latitude, longitude)
}

// GetSunset of today
func (service *weatherService) GetSunset(latitude, longitude float64) time.Time {
	return astrotime.NextSunset(time.Now(), latitude, longitude)
}

// CalcSunrise of today
func (service *weatherService) CalcSunrise(date time.Time, latitude, longitude float64) time.Time {
	return astrotime.CalcSunrise(date, latitude, longitude)
}

// CalcSunset of today
func (service *weatherService) CalcSunset(date time.Time, latitude, longitude float64) time.Time {
	return astrotime.CalcSunset(date, latitude, longitude)
}
