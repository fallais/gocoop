package routes

import (
	"net/http"
	"time"

	"gocoop-api/protocols"
	"gocoop-api/services"
	"gocoop-api/shared"
	"gocoop-api/utils"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// WeatherController is the controller of Weather
type WeatherController struct {
	weatherService services.WeatherService
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewWeatherController returns a new WeatherController
func NewWeatherController(s services.WeatherService) *WeatherController {
	return &WeatherController{
		weatherService: s,
	}
}

//------------------------------------------------------------------------------
// Routes
//------------------------------------------------------------------------------

// GetSunrise of today
func (ctrl *WeatherController) GetSunrise(w http.ResponseWriter, r *http.Request) {
	// GetSunrise of today
	nextSunrise := ctrl.weatherService.GetSunrise(shared.Latitude, shared.Longitude)

	// GetSunrise of yesterday
	yesterdaySunrise := ctrl.weatherService.CalcSunrise(time.Now().AddDate(0, 0, -1), shared.Latitude, shared.Longitude)

	sunrise := &protocols.Sun{
		Today:     nextSunrise,
		Yesterday: yesterdaySunrise,
	}

	// Execute the template
	utils.JSONResponse(w, http.StatusOK, sunrise)
}

// GetSunset of today
func (ctrl *WeatherController) GetSunset(w http.ResponseWriter, r *http.Request) {
	nextSunset := ctrl.weatherService.GetSunset(shared.Latitude, shared.Longitude)

	// GetSunset of yesterday
	yesterdaySunset := ctrl.weatherService.CalcSunset(time.Now().AddDate(0, 0, -1), shared.Latitude, shared.Longitude)

	sunset := &protocols.Sun{
		Today:     nextSunset,
		Yesterday: yesterdaySunset,
	}

	// Execute the template
	utils.JSONResponse(w, http.StatusOK, sunset)
}
