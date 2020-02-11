package routes

import (
	"net/http"

	"github.com/fallais/gocoop/internal/protocols"

	"github.com/alioygur/gores"
	"github.com/spf13/viper"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// MiscController is the controller of Misc.
type MiscController struct{}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewMiscController returns a new MiscController.
func NewMiscController() *MiscController {
	return &MiscController{}
}

//------------------------------------------------------------------------------
// Routes
//------------------------------------------------------------------------------

// Hello is the welcome message.
func (ctrl *MiscController) Hello(w http.ResponseWriter, r *http.Request) {
	// Prepare the data
	data := protocols.HelloResponse{
		Message: "gocoop",
		Version: "1.0.0",
	}

	// Execute the template
	gores.JSON(w, http.StatusOK, data)
}

// Cameras returns all the cameras.
func (ctrl *MiscController) Cameras(w http.ResponseWriter, r *http.Request) {
	cameras := viper.GetStringMapString("cameras")

	// Execute the template
	gores.JSON(w, http.StatusOK, cameras)
}
