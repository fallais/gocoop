package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"gocoop/internal/routes"
	"gocoop/internal/services"
	"gocoop/internal/system/middleware"
	"gocoop/pkg/coop"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"goji.io"
	"goji.io/pat"
)

func Run(cmd *cobra.Command, args []string) {
	// Flags
	configFile, err := cmd.Flags().GetString("config_file")
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

	// Create the coop instance
	c, err := coop.New()
	if err != nil {
		logrus.WithError(err).Fatalln("Error while creating the coop instance")
	}

	// Initialize the middlewares
	logrus.Infoln("Initializing the middlewares")
	corsMiddleware := middleware.Cors()
	xRequestIDMiddleware := middleware.XRequestID()
	logrus.Infoln("Successfully initialized the middlewares")

	// Initialize Web controllers
	logrus.Infoln("Initializing the services")
	coopService := services.NewCoopService(c)
	logrus.Infoln("Successfully initialized the services")

	// Initialize Web controllers
	logrus.Infoln("Initializing the Web controllers")
	coopCtrl := routes.NewCoopController(coopService)
	miscCtrl := routes.NewMiscController()
	logrus.Infoln("Successfully initialized the Web controllers")

	// Initialize CRON
	c.Check()
	cr := cron.New()
	cr.AddFunc("@every 5m", c.Check)
	cr.Start()

	// Create a new Goji multiplexer
	root := goji.NewMux()

	// Middlewares
	logrus.Infoln("Initializing the middlewares")
	root.Use(corsMiddleware)
	root.Use(xRequestIDMiddleware)
	logrus.Infoln("Successfully initialized the middlewares")

	// Define the routes for Web
	fs := http.FileServer(http.Dir("web"))
	root.Handle(pat.Get("/"), fs)
	root.Handle(pat.Get("/app/*"), fs)
	root.Handle(pat.Get("/static/*"), fs)

	// Define the routes for API
	root.HandleFunc(pat.Get("/api"), miscCtrl.Hello)
	root.HandleFunc(pat.Get("/api/v1/configuration"), miscCtrl.Configuration)
	root.HandleFunc(pat.Get("/api/v1/door"), coopCtrl.Status)
	root.HandleFunc(pat.Get("/api/v1/door/use"), coopCtrl.Use)

	// Serve
	logrus.Infoln("Starting the Web server")
	http.ListenAndServe(":8000", root)
}
