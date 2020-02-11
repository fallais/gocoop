package coop

import (
	"github.com/fallais/gocoop/pkg/coop/conditions"
	"github.com/fallais/gocoop/pkg/door"
	"github.com/fallais/gocoop/pkg/notifiers"
)

// options are the options for the coop.
type options struct {
	openingCondition conditions.Condition
	closingCondition conditions.Condition
	door             *door.Door
	notifiers        []notifiers.Notifier
	isAutomatic      bool
	latitude         float64
	longitude        float64
}

// Option is a single option.
type Option func(*options)

// WithAutomatic sets the automatic mode.
func WithAutomatic() Option {
	return func(o *options) {
		o.isAutomatic = true
	}
}

// WithOpeningCondition sets the opening condition.
func WithOpeningCondition(oc conditions.Condition) Option {
	return func(o *options) {
		o.openingCondition = oc
	}
}

// WithClosingCondition sets the closing condition.
func WithClosingCondition(cc conditions.Condition) Option {
	return func(o *options) {
		o.closingCondition = cc
	}
}
