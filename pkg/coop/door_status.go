package coop

// DoorStatus is the status of the door.
type DoorStatus string

const (
	// DoorOpened when the door is opened.
	DoorOpened DoorStatus = "Opened"

	// DoorClosed when the door is closed.
	DoorClosed DoorStatus = "Closed"

	// DoorOpening when the door is opening.
	DoorOpening DoorStatus = "Opening"

	// DoorClosing when the door is closing.
	DoorClosing DoorStatus = "Closing"

	// DoorUnknown when it is unknown.
	DoorUnknown DoorStatus = "Unknown"

	// DoorDefault is the default status when starting.
	DoorDefault DoorStatus = DoorUnknown
)
