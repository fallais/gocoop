package bts7960

import (
	"github.com/fallais/gocoop/pkg/motor"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// bts7960 is a motor driver.
type bts7960 struct {
	motor1A      int
	motor1B      int
	motor1Enable int
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewL293D returns a new L293D.
func NewL293D(pin1A, pin1B, pin1Enable int) motor.Motor {
	return &bts7960{
		motor1A:      pin1A,
		motor1B:      pin1B,
		motor1Enable: pin1Enable,
	}
}
