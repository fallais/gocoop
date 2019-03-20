package condition

import (
	"time"
)

// Condition is an opening or a closing condition.
type Condition interface {
	GetTime() time.Time
	GetNextTime() time.Time
}
