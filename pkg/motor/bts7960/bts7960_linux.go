//go:build linux
// +build linux

package bts7960

import (
	"context"
	"fmt"

	"github.com/fallais/gocoop/pkg/motor"

	rpi "github.com/nathan-osman/go-rpigpio"
	"github.com/sirupsen/logrus"
)

// Forward makes the motor work in the forward way.
func (d *bts7960) Forward(ctx context.Context) error {
	logrus.Infoln("Running forward")

	// Open the pinMotor1A
	logrus.WithFields(logrus.Fields{
		"pin_number": d.motor1A,
		"pin_name":   "1A",
		"mode":       "out",
	}).Infoln("Open the pin")
	pinMotor1A, err := rpi.OpenPin(d.motor1A, rpi.OUT)
	if err != nil {
		return fmt.Errorf("error while opening the pin: %s", err)
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
		return fmt.Errorf("error while opening the pin: %s", err)
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
		return fmt.Errorf("error while opening the pin: %s", err)
	}
	defer pinMotor1Enable.Close()

	// Set the motor rotation
	logrus.Infoln("Set the motor rotation")
	pinMotor1A.Write(rpi.HIGH)
	pinMotor1B.Write(rpi.LOW)

	// Enable the motor
	logrus.Infoln("Start the motor")
	pinMotor1Enable.Write(rpi.HIGH)

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	// Disable the motor
	logrus.Infoln("Stop the motor")
	pinMotor1Enable.Write(rpi.LOW)

	// Close all the pins
	pinMotor1Enable.Close()
	pinMotor1A.Close()
	pinMotor1B.Close()

	logrus.Infoln("Door has been stopped")

	return nil
}

// Backward makes the motor work in the backward way.
func (d *bts7960) Backward(ctx context.Context) error {
	logrus.Infoln("Running backward")

	// Open the pinMotor1A
	logrus.Infoln("Open the pin", d.motor1A, "in OUT mode")
	pinMotor1A, err := rpi.OpenPin(d.motor1A, rpi.OUT)
	if err != nil {
		return fmt.Errorf("error while opening the pin: %s", err)
	}
	defer pinMotor1A.Close()

	// Open the pinMotor1B
	logrus.Infoln("Open the pin", d.motor1B, "in OUT mode")
	pinMotor1B, err := rpi.OpenPin(d.motor1B, rpi.OUT)
	if err != nil {
		return fmt.Errorf("error while opening the pin: %s", err)
	}
	defer pinMotor1B.Close()

	// Open the pinMotor1Enable
	logrus.Infoln("Open the pin", d.motor1Enable, "in OUT mode")
	pinMotor1Enable, err := rpi.OpenPin(d.motor1Enable, rpi.OUT)
	if err != nil {
		return fmt.Errorf("error while opening the pin: %s", err)
	}
	defer pinMotor1Enable.Close()

	// Set the motor rotation
	logrus.Infoln("Set the motor rotation")
	pinMotor1A.Write(rpi.LOW)
	pinMotor1B.Write(rpi.HIGH)

	// Enable the motor
	logrus.Infoln("Start the motor")
	pinMotor1Enable.Write(rpi.HIGH)

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	// Disable the motor
	logrus.Infoln("Stop the motor")
	pinMotor1Enable.Write(rpi.LOW)

	// Close all the pins
	pinMotor1Enable.Close()
	pinMotor1A.Close()
	pinMotor1B.Close()

	logrus.Infoln("Motor is stopped")

	return nil
}

// Stop the motor.
func (d *bts7960) Stop() error {
	logrus.Infoln("Stopping the motor")

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
	logrus.Infoln("Stop the motor")
	pinMotor1Enable.Write(rpi.LOW)
	pinMotor1A.Write(rpi.LOW)
	pinMotor1B.Write(rpi.LOW)

	// Close all the pins
	pinMotor1Enable.Close()
	pinMotor1A.Close()
	pinMotor1B.Close()

	logrus.Infoln("Motor has been stopped")

	return nil
}
