//go:build linux
// +build linux

package bts7960

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

// Forward turns the motor forward.
func (d *bts7960) Forward(ctx context.Context) error {
	logrus.Infoln("Turn motor forward")

	// Access the pins
	err := rpio.Open()
	if err != nil {
		return fmt.Errorf("error while accessing the pins: %s", err)
	}
	defer rpio.Close()

	// Open the forwardPWM and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": d.forwardPWM,
		"mode":       "out",
	}).Infoln("Open the pin")
	forwardPWM := rpio.Pin(d.forwardPWM)
	forwardPWM.Output()

	// Open the forwardEnable and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": d.forwardEnable,
		"mode":       "out",
	}).Infoln("Open the pin")
	forwardEnable := rpio.Pin(d.forwardEnable)
	forwardEnable.Output()

	// Set the motor rotation
	logrus.Infoln("Set the motor rotation")
	forwardPWM.High()

	// Enable the motor
	logrus.Infoln("Start the motor")
	forwardEnable.High()

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	// Disable the motor
	logrus.Infoln("Stop the motor")
	forwardEnable.Low()
	logrus.Infoln("Motor has been stopped")

	return nil
}

// Backward turns the motor backward.
func (d *bts7960) Backward(ctx context.Context) error {
	logrus.Infoln("Turn motor backward")

	// Access the pins
	err := rpio.Open()
	if err != nil {
		return fmt.Errorf("error while accessing the pins: %s", err)
	}
	defer rpio.Close()

	// Open the reversePWM and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": d.reversePWM,
		"mode":       "out",
	}).Infoln("Open the pin")
	reversePWM := rpio.Pin(d.reversePWM)
	reversePWM.Output()

	// Open the reverseEnable and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": d.reverseEnable,
		"mode":       "out",
	}).Infoln("Open the pin")
	reverseEnable := rpio.Pin(d.reverseEnable)
	reverseEnable.Output()

	// Set the motor rotation
	logrus.Infoln("Set the motor rotation")
	reversePWM.High()

	// Enable the motor
	logrus.Infoln("Start the motor")
	reverseEnable.High()

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	// Disable the motor
	logrus.Infoln("Stop the motor")
	reverseEnable.Low()
	logrus.Infoln("Motor has been stopped")

	return nil
}

// Stop the motor.
func (d *bts7960) Stop() error {
	logrus.Infoln("Stopping the motor")

	// Access the pins
	err := rpio.Open()
	if err != nil {
		return fmt.Errorf("error while accessing the pins: %s", err)
	}
	defer rpio.Close()

	// Open the forwardEnable and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": d.forwardEnable,
		"mode":       "out",
	}).Infoln("Open the pin")
	forwardEnable := rpio.Pin(d.motor1B)
	forwardEnable.Output()

	// Open the reverseEnable and set OUT mode
	logrus.WithFields(logrus.Fields{
		"pin_number": d.reverseEnable,
		"mode":       "out",
	}).Infoln("Open the pin")
	reverseEnable := rpio.Pin(d.reverseEnable)
	reverseEnable.Output()

	// Set to LOW
	logrus.Infoln("Stop the motor")
	forwardEnable.Low()
	reverseEnable.Low()
	logrus.Infoln("Motor has been stopped")

	return nil
}
