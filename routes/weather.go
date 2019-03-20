package routes

import (
	"net/http"
	"time"

	"gocoop-api/protocols"
	"gocoop-api/services"
	"gocoop-api/utils"
	//"goji.io/pat"
	//"github.com/sirupsen/logrus"
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
	nextSunrise := ctrl.weatherService.GetSunrise(time.Now())

	// GetSunrise of yesterday
	yesterdaySunrise := ctrl.weatherService.GetSunrise(time.Now().AddDate(0, 0, -1))

	sunrise := &protocols.Sunrise{
		Today:     nextSunrise,
		Yesterday: yesterdaySunrise,
	}

	// Execute the template
	utils.JSONResponse(w, http.StatusOK, sunrise)
}

// GetSunset of today
func (ctrl *WeatherController) GetSunset(w http.ResponseWriter, r *http.Request) {
	nextSunset := ctrl.weatherService.GetSunset(time.Now())

	// GetSunset of yesterday
	yesterdaySunset := ctrl.weatherService.GetSunset(time.Now().AddDate(0, 0, -1))

	sunset := &protocols.Sunset{
		Today:     nextSunset,
		Yesterday: yesterdaySunset,
	}

	// Execute the template
	utils.JSONResponse(w, http.StatusOK, sunset)
}
