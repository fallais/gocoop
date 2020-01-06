package coop

import (
	"fmt"
	"time"

	"gocoop/pkg/coop/conditions"
	"gocoop/pkg/coop/conditions/sunbased"
	"gocoop/pkg/coop/conditions/timebased"
	"gocoop/pkg/coop/door"

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
	door             *door.Door
	status           Status
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
	switch viper.GetString("coop.opening.mode") {
	case "time_based":
		h, m, err := parseTime(viper.GetString("coop.opening.value"))
		if err != nil {
			return nil, fmt.Errorf("Error while parsing the time for the opening condition : %s", err)
		}

		openingCondition = timebased.NewTimeBasedCondition(h, m)

		break
	case "sun_based":
		// Parse the duration
		duration, err := time.ParseDuration(viper.GetString("coop.opening.value"))
		if err != nil {
			return nil, fmt.Errorf("Error when parsing the duration for the opening condition : %s", err)
		}

		openingCondition = sunbased.NewSunBasedCondition(duration, viper.GetFloat64("coop.latitude"), viper.GetFloat64("coop.longitude"))

		break
	default:
		return nil, fmt.Errorf("error with the opening mode: %s", viper.GetString("coop.opening.mode"))
	}

	// Create the closing condition
	var closingCondition conditions.Condition
	switch viper.GetString("coop.closing.mode") {
	case "time_based":
		h, m, err := parseTime(viper.GetString("coop.closing.value"))
		if err != nil {
			return nil, fmt.Errorf("Error while parsing the time for the closing condition : %s", err)
		}

		closingCondition = timebased.NewTimeBasedCondition(h, m)

		break
	case "sun_based":
		// Parse the duration
		duration, err := time.ParseDuration(viper.GetString("coop.closing.value"))
		if err != nil {
			return nil, fmt.Errorf("Error when parsing the duration for the closing condition : %s", err)
		}

		closingCondition = sunbased.NewSunBasedCondition(duration, viper.GetFloat64("coop.latitude"), viper.GetFloat64("coop.longitude"))

		break
	default:
		return nil, fmt.Errorf("error with the closing mode: %s", viper.GetString("closing.mode"))
	}

	return &Coop{
		openingCondition: openingCondition,
		closingCondition: closingCondition,
		latitude:         viper.GetFloat64("latitude"),
		longitude:        viper.GetFloat64("longitude"),
		status:           Unknown,
		door:             door.NewDoor(viper.GetDuration("door.opening_duration"), viper.GetDuration("door.closing_duration")),
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetStatus returns the status of the chicken coop.
func (coop *Coop) GetStatus() Status {
	return coop.status
}

// GetOpeningCondition returns the opening condition of the chicken coop.
func (coop *Coop) GetOpeningCondition() conditions.Condition {
	return coop.openingCondition
}

// GetClosingCondition returns the closing condition of the chicken coop.
func (coop *Coop) GetClosingCondition() conditions.Condition {
	return coop.closingCondition
}

// UpdateStatus updates the status of the chicken coop.
func (coop *Coop) UpdateStatus(status string) error {
	switch status {
	case "opened":
		coop.status = Opened
	case "closed":
		coop.status = Closed
	default:
		return fmt.Errorf("bad status")
	}

	return nil
}

// Open opens the chicken coop.
func (coop *Coop) Open() error {
	switch coop.status {
	case Unknown:
		return fmt.Errorf("cannot open the coop because the status unknown")
	case Opened:
		return fmt.Errorf("coop is already opened")
	case Opening:
		return fmt.Errorf("coop is already opening")
	case Closing:
		return fmt.Errorf("coop is already closing")
	}

	// Update the status of the coop
	coop.status = Opening

	// Open the door
	err := coop.door.Open()
	if err != nil {
		// Update the status of the coop
		coop.status = Unknown

		return fmt.Errorf("error while opening the door: %s", err)
	}

	// Update the status of the coop
	coop.status = Opened

	return nil
}

// Close closes the chicken coop.
func (coop *Coop) Close() error {
	switch coop.status {
	case Unknown:
		return fmt.Errorf("cannot open the coop because the status unknown")
	case Closed:
		return fmt.Errorf("coop is already closed")
	case Opening:
		return fmt.Errorf("coop is already opening")
	case Closing:
		return fmt.Errorf("coop is already closing")
	}

	// Update the status of the coop
	coop.status = Closing

	// Close the door
	err := coop.door.Close()
	if err != nil {
		// Update the status of the coop
		coop.status = Unknown

		return fmt.Errorf("error while opening the door: %s", err)
	}

	// Update the status of the coop
	coop.status = Closed

	return nil
}

// Check performs a check of the door of the chicken coop.
func (coop *Coop) Check() {
	logrus.WithFields(logrus.Fields{
		"status":       coop.status,
		"opening_time": coop.openingCondition.GetOpeningTime(),
		"closing_time": coop.closingCondition.GetClosingTime(),
	}).Infoln("Checking the coop")

	// Process the status
	switch coop.status {
	case Unknown:
		logrus.Warningln("The status is unknown, skipping")
		break
	case Opening:
		logrus.Infoln("The coop is opening, skipping")
		break
	case Closing:
		logrus.Infoln("The coop is closing, skipping")
		break
	case Closed:
		if coop.shouldBeOpened(time.Now()) {
			logrus.WithFields(logrus.Fields{
				"status":       coop.status,
				"opening_time": coop.openingCondition.GetOpeningTime(),
				"closing_time": coop.closingCondition.GetClosingTime(),
			}).Warnln("The coop should be opened")

			// Open the coop
			err := coop.Open()
			if err != nil {
				logrus.Errorf("error while opening the coop: %s", err)
				return
			}

			logrus.Infoln("The coop has been opened")
		}

		break
	case Opened:
		if coop.shouldBeClosed(time.Now()) {
			logrus.WithFields(logrus.Fields{
				"status":       coop.status,
				"opening_time": coop.openingCondition.GetOpeningTime(),
				"closing_time": coop.closingCondition.GetClosingTime(),
			}).Warnln("The coop should be closed")

			// Close the coop
			err := coop.Close()
			if err != nil {
				logrus.Errorf("Error when closing the coop: %s", err)
				return
			}

			logrus.Infoln("The coop has been closed")
		}

		break
	default:
		logrus.Errorf("Wrong status for the coop : %s", coop.status)
		return
	}

	logrus.WithFields(logrus.Fields{
		"status":       coop.status,
		"opening_time": coop.openingCondition.GetOpeningTime(),
		"closing_time": coop.closingCondition.GetClosingTime(),
	}).Infoln("Coop has been checked")
}
