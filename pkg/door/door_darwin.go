// +build darwin

package door

import (
	"time"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Door is a physical door manipulated with a motor.
type door struct {
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewDoor returns a new Door.
func NewDoor(pin1A, pin1B, pin1Enable int, openingDuration, closingDuration time.Duration) Door {
	return &door{}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Open the door
func (d *door) Open() error {
	logrus.Warningln("Coop cannot work on MacOS")
	return nil
}

// Close the door
func (d *door) Close() error {
	logrus.Warningln("Coop cannot work on MacOS")
	return nil
}

// Close the door
func (d *door) Stop() error {
	logrus.Warningln("Coop cannot work on MacOS")
	return nil
}
