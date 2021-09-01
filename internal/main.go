package internal

import (
	"bytes"
	"embed"
	"io/ioutil"
	"net/http"

	"github.com/fallais/gocoop/internal/routes"
	"github.com/fallais/gocoop/internal/services"
	"github.com/fallais/gocoop/internal/system"
	"github.com/fallais/gocoop/pkg/coop"
	"github.com/fallais/gocoop/pkg/door"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// StaticFS is the embed for the static files.
var StaticFS embed.FS

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	// Flags
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		logrus.WithError(err).Fatalln("Error while getting the flag for configuration data")
	}

	// Read configuration file
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		logrus.WithError(err).Fatalln("Error while reading configuration file")
	}

	// Initialize configuration values with Viper
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		logrus.WithError(err).Fatalln("Error when reading configuration data")
	}

	// Create the opening condition
	logrus.Infoln("Creating the opening and closing conditions")
	openingCondition, closingCondition, err := system.SetupConditions()
	if err != nil {
		logrus.WithError(err).Fatalln("Error while creating the opening and closing conditions")
	}
	logrus.Infoln("Successfully created the opening and closing conditions")

	// Door
	logrus.WithFields(logrus.Fields{
		"pin_1A":      viper.GetString("door.pin_1A"),
		"pin_1B":      viper.GetString("door.pin_1B"),
		"pin_enable1": viper.GetString("door.pin_enable1"),
	}).Infoln("Creating the door")
	d := door.NewDoor(viper.GetInt("door.pin_1A"), viper.GetInt("door.pin_1B"), viper.GetInt("door.pin_enable1"), viper.GetDuration("door.opening_duration"), viper.GetDuration("door.closing_duration"))
	logrus.Infoln("Successfully created the door")

	// Notifiers
	notifs := system.SetupNotifiers()

	// Create the coop instance
	c, err := coop.New(viper.GetFloat64("coop.latitude"), viper.GetFloat64("coop.longitude"), d, openingCondition, closingCondition, coop.WithAutomatic(), coop.WithNotifiers(notifs))
	if err != nil {
		logrus.WithError(err).Fatalln("Error while creating the coop instance")
	}

	// Initialize Web controllers
	logrus.Infoln("Initializing the services")
	coopService := services.NewCoopService(c)
	logrus.Infoln("Successfully initialized the services")

	// Initialize Web controllers
	logrus.Infoln("Initializing the Web controllers")
	//coopCtrl := routes.NewCoopController(coopService)
	miscCtrl := routes.NewMiscController(coopService)
	//securityCtrl := routes.NewSecurityController(viper.GetString("general.gui_username"), viper.GetString("general.gui_password"))
	logrus.Infoln("Successfully initialized the Web controllers")

	// Static files
	var staticFS = http.FS(StaticFS)
	fs := http.FileServer(staticFS)

	// Handlers
	http.Handle("/static/", fs)
	http.HandleFunc("/", miscCtrl.Index)
	http.HandleFunc("/configuration", miscCtrl.Configuration)

	// Serve
	logrus.Infoln("Starting the Web server")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		logrus.WithError(err).Fatalln("Error while starting the Web server")
	}
}
