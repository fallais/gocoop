package coop

import (
	"fmt"
	"time"

	"github.com/fallais/gocoop/internal/protocols"
	"github.com/fallais/gocoop/pkg/coop/conditions"
	"github.com/fallais/gocoop/pkg/door"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Coop represents a chicken coop.
type Coop struct {
	door             door.Door
	openingCondition conditions.Condition
	closingCondition conditions.Condition
	opts             options
	status           Status
	latitude         float64
	longitude        float64
	ticker           *time.Ticker
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New returns a new Coop with given latitude and longitude, a door, and options.
func New(latitude, longitude float64, door door.Door, openingCondition, closingCondition conditions.Condition, opts ...Option) (*Coop, error) {
	// Check latitude and longtitude
	if latitude == 0 && longitude == 0 {
		return nil, ErrIncorrectPosition
	}

	c := &Coop{
		door:             door,
		openingCondition: openingCondition,
		closingCondition: closingCondition,
		latitude:         latitude,
		longitude:        longitude,
		ticker:           time.NewTicker(CheckFrequency),
		status:           Unknown,
	}

	// Set options
	for _, opt := range opts {
		opt(&c.opts)
	}

	// Watch the clock
	go c.watch()

	// Notify that the status is unknown
	go c.notify()

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
	for _, notifier := range coop.opts.notifiers {
		err := notifier.Notify(NotificationMessage)
		if err != nil {
			logrus.Errorf("error while notifying: %s", err)
		}
	}
}

// Status returns the status of the chicken coop.
func (coop *Coop) Status() Status {
	return coop.status
}

// IsAutomatic returns the automatic mode.
func (coop *Coop) IsAutomatic() bool {
	return coop.opts.isAutomatic
}

// Latitude returns the latitude of the chicken coop.
func (coop *Coop) Latitude() float64 {
	return coop.latitude
}

// Longitude returns the longitude of the chicken coop.
func (coop *Coop) Longitude() float64 {
	return coop.longitude
}

// OpeningCondition returns the opening condition of the chicken coop.
func (coop *Coop) OpeningCondition() conditions.Condition {
	return coop.openingCondition
}

// ClosingCondition returns the closing condition of the chicken coop.
func (coop *Coop) ClosingCondition() conditions.Condition {
	return coop.closingCondition
}

// OpeningTime returns the opening time of the chicken coop.
func (coop *Coop) OpeningTime() time.Time {
	return coop.openingCondition.OpeningTime()
}

// ClosingTime returns the opening time of the chicken coop.
func (coop *Coop) ClosingTime() time.Time {
	return coop.closingCondition.ClosingTime()
}

// Update updates the chicken coop.
func (coop *Coop) Update(input protocols.CoopUpdateRequestService) error {
	// Update the status
	switch input.Status {
	case "opened":
		coop.status = Opened
	case "closed":
		coop.status = Closed
	default:
		return ErrIncorrectStatus
	}

	// Update the automatic mode
	coop.opts.isAutomatic = input.IsAutomatic

	// Update the opening condition
	coop.openingCondition = input.OpeningCondition

	// Update the closing condition
	coop.closingCondition = input.ClosingCondition

	return nil
}

// Open opens the chicken coop.
func (coop *Coop) Open() error {
	// Check the automatic mode
	if coop.opts.isAutomatic {
		return ErrAutomaticModeEnabled
	}

	return coop.open()
}

func (coop *Coop) open() error {
	// Check the incompatible status
	switch coop.status {
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
	// Check the automatic mode
	if coop.opts.isAutomatic {
		return fmt.Errorf("cannot close the coop because automatic mode is set")
	}

	return coop.close()
}

func (coop *Coop) close() error {
	// Check the incompatible statuses
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
	// Check the automatic mode
	if !coop.opts.isAutomatic {
		logrus.WithFields(logrus.Fields{
			"status": coop.status,
		}).Warningln("Automatic mode is disabled")
		return
	}

	logrus.WithFields(logrus.Fields{
		"status":       coop.status,
		"opening_time": coop.openingCondition.OpeningTime(),
		"closing_time": coop.closingCondition.ClosingTime(),
	}).Debugln("Checking the coop")

	// Process the status
	switch coop.status {
	case Unknown:
		logrus.Warningln("The status is unknown")
	case Opening:
		logrus.Infoln("The coop is opening")
	case Closing:
		logrus.Infoln("The coop is closing")
	case Closed:
		if coop.shouldBeOpened(time.Now()) {
			logrus.WithFields(logrus.Fields{
				"status":       coop.status,
				"opening_time": coop.openingCondition.OpeningTime(),
				"closing_time": coop.closingCondition.ClosingTime(),
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
				"status":       coop.status,
				"opening_time": coop.openingCondition.OpeningTime(),
				"closing_time": coop.closingCondition.ClosingTime(),
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
		logrus.Errorf("Wrong status for the coop : %s", coop.status)
	}

	logrus.WithFields(logrus.Fields{
		"status":       coop.status,
		"opening_time": coop.openingCondition.OpeningTime(),
		"closing_time": coop.closingCondition.ClosingTime(),
	}).Debugln("Coop has been checked")
}
