package models

import (
	"time"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Door describes the model of a Door
type Door struct {
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
