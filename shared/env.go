package shared

type Action struct {
	Mode     string            `json:"mode"`
	Settings map[string]string `json:"settings"`
}

type Door struct {
	Open  Action `json:"open"`
	Close Action `json:"close"`
}

// Configuration is the configuration.
type Configuration struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Door      Door    `json:"door"`
}

var (
	// Config holds the configuration.
	Config Configuration
)
