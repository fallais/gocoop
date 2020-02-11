package coop

import (
	"github.com/fallais/gocoop/pkg/notifiers"
)

// options are the options for the coop.
type options struct {
	notifiers   []notifiers.Notifier
	isAutomatic bool
}

// Option is a single option.
type Option func(*options)

// WithAutomatic sets the automatic mode.
func WithAutomatic() Option {
	return func(o *options) {
		o.isAutomatic = true
	}
}

// WithNotifiers sets the notifiers.
func WithNotifiers(n []notifiers.Notifier) Option {
	return func(o *options) {
		o.notifiers = n
	}
}
