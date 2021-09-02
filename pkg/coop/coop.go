package coop

import (
	"fmt"
	"time"

	"github.com/fallais/gocoop/pkg/coop/conditions"
	"github.com/fallais/gocoop/pkg/coop/conditions/sunbased"
	"github.com/fallais/gocoop/pkg/coop/conditions/timebased"
	"github.com/fallais/gocoop/pkg/door"
	"github.com/fallais/gocoop/pkg/notifiers"
	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Coop represents a chicken coop.
type Coop struct {
	door      door.Door
	ticker    *time.Ticker
	notifiers []notifiers.Notifier

	OpeningCondition conditions.Condition
	ClosingCondition conditions.Condition
	Status           Status
	Latitude         float64
	Longitude        float64
	IsAutomatic      bool
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New returns a new Coop with given latitude and longitude, a door, and options.
func New(latitude, longitude float64, door door.Door, openingConditionMode, openingConditionValue, closingConditionMode, closingConditionValue string, notifiers []notifiers.Notifier, isAutomatic, notifyAtStartup bool) (*Coop, error) {
	// Check latitude and longtitude
	if latitude == 0 && longitude == 0 {
		return nil, ErrIncorrectPosition
	}

	// Create the opening condition
	var openingCondition conditions.Condition
	switch openingConditionMode {
	case "time_based":
		oc, err := timebased.NewTimeBasedCondition(openingConditionValue)
		if err != nil {
			return nil, fmt.Errorf("error while creating the opening condition: %s", err)
		}

		openingCondition = oc
	case "sun_based":
		oc, err := sunbased.NewSunBasedCondition(openingConditionValue, latitude, longitude)
		if err != nil {
			return nil, fmt.Errorf("error while creating the opening condition: %s", err)
		}

		openingCondition = oc
	default:
		return nil, fmt.Errorf("error with the opening mode: %s", viper.GetString("coop.opening.mode"))
	}

	// Create the closing condition
	var closingCondition conditions.Condition
	switch closingConditionMode {
	case "time_based":
		cc, err := timebased.NewTimeBasedCondition(closingConditionValue)
		if err != nil {
			return nil, fmt.Errorf("error while creating the closing condition: %s", err)
		}

		closingCondition = cc
	case "sun_based":
		cc, err := sunbased.NewSunBasedCondition(closingConditionValue, latitude, longitude)
		if err != nil {
			return nil, fmt.Errorf("error while creating the closing condition: %s", err)
		}

		closingCondition = cc
	default:
		return nil, fmt.Errorf("error with the closing mode: %s", viper.GetString("coop.closing.mode"))
	}

	c := &Coop{
		door:             door,
		notifiers:        notifiers,
		ticker:           time.NewTicker(CheckFrequency),
		OpeningCondition: openingCondition,
		ClosingCondition: closingCondition,
		Latitude:         latitude,
		Longitude:        longitude,
		Status:           DefaultStatus,
		IsAutomatic:      isAutomatic,
	}

	// Watch the clock
	go c.watch()

	// Notify that the status is unknown
	if notifyAtStartup {
		go c.notify()
	}

	return c, nil

}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

func (coop *Coop) watch() {
	for range coop.ticker.C {
		go coop.Check()
	}
}

func (coop *Coop) notify() {
	logrus.Infoln("Notifying")
	for _, notifier := range coop.notifiers {
		err := notifier.Notify(NotificationMessage)
		if err != nil {
			logrus.Errorf("error while notifying: %s", err)
		}
	}
}

// NextOpeningTime returns the next opening time of the chicken coop.
func (coop *Coop) NextOpeningTime() time.Time {
	return coop.OpeningCondition.NextOpeningTime()
}

// NextClosingTime returns the next closing time of the chicken coop.
func (coop *Coop) NextClosingTime() time.Time {
	return coop.ClosingCondition.NextClosingTime()
}

// Open opens the chicken coop.
func (coop *Coop) Open() error {
	// Check the automatic mode
	if coop.IsAutomatic {
		return ErrAutomaticModeEnabled
	}

	return coop.open()
}

func (coop *Coop) open() error {
	// Check the incompatible status
	switch coop.Status {
	case Unknown:
		return fmt.Errorf("cannot open the coop because the status unknown")
	case Opened:
		return ErrCoopAlreadyOpened
	case Opening:
		return ErrCoopAlreadyOpening
	case Closing:
		return ErrCoopAlreadyClosing
	}

	// Update the status of the coop
	coop.Status = Opening

	// Open the door
	err := coop.door.Open()
	if err != nil {
		// Update the status of the coop
		coop.Status = Unknown

		return fmt.Errorf("error while opening the door: %s", err)
	}

	// Update the status of the coop
	coop.Status = Opened

	return nil
}

// Close closes the chicken coop.
func (coop *Coop) Close() error {
	// Check the automatic mode
	if coop.IsAutomatic {
		return fmt.Errorf("cannot close the coop because automatic mode is set")
	}

	return coop.close()
}

func (coop *Coop) close() error {
	// Check the incompatible statuses
	switch coop.Status {
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
	coop.Status = Closing

	// Close the door
	err := coop.door.Close()
	if err != nil {
		// Update the status of the coop
		coop.Status = Unknown

		return fmt.Errorf("error while opening the door: %s", err)
	}

	// Update the status of the coop
	coop.Status = Closed

	return nil
}

// Check performs a check of the door of the chicken coop.
func (coop *Coop) Check() {
	// Check the automatic mode
	if !coop.IsAutomatic {
		logrus.WithFields(logrus.Fields{
			"status": coop.Status,
		}).Warningln("Automatic mode is disabled")
		return
	}

	logrus.WithFields(logrus.Fields{
		"status":       coop.Status,
		"opening_time": coop.OpeningCondition.OpeningTime(),
		"closing_time": coop.ClosingCondition.ClosingTime(),
	}).Debugln("Checking the coop")

	// Process the status
	switch coop.Status {
	case Unknown:
		logrus.Warningln("The status is unknown")
	case Opening:
		logrus.Infoln("The coop is opening")
	case Closing:
		logrus.Infoln("The coop is closing")
	case Closed:
		if coop.shouldBeOpened(time.Now()) {
			logrus.WithFields(logrus.Fields{
				"status":       coop.Status,
				"opening_time": coop.OpeningCondition.OpeningTime(),
				"closing_time": coop.ClosingCondition.ClosingTime(),
			}).Warnln("The coop should be opened")

			// Open the coop
			err := coop.open()
			if err != nil {
				logrus.Errorf("error while opening the coop: %s", err)
				return
			}

			logrus.Infoln("The coop has been opened")
		}
	case Opened:
		if coop.shouldBeClosed(time.Now()) {
			logrus.WithFields(logrus.Fields{
				"status":       coop.Status,
				"opening_time": coop.OpeningCondition.OpeningTime(),
				"closing_time": coop.ClosingCondition.ClosingTime(),
			}).Warnln("The coop should be closed")

			// Close the coop
			err := coop.close()
			if err != nil {
				logrus.Errorf("Error when closing the coop: %s", err)
				return
			}

			logrus.Infoln("The coop has been closed")
		}
	default:
		logrus.Errorf("Wrong status for the coop : %s", coop.Status)
	}

	logrus.WithFields(logrus.Fields{
		"status":       coop.Status,
		"opening_time": coop.OpeningCondition.OpeningTime(),
		"closing_time": coop.ClosingCondition.ClosingTime(),
	}).Debugln("Coop has been checked")
}
