package bts7960

import (
	"github.com/fallais/gocoop/pkg/motor"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Motor driver for BTS7960.
type bts7960 struct {
	forwardPWM    int
	reversePWM    int
	forwardEnable int
	reverseEnable int
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewBTS7960 returns a new BTS7960 motor driver.
func NewBTS7960(forwardPWM, reversePWM, forwardEnable, reverseEnable int) motor.Motor {
	return &bts7960{
		forwardPWM:    forwardPWM,
		reversePWM:    reversePWM,
		forwardEnable: forwardEnable,
		reverseEnable: reverseEnable,
	}
}
