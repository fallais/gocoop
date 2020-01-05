package services

import (
	"fmt"

	"gocoop/pkg/coop"
)

//------------------------------------------------------------------------------
// Constants
//------------------------------------------------------------------------------

const ErrCoopOpened = "The coop is already opened"
const ErrCoopOpening = "The coop is opening"
const ErrCoopClosed = "The coop is already closed"
const ErrCoopClosing = "The coop is closing"
const ErrCoopOpeningOrClosing = "The coop is not being used"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type coopService struct {
	coop *coop.Coop
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewCoopService returns a new CoopService.
func NewCoopService(coop *coop.Coop) CoopService {
	return &coopService{
		coop: coop,
	}
}

//------------------------------------------------------------------------------
// Services
//------------------------------------------------------------------------------

// Status returns the status of the coop.
func (service *coopService) Status() coop.Status {
	return service.coop.Status()
}

// Open the Coop
func (service *coopService) Open() error {
	// Get the status of the coop
	status := service.coop.Status()

	// Check if coop is opened
	if status == coop.Opened {
		return fmt.Errorf(ErrCoopOpened)
	}

	// Check if coop is opening
	if status == coop.Opening {
		return fmt.Errorf(ErrCoopOpening)
	}

	return service.coop.Open()
}

// Close the Coop
func (service *coopService) Close() error {
	// Get the status of the coop
	status := service.coop.Status()

	// Check if coop is closed
	if status == coop.Closed {
		return fmt.Errorf(ErrCoopClosed)
	}

	// Check if coop is closing
	if status == coop.Closing {
		return fmt.Errorf(ErrCoopClosing)
	}

	return service.coop.Close()
}
