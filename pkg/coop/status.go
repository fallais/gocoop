package coop

// Status is the status of the coop.
type Status string

const (
	// Opened when the coop is opened.
	Opened Status = "Opened"

	// Closed when the coop is closed.
	Closed Status = "Closed"

	// Opening when the coop is opening.
	Opening Status = "Opening"

	// Closing when the coop is closing.
	Closing Status = "Closing"

	// Unknown when it is unknown.
	Unknown Status = "Unknown"
)
