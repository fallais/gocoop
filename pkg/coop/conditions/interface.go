package conditions

import (
	"time"
)

// Condition is an opening or a closing condition.
type Condition interface {
	GetOpeningTime() time.Time
	//GetNextOpeningTime() time.Time
	GetClosingTime() time.Time
	//GetNextClosingTime() time.Time
}
