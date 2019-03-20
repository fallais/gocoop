package coop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"gocoop-api/coop/conditions"
	"gocoop-api/raspberry/door"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Coop represents a chicken coop.
type Coop struct {
	openingCondition conditions.Condition
	closingCondition conditions.Condition
	location         *time.Location
	doors            []*door.Door
	latitude         float64
	longitude        float64
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New returns a new Coop with the given configuration file.
func New(filename string) (*Coop, error) {
	var configuration Configuration

	// Read configuration file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error while reading the configuration file : %s", err)
	}

	// Unmarshal configuration file
	err = json.Unmarshal(file, &configuration)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the configuration : %s", err)
	}

	// Load the location
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		return nil, fmt.Errorf("Failed to load location : %s", err)
	}

	// Create the opening condition
	var openingCondition conditions.Condition
	switch configuration.Opening.Mode {
	case "time_based":
		h, m, err := parseTime(configuration.Opening.Value)
		if err != nil {
			return nil, fmt.Errorf("Error while parsing the time for the opening conditions : %s", err)
		}

		openingCondition = conditions.NewTimeBasedCondition(h, m, loc)

		break
	case "sun_based":
		// Parse the duration
		duration, err := time.ParseDuration(configuration.Opening.Value)
		if err != nil {
			return nil, fmt.Errorf("Error when parsing the duration for the opening conditions : %s", err)
		}

		openingCondition = conditions.NewSunBasedCondition(duration, configuration.Latitude, configuration.Longitude, loc)

		break
	}

	// Create the closing condition
	var closingCondition conditions.Condition
	switch configuration.Closing.Mode {
	case "time_based":
		h, m, err := parseTime(configuration.Closing.Value)
		if err != nil {
			return nil, fmt.Errorf("Error while parsing the time for the opening conditions : %s", err)
		}

		closingCondition = conditions.NewTimeBasedCondition(h, m, loc)

		break
	case "sun_based":
		// Parse the duration
		duration, err := time.ParseDuration(configuration.Closing.Value)
		if err != nil {
			return nil, fmt.Errorf("Error when parsing the duration for the opening conditions : %s", err)
		}

		closingCondition = conditions.NewSunBasedCondition(duration, configuration.Latitude, configuration.Longitude, loc)

		break
	}

	// Create the doors
	var doors []*door.Door
	door := door.New()
	doors = append(doors, door)

	return &Coop{
		openingCondition: openingCondition,
		closingCondition: closingCondition,
		location:         loc,
		latitude:         configuration.Latitude,
		longitude:        configuration.Longitude,
		doors:            doors,
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

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
		case door.DoorUnknown:
			logrus.Infoln("The door is in unknown status, set the status first !")
			break
		case door.DoorOpening:
		case door.DoorClosing:
			logrus.Infoln("The door is already being used, waiting")
			break
		case door.DoorClosed:
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
		case door.DoorOpened:
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
