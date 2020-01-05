package routes

import (
	"fmt"
	"net/http"

	"gocoop/internal/protocols"
	"gocoop/internal/services"

	"github.com/alioygur/gores"
	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// CoopController is the controller of Coop
type CoopController struct {
	coopService services.CoopService
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewCoopController returns a new CoopController
func NewCoopController(coopService services.CoopService) *CoopController {
	return &CoopController{
		coopService: coopService,
	}
}

//------------------------------------------------------------------------------
// Routes
//------------------------------------------------------------------------------

// Status of the coop
func (ctrl *CoopController) Status(w http.ResponseWriter, r *http.Request) {
	// Get the status of the coop
	status := ctrl.coopService.Status()

	// Response
	gores.JSON(w, http.StatusOK, status)
}

// Use the coop
func (ctrl *CoopController) Use(w http.ResponseWriter, r *http.Request) {
	// Retreive the parameters
	action := r.URL.Query().Get("action")

	// Process the action
	switch action {
	case "open":
		err := ctrl.coopService.Open()
		if err != nil {
			logrus.Errorln(err)

			// Prepare the data
			data := protocols.APIControllerResponse{
				ErrorID:          "DOOR/USE/001",
				ErrorMessage:     "Error when opening the coop",
				ErrorDescription: fmt.Sprint(err),
			}

			// Execute the template
			gores.JSON(w, http.StatusNotAcceptable, data)
			return
		}

		// Execute the template
		gores.JSON(w, http.StatusNoContent, nil)
		break

	case "close":
		err := ctrl.coopService.Close()
		if err != nil {
			logrus.Errorln(err)

			// Prepare the data
			data := protocols.APIControllerResponse{
				ErrorMessage:     "Error when closing the coop",
				ErrorDescription: err.Error(),
			}

			// Execute the template
			gores.JSON(w, http.StatusNotAcceptable, data)
			return
		}

		// Execute the template
		gores.JSON(w, http.StatusNoContent, nil)
		break

	default:
		logrus.Errorln("The action is not correct")

		// Prepare the data
		data := protocols.APIControllerResponse{
			ErrorMessage: "The action is not correct",
		}

		// Execute the template
		gores.JSON(w, http.StatusNotAcceptable, data)
		return
	}
}
