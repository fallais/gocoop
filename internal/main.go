package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"gocoop/internal/cache"
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

	// Initialize cache
	logrus.Infoln("Initializing the Redis cache")
	store, err := cache.NewRedisCache(viper.GetString("redis_host"), viper.GetString("redis_password"), 12*time.Hour)
	if err != nil {
		logrus.WithError(err).Fatalln("Error when initializing connection to Redis cache")
	}
	logrus.Infoln("Successfully initialized the Redis cache")

	// Initialize the middlewares
	logrus.Infoln("Initializing the middlewares")
	corsMiddleware := middleware.Cors()
	xRequestIDMiddleware := middleware.XRequestID()
	jwtMiddleware := middleware.JWT(store, viper.GetString("general.private_key"))
	logrus.Infoln("Successfully initialized the middlewares")

	// Initialize Web controllers
	logrus.Infoln("Initializing the services")
	coopService := services.NewCoopService(c)
	jwtService := services.NewJwtService(store, viper.GetString("general.private_key"))
	logrus.Infoln("Successfully initialized the services")

	// Initialize Web controllers
	logrus.Infoln("Initializing the Web controllers")
	coopCtrl := routes.NewCoopController(coopService)
	jwtCtrl := routes.NewJwtController(jwtService, viper.GetString("gui_username"), viper.GetString("gui_password"))
	logrus.Infoln("Successfully initialized the Web controllers")

	// Initialize CRON
	cr := cron.New()
	cr.AddFunc("@every 5m", c.Check)
	cr.Start()

	// Create a new Goji multiplexer
	root := goji.NewMux()

	// Middlewares
	root.Use(corsMiddleware)
	root.Use(xRequestIDMiddleware)

	// Unauthenticated route
	root.HandleFunc(pat.Post("/api/v1/login"), jwtCtrl.Login)
	root.HandleFunc(pat.Get("/api/v1/refresh"), jwtCtrl.Refresh)
	root.HandleFunc(pat.Get("/api/v1/logout"), jwtCtrl.Logout)

	// Authenticated routes
	authenticated := goji.SubMux()
	authenticated.Use(jwtMiddleware)
	authenticated.HandleFunc(pat.Get("/api/v1/coop"), coopCtrl.Get)
	//authenticated.HandleFunc(pat.Post("/api/v1/coop"), coopCtrl.Update)
	authenticated.HandleFunc(pat.Get("/api/v1/coop/status"), coopCtrl.GetStatus)
	authenticated.HandleFunc(pat.Post("/api/v1/coop/status"), coopCtrl.UpdateStatus)
	authenticated.HandleFunc(pat.Post("/api/v1/coop/open"), coopCtrl.Open)
	authenticated.HandleFunc(pat.Post("/api/v1/coop/close"), coopCtrl.Close)

	// Merge the muxes
	root.Handle(pat.New("/*"), authenticated)

	// Static files
	/* 	fs := http.FileServer(http.Dir(viper.GetString("static_dir")))
	   	root.Handle(pat.Get("/*"), fs) */

	// Handlers
	http.Handle("/", root)

	// Serve
	logrus.Infoln("Starting the Web server")
	http.ListenAndServe(":8000", root)
}
