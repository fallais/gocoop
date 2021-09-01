package routes

import (
	"embed"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/alioygur/gores"
	"github.com/fallais/gocoop/internal/protocols"
	"github.com/fallais/gocoop/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// TemplatesFS ...
var TemplatesFS embed.FS

// MiscController is the controller of Misc.
type MiscController struct {
	coopService services.CoopService
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewMiscController returns a new MiscController.
func NewMiscController(coopService services.CoopService) *MiscController {
	return &MiscController{
		coopService: coopService,
	}
}

//------------------------------------------------------------------------------
// Routes
//------------------------------------------------------------------------------

// Index is the index page.
func (ctrl *MiscController) Index(w http.ResponseWriter, r *http.Request) {
	// Get the coop
	coop := ctrl.coopService.Get()

	// Prepare the response
	response := protocols.CoopResponse{
		OpeningCondition: map[string]string{
			"mode":  coop.OpeningCondition().Mode(),
			"value": coop.OpeningCondition().Value(),
		},
		ClosingCondition: map[string]string{
			"mode":  coop.ClosingCondition().Mode(),
			"value": coop.ClosingCondition().Value(),
		},
		OpeningTime: coop.OpeningTime(),
		ClosingTime: coop.ClosingTime(),
		Latitude:    coop.Latitude(),
		Longitude:   coop.Longitude(),
		Status:      string(coop.Status()),
		IsAutomatic: coop.IsAutomatic(),
	}

	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(TemplatesFS, "templates/index.html.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// Header
	w.Header().Add("Content-Type", "text/html")

	// Execute
	t.Execute(w, response)
}

// Configuration is the configuration page.
func (ctrl *MiscController) Configuration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctrl.getConfiguration(w, r)
	case "POST":
		ctrl.updateConfiguration(w, r)
	default:
	}
}

func (ctrl *MiscController) getConfiguration(w http.ResponseWriter, r *http.Request) {
	// Get the coop
	coop := ctrl.coopService.Get()

	// Prepare the response
	response := protocols.CoopResponse{
		OpeningCondition: map[string]string{
			"mode":  coop.OpeningCondition().Mode(),
			"value": coop.OpeningCondition().Value(),
		},
		ClosingCondition: map[string]string{
			"mode":  coop.ClosingCondition().Mode(),
			"value": coop.ClosingCondition().Value(),
		},
		OpeningTime: coop.OpeningTime(),
		ClosingTime: coop.ClosingTime(),
		Latitude:    coop.Latitude(),
		Longitude:   coop.Longitude(),
		Status:      string(coop.Status()),
		IsAutomatic: coop.IsAutomatic(),
	}

	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(TemplatesFS, "templates/configuration.html.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// Header
	w.Header().Add("Content-Type", "text/html")

	// Execute
	t.Execute(w, response)
}

func (ctrl *MiscController) updateConfiguration(w http.ResponseWriter, r *http.Request) {
	// Parse the form
	err := r.ParseForm()
	if err != nil {
		logrus.WithError(err).Errorln("error while parsing the form")
		return
	}

	// Parse the automatic mode (ignore the error)
	isAutomatic, _ := strconv.ParseBool(r.FormValue("is_automatic"))

	// Parse the latitude
	latitude, _ := strconv.ParseFloat(strings.TrimSpace(r.FormValue("latitude")), 64)
	longitude, _ := strconv.ParseFloat(strings.TrimSpace(r.FormValue("longitude")), 64)

	// Create the request
	update := protocols.CoopUpdateRequestController{
		Status:      r.FormValue("status"),
		Latitude:    latitude,
		Longitude:   longitude,
		IsAutomatic: isAutomatic,
		OpeningCondition: protocols.Condition{
			Mode:  r.FormValue("opening_mode"),
			Value: r.FormValue("opening_value"),
		},
		ClosingCondition: protocols.Condition{
			Mode:  r.FormValue("closing_mode"),
			Value: r.FormValue("closing_value"),
		},
	}

	// Update the coop
	err = ctrl.coopService.Update(update)
	if err != nil {
		logrus.WithError(err).Errorln("error while updating the coop")
		return
	}

	// Get the coop
	coop := ctrl.coopService.Get()

	// Prepare the response
	response := protocols.CoopResponse{
		OpeningCondition: map[string]string{
			"mode":  coop.OpeningCondition().Mode(),
			"value": coop.OpeningCondition().Value(),
		},
		ClosingCondition: map[string]string{
			"mode":  coop.ClosingCondition().Mode(),
			"value": coop.ClosingCondition().Value(),
		},
		OpeningTime: coop.OpeningTime(),
		ClosingTime: coop.ClosingTime(),
		Latitude:    coop.Latitude(),
		Longitude:   coop.Longitude(),
		Status:      string(coop.Status()),
		IsAutomatic: coop.IsAutomatic(),
	}

	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(TemplatesFS, "templates/configuration.html.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// Header
	w.Header().Add("Content-Type", "text/html")

	// Execute
	t.Execute(w, response)
}

// Cameras returns all the cameras.
func (ctrl *MiscController) Cameras(w http.ResponseWriter, r *http.Request) {
	cameras := viper.GetStringMapString("cameras")

	// Execute the template
	gores.JSON(w, http.StatusOK, cameras)
}
