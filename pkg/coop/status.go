package coop

// Status is the status of the coop.
type Status string

const (
	// Opened when the coop is opened.
	Opened Status = "opened"

	// Closed when the coop is closed.
	Closed Status = "closed"

	// Opening when the coop is opening.
	Opening Status = "opening"

	// Closing when the coop is closing.
	Closing Status = "closing"

	// Unknown when it is unknown.
	Unknown Status = "unknown"
)
