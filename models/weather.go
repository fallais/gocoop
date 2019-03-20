package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/dchest/uniuri"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Weather describes the model of a Weather
type Weather struct {
	ID         string    `bson:"_id" json:"_id" valid:"alphanum,printableascii"`
	URL        string    `json:"url" valid:"url"`
	Token      string    `json:"token"`
	Version    string    `json:"version"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewWeather returns a valid Weather instance
func NewWeather(url, token, version string) *Weather {
	return &Weather{
		ID:         uniuri.New(),
		URL:        url,
		Token:      token,
		Version:    version,
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Validate entity
func (m *Weather) Validate() error {
	_, err := govalidator.ValidateStruct(m)

	return err
}

// Touch update the modified date
func (m *Weather) Touch() {
	m.ModifiedAt = time.Now().UTC()
}
