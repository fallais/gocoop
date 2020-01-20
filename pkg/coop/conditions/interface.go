package conditions

import (
	"time"
)

// Condition is an opening or a closing condition.
type Condition interface {
	OpeningTime() time.Time
	ClosingTime() time.Time
	Mode() string
	Value() string
}
