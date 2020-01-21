package routes

import (
	"net/http"

	"gocoop/internal/protocols"
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
		Message: "gocoop",
		Version: "1.0.0",
	}

	// Execute the template
	utils.JSONResponse(w, http.StatusOK, data)
}
