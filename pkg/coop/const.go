package coop

import (
	"errors"
	"time"
)

// ErrAutomaticModeEnabled is raised when the automatic mode is enabled.
var ErrAutomaticModeEnabled = errors.New("cannot close the coop because automatic mode is enabled")

// ErrCoopAlreadyOpening is raised when the coop is already opening.
var ErrCoopAlreadyOpening = errors.New("coop is already opening")

// ErrCoopAlreadyClosing is raised when the coop is already closing.
var ErrCoopAlreadyClosing = errors.New("coop is already closing")

// ErrIncorrectPosition is raised when latitude and longitude are null.
var ErrIncorrectPosition = errors.New("incorrect latitude and longitude")

// CheckFrequency is the frequency for checking the coop.
const CheckFrequency = 10 * time.Second

// NotificationMessage is the notification message.
const NotificationMessage = "The status of the coop is unknown."
