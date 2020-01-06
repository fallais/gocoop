package coop

// Cond is a condition
type Cond struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

// Configuration is the configuration.
type Configuration struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Opening   Cond    `json:"opening"`
	Closing   Cond    `json:"closing"`
}
