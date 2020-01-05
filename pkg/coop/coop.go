package coop

import (
	"fmt"
	"time"

	"gocoop/pkg/coop/conditions"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Coop represents a chicken coop.
type Coop struct {
	openingCondition conditions.Condition
	closingCondition conditions.Condition
	location         *time.Location
	doors            []*Door
	latitude         float64
	longitude        float64
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New returns a new Coop.
func New() (*Coop, error) {
	// Create the opening condition
	var openingCondition conditions.Condition

	switch viper.GetString("opening.mode") {
	case "time_based":
		h, m, err := parseTime(viper.GetString("opening.value"))
		if err != nil {
			return nil, fmt.Errorf("Error while parsing the time for the opening conditions : %s", err)
		}

		openingCondition = conditions.NewTimeBasedCondition(h, m)

		break
	case "sun_based":
		// Parse the duration
		duration, err := time.ParseDuration(viper.GetString("opening.value"))
		if err != nil {
			return nil, fmt.Errorf("Error when parsing the duration for the opening conditions : %s", err)
		}

		openingCondition = conditions.NewSunBasedCondition(duration, viper.GetFloat64("latitude"), viper.GetFloat64("longitude"))

		break
	}

	// Create the closing condition
	var closingCondition conditions.Condition
	switch viper.GetString("closing.mode") {
	case "time_based":
		h, m, err := parseTime(viper.GetString("closing.value"))
		if err != nil {
			return nil, fmt.Errorf("Error while parsing the time for the opening conditions : %s", err)
		}

		closingCondition = conditions.NewTimeBasedCondition(h, m)

		break
	case "sun_based":
		// Parse the duration
		duration, err := time.ParseDuration(viper.GetString("closing.value"))
		if err != nil {
			return nil, fmt.Errorf("Error when parsing the duration for the opening conditions : %s", err)
		}

		closingCondition = conditions.NewSunBasedCondition(duration, viper.GetFloat64("latitude"), viper.GetFloat64("longitude"))

		break
	}

	// Create the doors
	var doors []*Door
	door := NewDoor()
	doors = append(doors, door)

	return &Coop{
		openingCondition: openingCondition,
		closingCondition: closingCondition,
		latitude:         viper.GetFloat64("latitude"),
		longitude:        viper.GetFloat64("longitude"),
		doors:            doors,
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Status returns the status of the chicken coop.
func (coop *Coop) Status() Status {
	return Unknown
}

// Open opens the chicken coop.
func (coop *Coop) Open() error {
	return nil
}

// Close closes the chicken coop.
func (coop *Coop) Close() error {
	return nil
}

// Check performs a check for al the doors of the chicken coop.
func (coop *Coop) Check() {
	logrus.Infoln("Checking the doors of the coop")

	// Check the doors
	for _, d := range coop.doors {
		// Get the status of the door
		logrus.Infoln("Getting the status of the door")
		status := d.GetStatus()
		logrus.WithFields(logrus.Fields{
			"status": status,
		}).Infoln("Successfully got the status of the door")

		// Process the status
		switch status {
		case DoorUnknown:
			logrus.Infoln("The door is in unknown status, set the status first !")
			break
		case DoorOpening:
		case DoorClosing:
			logrus.Infoln("The door is already being used, waiting")
			break
		case DoorClosed:
			if coop.shouldBeOpened(time.Now().UTC()) {
				logrus.WithFields(logrus.Fields{
					"status":       status,
					"opening_time": coop.openingCondition.GetTime(),
					"closing_time": coop.closingCondition.GetTime(),
				}).Warnln("The door should be opened")

				// Open the door
				err := d.Open()
				if err != nil {
					logrus.Errorf("Error when opening the door : %s", err)
					return
				}

				logrus.Infoln("The door has been opened")
			}

			break
		case DoorOpened:
			if coop.shouldBeClosed(time.Now().UTC()) {
				// Close the door
				err := d.Close()
				if err != nil {
					logrus.Errorf("Error when closing the door : %s", err)
					return
				}

				logrus.Infoln("The door has been closed")
			}

			break
		default:
			logrus.Errorf("Wrong status for the door : %s", status)
			return
		}
	}

	logrus.Infoln("Doors of the coop has been checked")
}

//------------------------------------------------------------------------------
// Helpers
//------------------------------------------------------------------------------

func (coop *Coop) shouldBeClosed(date time.Time) bool {
	// Get the closing time from the conditions
	closingTime := coop.closingCondition.GetTime()

	// Get the opening time from the conditions
	openingTime := coop.openingCondition.GetTime()

	// Check the time
	if date.Before(openingTime) || date.After(closingTime) {
		return true
	}

	return false
}

func (coop *Coop) shouldBeOpened(date time.Time) bool {
	// Get the closing time from the conditions
	closingTime := coop.closingCondition.GetTime()

	// Get the opening time from the conditions
	openingTime := coop.openingCondition.GetTime()

	// Check the time
	if date.After(openingTime) && date.Before(closingTime) {
		return true
	}

	return false
}
