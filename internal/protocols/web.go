package protocols

// APIControllerResponse is the response for API.
type APIControllerResponse struct {
	ErrorID          string `json:"error_id,omitempty"`
	ErrorMessage     string `json:"error_message,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// JWToken ...
type JWToken struct {
	Token string `json:"token,omitempty"`
}

// JWTokenRequest ...
type JWTokenRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
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
