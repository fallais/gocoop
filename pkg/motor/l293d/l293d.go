package l293d

import (
	"github.com/fallais/gocoop/pkg/motor"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// l293d is a motor driver.
type l293d struct {
	motor1A      int
	motor1B      int
	motor1Enable int
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewL293D returns a new L293D.
func NewL293D(pin1A, pin1B, pin1Enable int) motor.Motor {
	return &l293d{
		motor1A:      pin1A,
		motor1B:      pin1B,
		motor1Enable: pin1Enable,
	}
}
