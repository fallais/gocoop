package services

import (
	"fmt"

	"gocoop-api/raspberry/door"
)

//------------------------------------------------------------------------------
// Constants
//------------------------------------------------------------------------------

const ErrDoorOpened = "The door is already opened"
const ErrDoorOpening = "The door is opening"
const ErrDoorClosed = "The door is already closed"
const ErrDoorClosing = "The door is closing"
const ErrDoorOpeningOrClosing = "The door is not being used"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type doorService struct {
	door *door.Door
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewDoorService returns a new DoorService.
func NewDoorService(door *door.Door) DoorService {
	return &doorService{
		door: door,
	}
}

//------------------------------------------------------------------------------
// Services
//------------------------------------------------------------------------------

// Status of the Door
func (service *doorService) Status() door.Status {
	return service.door.GetStatus()
}

// Open the Door
func (service *doorService) Open() error {
	// Get the status of the door
	status := service.door.GetStatus()

	// Check if door is opened
	if status == door.DoorOpened {
		return fmt.Errorf(ErrDoorOpened)
	}

	// Check if door is opening
	if status == door.DoorOpening {
		return fmt.Errorf(ErrDoorOpening)
	}

	return service.door.Open()
}

// Close the Door
func (service *doorService) Close() error {
	// Get the status of the door
	status := service.door.GetStatus()

	// Check if door is closed
	if status == door.DoorClosed {
		return fmt.Errorf(ErrDoorClosed)
	}

	// Check if door is closing
	if status == door.DoorClosing {
		return fmt.Errorf(ErrDoorClosing)
	}

	return service.door.Close()
}

// Stop the Door
func (service *doorService) Stop() error {
	// Get the status of the door
	status := service.door.GetStatus()

	// Check if the door is opening or closing
	if status != door.DoorOpening && status != door.DoorClosing {
		return fmt.Errorf(ErrDoorOpeningOrClosing)
	}

	return service.door.Stop()
}
