package main

import (
	"flag"
	"net/http"
	"time"

	"gocoop-api/coop"
	"gocoop-api/raspberry/door"
	"gocoop-api/routes"
	"gocoop-api/services"
	"gocoop-api/system/middleware"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"goji.io"
	"goji.io/pat"
)

var (
	logging           = flag.String("logging", "info", "Logging level")
	configurationFile = flag.String("configuration_file", "configuration.json", "Configuration file")
)

func init() {
	// Parse the flags
	flag.Parse()

	// Set localtime to UTC
	time.Local = time.UTC

	// Set the logging level
	level, err := logrus.ParseLevel(*logging)
	if err != nil {
		logrus.Fatalln("Invalid log level ! (panic, fatal, error, warn, info, debug)")
	}
	logrus.SetLevel(level)

	// Set the TextFormatter
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	logrus.Infoln("gocoop-api is starting")
}

func main() {
	// Create the coop instance
	c, err := coop.New(*configurationFile)
	if err != nil {
		logrus.Fatalf("Error while creating the coop instance : %s", err)
	}

	// Create the door
	d := door.New()

	// Initialize the middlewares
	logrus.Infoln("Initializing the middlewares")
	corsMiddleware := middleware.Cors()
	xRequestIDMiddleware := middleware.XRequestID()
	logrus.Infoln("Successfully initialized the middlewares")

	// Initialize Web controllers
	logrus.Infoln("Initializing the services")
	doorService := services.NewDoorService(d)
	weatherService := services.NewWeatherService()
	logrus.Infoln("Successfully initialized the services")

	// Initialize Web controllers
	logrus.Infoln("Initializing the Web controllers")
	doorCtrl := routes.NewDoorController(doorService)
	weatherCtrl := routes.NewWeatherController(weatherService)
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
	root.HandleFunc(pat.Get("/api"), miscCtrl.Hello)
	root.HandleFunc(pat.Get("/api/v1/configuration"), miscCtrl.Configuration)
	root.HandleFunc(pat.Get("/api/v1/door"), doorCtrl.Status)
	root.HandleFunc(pat.Get("/api/v1/door/use"), doorCtrl.Use)
	root.HandleFunc(pat.Get("/api/v1/weather/sunrise"), weatherCtrl.GetSunrise)
	root.HandleFunc(pat.Get("/api/v1/weather/sunset"), weatherCtrl.GetSunset)

	// Serve
	logrus.Infoln("Starting the Web server")
	http.ListenAndServe(":8000", root)
}
