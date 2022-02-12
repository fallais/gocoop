//go:build darwin
// +build darwin

package bts7960

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Forward makes the motor work in the forward way.
func (d *bts7960) Forward(ctx context.Context) error {
	logrus.Warningln("We are on MAC OS, no action")

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	return nil
}

// Backward makes the motor work in the backward way.
func (d *bts7960) Backward(ctx context.Context) error {
	logrus.Warningln("We are on MAC OS, no action")

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	return nil
}

// Stop the motor.
func (d *bts7960) Stop() error {
	logrus.Warningln("We are on MAC OS, no action")

	return nil
}
