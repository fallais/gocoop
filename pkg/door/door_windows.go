// +build windows

package door

import (
	"time"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Door is a physical door manipulated with a motor.
type Door struct {
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewDoor returns a new Door.
func NewDoor(pin1A, pin1B, pin1Enable int, openingDuration, closingDuration time.Duration) *Door {
	return &Door{}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Open the door
func (d *Door) Open() error {
	logrus.Warningln("Coop cannot work on Windows")
	return nil
}

// Close the door
func (d *Door) Close() error {
	logrus.Warningln("Coop cannot work on Windows")
	return nil
}
