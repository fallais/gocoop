//go:build linux
// +build linux

package l293d

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

// Forward turns the motor forward.
func (motor *l293d) Forward(ctx context.Context) error {
	logrus.Infoln("Turn motor forward")

	// Access the pins
	err := rpio.Open()
	if err != nil {
		return fmt.Errorf("error while accessing the pins: %s", err)
	}
	defer rpio.Close()

	// Open the pinInput1 and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": motor.pinInput1,
		"mode":       "out",
	}).Infoln("Open the pin")
	pinInput1 := rpio.Pin(motor.pinInput1)
	pinInput1.Output()

	// Open the pinInput2 and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": motor.pinInput2,
		"mode":       "out",
	}).Infoln("Open the pin")
	pinInput2 := rpio.Pin(motor.pinInput2)
	pinInput2.Output()

	// Open the pinEnable1 and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": motor.pinEnable1,
		"mode":       "out",
	}).Infoln("Open the pin")
	pinEnable1 := rpio.Pin(motor.pinEnable1)
	pinEnable1.Output()

	// Set the motor rotation
	logrus.Infoln("Set the motor rotation")
	pinInput1.High()
	pinInput2.Low()

	// Enable the motor
	logrus.Infoln("Start the motor")
	pinEnable1.High()

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	// Disable the motor
	logrus.Infoln("Stop the motor")
	pinEnable1.Low()

	logrus.Infoln("Door has been stopped")

	return nil
}

// Backward turns the motor backward.
func (motor *l293d) Backward(ctx context.Context) error {
	logrus.Infoln("Turn motor backward")

	// Access the pins
	err := rpio.Open()
	if err != nil {
		return fmt.Errorf("error while accessing the pins: %s", err)
	}
	defer rpio.Close()

	// Open the pinInput1 and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": motor.pinInput1,
		"mode":       "out",
	}).Infoln("Open the pin")
	pinInput1 := rpio.Pin(motor.pinInput1)
	pinInput1.Output()

	// Open the pinInput2 and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": motor.pinInput2,
		"mode":       "out",
	}).Infoln("Open the pin")
	pinInput2 := rpio.Pin(motor.pinInput2)
	pinInput2.Output()

	// Open the pinEnable1 and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": motor.pinEnable1,
		"mode":       "out",
	}).Infoln("Open the pin")
	pinEnable1 := rpio.Pin(motor.pinEnable1)
	pinEnable1.Output()

	// Set the motor rotation
	logrus.Infoln("Set the motor rotation")
	pinInput1.Low()
	pinInput2.Low()

	// Enable the motor
	logrus.Infoln("Start the motor")
	pinEnable1.High()

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	// Disable the motor
	logrus.Infoln("Stop the motor")
	pinMotor1Enable.Low()

	logrus.Infoln("Motor is stopped")

	return nil
}

// Stop the motor.
func (d *l293d) Stop() error {
	logrus.Infoln("Stopping the motor")

	// Access the pins
	err := rpio.Open()
	if err != nil {
		return fmt.Errorf("error while accessing the pins: %s", err)
	}
	defer rpio.Close()

	// Open the pinEnable1 and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": motor.pinEnable1,
		"mode":       "out",
	}).Infoln("Open the pin")
	pinEnable1 := rpio.Pin(motor.pinEnable1)
	pinEnable1.Output()

	// Set pinEnable1 to LOW
	logrus.Infoln("Stop the motor")
	pinEnable1.Low()

	logrus.Infoln("Motor has been stopped")

	return nil
}
