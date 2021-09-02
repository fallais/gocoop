package system

import (
	"fmt"

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
