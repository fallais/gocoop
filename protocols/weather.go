package protocols

import "time"

//------------------------------------------------------------------------------
// Controller
//------------------------------------------------------------------------------

// OpenWeatherMap ...
type OpenWeatherMap struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeH float64 `json:"3h"`
	} `json:"rain"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

// WeatherCreateRequest contains ...
type WeatherCreateRequest struct {
	URL     string `json:"url,omitempty"`
	Token   string `json:"token,omitempty"`
	Version string `json:"version,omitempty"`
}

// WeatherUpdateRequest contains ...
type WeatherUpdateRequest struct {
	ID      string `json:"_id,omitempty"`
	URL     string `json:"url,omitempty"`
	Token   string `json:"token,omitempty"`
	Version string `json:"version,omitempty"`
}

// Sunrise contains ...
type Sunrise struct {
	Today     time.Time `json:"today,omitempty"`
	Yesterday time.Time `json:"yesterday,omitempty"`
}

// Sunset contains ...
type Sunset struct {
	Today     time.Time `json:"today,omitempty"`
	Yesterday time.Time `json:"yesterday,omitempty"`
}
