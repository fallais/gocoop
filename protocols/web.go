package protocols

import (
	"time"
)

// APIControllerResponse is the structure template for API
type APIControllerResponse struct {
	ErrorID          string `json:"error_id"`
	ErrorMessage     string `json:"error_message"`
	ErrorDescription string `json:"error_description"`
}

// HelloResponse contains the result of the Hello request.
type HelloResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

// Configuration contains the result of the Configuration.
type Configuration struct {
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	MinutesAfterSunrise string  `json:"minutes_after_sunrise"`
	MinutesAfterSunset  string  `json:"minutes_after_sunset"`
}

type Sun struct {
	Today     time.Time
	Yesterday time.Time
}
