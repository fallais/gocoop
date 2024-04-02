//go:build linux
// +build linux

package noop

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Forward turns the motor forward.
func (d *noop) Forward(ctx context.Context) error {
	logrus.Infoln("Turn motor forward")

	// Enable the motor
	logrus.Infoln("Start the motor")

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	// Disable the motor
	logrus.Infoln("Stop the motor")
	logrus.Infoln("Motor has been stopped")

	return nil
}

// Backward turns the motor backward.
func (d *noop) Backward(ctx context.Context) error {
	logrus.Infoln("Turn motor backward")

	// Enable the motor
	logrus.Infoln("Start the motor")

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	// Disable the motor
	logrus.Infoln("Stop the motor")
	logrus.Infoln("Motor has been stopped")

	return nil
}

// Stop the motor.
func (d *noop) Stop() error {
	logrus.Infoln("Stopping the motor")

	// Set to LOW
	logrus.Infoln("Stop the motor")
	logrus.Infoln("Motor has been stopped")

	return nil
}
