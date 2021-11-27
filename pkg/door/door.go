package door

import (
	"context"
	"time"

	"github.com/fallais/gocoop/pkg/motor"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Door is a physical door manipulated with a motor.
type door struct {
	motor           motor.Motor
	openingDuration time.Duration
	closingDuration time.Duration
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewDoor returns a new Door.
func NewDoor(motor motor.Motor, openingDuration, closingDuration time.Duration) Door {
	return &door{
		motor:           motor,
		openingDuration: openingDuration,
		closingDuration: closingDuration,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Open the door
func (d *door) Open() error {
	logrus.Infoln("Opening the door")

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), d.openingDuration)
	defer cancel()

	// Run the motor in forward
	d.motor.Forward(ctx)

	logrus.Infoln("Door has been opened")

	return nil
}

// Close the door
func (d *door) Close() error {
	logrus.Infoln("Closing the door")

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), d.closingDuration)
	defer cancel()

	// Run the motor in backward
	d.motor.Backward(ctx)

	logrus.Infoln("Door has been closed")

	return nil
}

// Stop the door
func (d *door) Stop() error {
	logrus.Infoln("Stopping the door")

	// Stop the motor
	d.motor.Stop()

	logrus.Infoln("Door has been stopped")

	return nil
}
