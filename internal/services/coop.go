package services

import (
	"fmt"

	"github.com/fallais/gocoop/internal/protocols"
	"github.com/fallais/gocoop/pkg/coop"
	"github.com/fallais/gocoop/pkg/coop/conditions/sunbased"
	"github.com/fallais/gocoop/pkg/coop/conditions/timebased"

	"github.com/spf13/viper"
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

// Get returns the the coop.
func (service *coopService) Get() *coop.Coop {
	return service.coop
}

// Update updates the coop.
func (service *coopService) Update(input protocols.CoopUpdateRequestController) error {
	var inputService protocols.CoopUpdateRequestService

	// Create the opening condition
	switch input.OpeningCondition.Mode {
	case "time_based":
		openingCondition, err := timebased.NewTimeBasedCondition(input.OpeningCondition.Value)
		if err != nil {
			return fmt.Errorf("Error while creating the opening condition: %s", err)
		}

		inputService.OpeningCondition = openingCondition
	case "sun_based":
		openingCondition, err := sunbased.NewSunBasedCondition(input.OpeningCondition.Value, viper.GetFloat64("coop.latitude"), viper.GetFloat64("coop.longitude"))
		if err != nil {
			return fmt.Errorf("Error while creating the opening condition")
		}

		inputService.OpeningCondition = openingCondition
	default:
		return fmt.Errorf("opening mode is incorrect: %s", input.OpeningCondition.Mode)
	}

	// Create the closing condition
	switch input.ClosingCondition.Mode {
	case "time_based":
		closingCondition, err := timebased.NewTimeBasedCondition(input.ClosingCondition.Value)
		if err != nil {
			return fmt.Errorf("Error when creating the closing condition")
		}

		inputService.ClosingCondition = closingCondition
	case "sun_based":
		closingCondition, err := sunbased.NewSunBasedCondition(input.ClosingCondition.Value, viper.GetFloat64("coop.latitude"), viper.GetFloat64("coop.longitude"))
		if err != nil {
			return fmt.Errorf("Error when creating the closing condition")
		}

		inputService.ClosingCondition = closingCondition
	default:
		return fmt.Errorf("closing mode is incorrect: %s", input.ClosingCondition.Mode)
	}

	inputService.Status = input.Status
	inputService.IsAutomatic = input.IsAutomatic

	return service.coop.Update(inputService)
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
