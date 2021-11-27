package motor

import "context"

// Motor is an electrical motor or a linear actuator.
type Motor interface {
	Forward(context.Context) error
	Backward(context.Context) error
	Stop() error
}
