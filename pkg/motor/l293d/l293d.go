package l293d

import (
	"github.com/fallais/gocoop/pkg/motor"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// l293d is a motor driver.
type l293d struct {
	pinInput1  int
	pinInput2  int
	pinEnable1 int
	pinInput3  int
	pinInput4  int
	pinEnable2 int
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewL293D returns a new L293D.
func NewL293D(pinInput1, pinInput2, pinEnable1 int) motor.Motor {
	return &l293d{
		pinInput1:  pinInput1,
		pinInput2:  pinInput2,
		pinEnable1: pinEnable1,
	}
}
