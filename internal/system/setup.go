package system

import (
	"fmt"

	"github.com/fallais/gocoop/pkg/coop/conditions"
	"github.com/fallais/gocoop/pkg/coop/conditions/sunbased"
	"github.com/fallais/gocoop/pkg/coop/conditions/timebased"

	"github.com/fallais/gocoop/pkg/notifiers"
	"github.com/fallais/gocoop/pkg/notifiers/sms/free"
	"github.com/spf13/viper"
)

// SetupNotifiers ...
func SetupNotifiers() []notifiers.Notifier {
	var providers []notifiers.Notifier

	// Create the providers
	for key := range viper.GetStringMap("notifications") {
		// Extract the sub tree
		sub := viper.Sub(fmt.Sprintf("notifications.%s", key))

		// Create
		switch sub.GetString("type") {
		case "sms":
			switch sub.GetString("vendor") {
			case "free":
				prvd := free.NewProvider(sub.GetStringMap("settings"))
				providers = append(providers, prvd)
			}
		}
	}

	return providers
}

// SetupConditions returns the opening and closing conditions.
func SetupConditions() (conditions.Condition, conditions.Condition, error) {
	var openingCondition conditions.Condition
	var closingCondition conditions.Condition

	// Create the opening condition
	switch viper.GetString("coop.opening.mode") {
	case "time_based":
		oc, err := timebased.NewTimeBasedCondition(viper.GetString("coop.opening.value"))
		if err != nil {
			return nil, nil, fmt.Errorf("error while creating the opening condition: %s", err)
		}

		openingCondition = oc
	case "sun_based":
		oc, err := sunbased.NewSunBasedCondition(viper.GetString("coop.opening.value"), viper.GetFloat64("coop.latitude"), viper.GetFloat64("coop.longitude"))
		if err != nil {
			return nil, nil, fmt.Errorf("error while creating the opening condition: %s", err)
		}

		openingCondition = oc
	default:
		return nil, nil, fmt.Errorf("error with the opening mode: %s", viper.GetString("coop.opening.mode"))
	}

	// Create the closing condition
	switch viper.GetString("coop.closing.mode") {
	case "time_based":
		cc, err := timebased.NewTimeBasedCondition(viper.GetString("coop.closing.value"))
		if err != nil {
			return nil, nil, fmt.Errorf("error while creating the closing condition: %s", err)
		}

		closingCondition = cc
	case "sun_based":
		cc, err := sunbased.NewSunBasedCondition(viper.GetString("coop.closing.value"), viper.GetFloat64("coop.latitude"), viper.GetFloat64("coop.longitude"))
		if err != nil {
			return nil, nil, fmt.Errorf("error while creating the closing condition: %s", err)
		}

		closingCondition = cc
	default:
		return nil, nil, fmt.Errorf("error with the closing mode: %s", viper.GetString("coop.closing.mode"))
	}

	return openingCondition, closingCondition, nil
}
