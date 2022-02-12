//go:build darwin
// +build darwin

package l293d

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Forward makes the motor work in the forward way.
func (motor *l293d) Forward(ctx context.Context) error {
	logrus.Warningln("We are on MAC OS, no action")

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	return nil
}

// Backward makes the motor work in the backward way.
func (motor *l293d) Backward(ctx context.Context) error {
	logrus.Warningln("We are on MAC OS, no action")

	// Wait
	until, _ := ctx.Deadline()
	logrus.Infoln("Wait until", until)
	<-ctx.Done()

	return nil
}

// Stop the motor.
func (motor *l293d) Stop() error {
	logrus.Warningln("We are on MAC OS, no action")

	return nil
}
