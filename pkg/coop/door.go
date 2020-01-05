package coop

import (
	"time"

	rpi "github.com/nathan-osman/go-rpigpio"
	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Door is a physical door manipulated with a motor.
type Door struct {
	motor1A      int
	motor1B      int
	motor1Enable int
	waitOpen     time.Duration
	waitClose    time.Duration
	status       DoorStatus
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewDoor returns a new Door.
func NewDoor() *Door {
	return &Door{
		motor1A:      23,
		motor1B:      24,
		motor1Enable: 25,
		waitOpen:     65 * time.Second,
		waitClose:    60 * time.Second,
		status:       DoorDefault,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetStatus returns the status of the door.
func (d *Door) GetStatus() DoorStatus {
	return d.status
}

// Open the door
func (d *Door) Open() error {
	logrus.Infoln("Opening the door")

	// Open the pinMotor1A
	logrus.WithFields(logrus.Fields{
		"pin_number": d.motor1A,
		"pin_name":   "1A",
		"mode":       "out",
	}).Infoln("Open the pin")
	pinMotor1A, err := rpi.OpenPin(d.motor1A, rpi.OUT)
	if err != nil {
		return err
	}
	defer pinMotor1A.Close()

	// Open the pinMotor1B
	logrus.WithFields(logrus.Fields{
		"pin_number": d.motor1B,
		"pin_name":   "1B",
		"mode":       "out",
	}).Infoln("Open the pin")
	pinMotor1B, err := rpi.OpenPin(d.motor1B, rpi.OUT)
	if err != nil {
		return err
	}
	defer pinMotor1B.Close()

	// Open the pinMotor1Enable
	logrus.WithFields(logrus.Fields{
		"pin_number": d.motor1Enable,
		"pin_name":   "1Enable",
		"mode":       "out",
	}).Infoln("Open the pin")
	pinMotor1Enable, err := rpi.OpenPin(d.motor1Enable, rpi.OUT)
	if err != nil {
		return err
	}
	defer pinMotor1Enable.Close()

	// Update the status of the door
	d.status = DoorOpening

	// Set the motor rotation
	logrus.Infoln("Set the motor rotation")
	pinMotor1A.Write(rpi.HIGH)
	pinMotor1B.Write(rpi.LOW)

	// Enable the motor
	logrus.Infoln("Enable the motor")
	pinMotor1Enable.Write(rpi.HIGH)

	// Wait
	logrus.Infoln("Wait for", d.waitOpen)
	time.Sleep(d.waitOpen)

	// Disable the motor
	logrus.Infoln("Disable the motor")
	pinMotor1Enable.Write(rpi.LOW)

	// Close all the pins
	pinMotor1Enable.Close()
	pinMotor1A.Close()
	pinMotor1B.Close()

	// Check if the door as been stopped
	if d.status == DoorUnknown {
		return nil
	}

	// Update the status of the door
	d.status = DoorOpened

	logrus.Infoln("Door has been opened")

	return nil
}

// Close the door
func (d *Door) Close() error {
	logrus.Infoln("Closing the door")

	// Open the pinMotor1A
	logrus.Infoln("Open the pin", d.motor1A, "in OUT mode")
	pinMotor1A, err := rpi.OpenPin(d.motor1A, rpi.OUT)
	if err != nil {
		return err
	}
	defer pinMotor1A.Close()

	// Open the pinMotor1B
	logrus.Infoln("Open the pin", d.motor1B, "in OUT mode")
	pinMotor1B, err := rpi.OpenPin(d.motor1B, rpi.OUT)
	if err != nil {
		return err
	}
	defer pinMotor1B.Close()

	// Open the pinMotor1Enable
	logrus.Infoln("Open the pin", d.motor1Enable, "in OUT mode")
	pinMotor1Enable, err := rpi.OpenPin(d.motor1Enable, rpi.OUT)
	if err != nil {
		return err
	}
	defer pinMotor1Enable.Close()

	// Update the status of the door
	d.status = DoorClosing

	// Set the motor rotation
	logrus.Infoln("Set the motor rotation")
	pinMotor1A.Write(rpi.LOW)
	pinMotor1B.Write(rpi.HIGH)

	// Enable the motor
	logrus.Infoln("Enable the motor")
	pinMotor1Enable.Write(rpi.HIGH)

	// Wait
	logrus.Infoln("Wait for", d.waitClose)
	time.Sleep(d.waitClose)

	// Disable the motor
	logrus.Infoln("Disable the motor")
	pinMotor1Enable.Write(rpi.LOW)

	// Close all the pins
	pinMotor1Enable.Close()
	pinMotor1A.Close()
	pinMotor1B.Close()

	// Check if the door as been stopped
	if d.status == DoorUnknown {
		return nil
	}

	// Update the status of the door
	d.status = DoorClosed

	logrus.Infoln("Door has been closed")

	return nil
}

// Stop the door
func (d *Door) Stop() error {
	logrus.Infoln("Stopping the door")

	// Open the pinMotor1A
	logrus.Infoln("Open the pin", d.motor1A, "in OUT mode")
	pinMotor1A, err := rpi.OpenPin(d.motor1A, rpi.OUT)
	if err != nil {
		return err
	}
	defer pinMotor1A.Close()

	// Open the pinMotor1B
	logrus.Infoln("Open the pin", d.motor1B, "in OUT mode")
	pinMotor1B, err := rpi.OpenPin(d.motor1B, rpi.OUT)
	if err != nil {
		return err
	}
	defer pinMotor1B.Close()

	// Open the pinMotor1Enable
	logrus.Infoln("Open the pin", d.motor1Enable, "in OUT mode")
	pinMotor1Enable, err := rpi.OpenPin(d.motor1Enable, rpi.OUT)
	if err != nil {
		return err
	}
	defer pinMotor1Enable.Close()

	// Disable all the pins
	logrus.Infoln("Disable the motor")
	pinMotor1Enable.Write(rpi.LOW)
	pinMotor1A.Write(rpi.LOW)
	pinMotor1B.Write(rpi.LOW)

	// Close all the pins
	pinMotor1Enable.Close()
	pinMotor1A.Close()
	pinMotor1B.Close()

	// Update the status of the door
	d.status = DoorUnknown

	logrus.Infoln("Door has been stopped")

	return nil
}
