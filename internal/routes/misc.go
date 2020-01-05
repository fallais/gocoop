package routes

import (
	"net/http"

	"gocoop/internal/protocols"
	"gocoop/internal/shared"
	"gocoop/internal/utils"
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

// Hello is the welcome message
func (ctrl *MiscController) Hello(w http.ResponseWriter, r *http.Request) {
	// Prepare the data
	data := protocols.HelloResponse{
		Message: "Go Coop API",
		Version: "1",
	}

	// Execute the template
	utils.JSONResponse(w, http.StatusOK, data)
}

// Configuration is the actual configuration
func (ctrl *MiscController) Configuration(w http.ResponseWriter, r *http.Request) {
	// Prepare the data
	data := protocols.Configuration{
		Latitude:  shared.Latitude,
		Longitude: shared.Longitude,
	}

	// Execute the template
	utils.JSONResponse(w, http.StatusOK, data)
}
