package routes

import (
	"net/http"
	"strings"

	"github.com/fallais/gocoop/internal/protocols"
	"github.com/fallais/gocoop/internal/services"
	"github.com/fallais/gocoop/internal/utils"

	"github.com/alioygur/gores"
	"github.com/sirupsen/logrus"
)

// ErrParseRequest is raised when the request parsing fails.
const ErrParseRequest = "Error while parsing the request"

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

// Get returns the coop.
func (ctrl *CoopController) Get(w http.ResponseWriter, r *http.Request) {
	// Get the coop
	coop := ctrl.coopService.Get()

	// Prepare the response
	response := &protocols.CoopResponse{
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

	// Response
	gores.JSON(w, http.StatusOK, response)
}

// Update updates the settings of the coop.
func (ctrl *CoopController) Update(w http.ResponseWriter, r *http.Request) {
	var input protocols.CoopUpdateRequestController

	// Decode the request
	err := utils.ParseRequest(r, &input)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"input": input,
		}).WithError(err).Errorln(ErrParseRequest)

		// Prepare the template
		response := protocols.APIControllerResponse{
			ErrorMessage:     ErrParseRequest,
			ErrorDescription: err.Error(),
		}

		// Execute the template
		gores.JSON(w, http.StatusNotAcceptable, response)
		return
	}

	// Check the status
	if len(strings.TrimSpace(input.Status)) == 0 {
		response := &protocols.APIControllerResponse{
			ErrorMessage: "Status cannot be blank",
		}

		// Execute the template
		gores.JSON(w, http.StatusPreconditionFailed, response)
		return
	}

	// Check the opening condition mode
	if len(strings.TrimSpace(input.OpeningCondition.Mode)) == 0 {
		response := &protocols.APIControllerResponse{
			ErrorMessage: "Opening condition mode cannot be blank",
		}

		// Execute the template
		gores.JSON(w, http.StatusPreconditionFailed, response)
		return
	}

	// Check the opening condition value
	if len(strings.TrimSpace(input.OpeningCondition.Value)) == 0 {
		response := &protocols.APIControllerResponse{
			ErrorMessage: "Opening condition value cannot be blank",
		}

		// Execute the template
		gores.JSON(w, http.StatusPreconditionFailed, response)
		return
	}

	// Update the the coop
	err = ctrl.coopService.Update(input)
	if err != nil {
		response := &protocols.APIControllerResponse{
			ErrorMessage:     "Unable to update the coop",
			ErrorDescription: err.Error(),
		}

		// Publish the respons
		gores.JSON(w, http.StatusInternalServerError, response)
		return
	}

	// Response
	gores.JSON(w, http.StatusNoContent, nil)
}

// Open the coop
func (ctrl *CoopController) Open(w http.ResponseWriter, r *http.Request) {
	err := ctrl.coopService.Open()
	if err != nil {
		logrus.Errorln(err)

		// Prepare the data
		data := protocols.APIControllerResponse{
			ErrorMessage:     "Error when opening the coop",
			ErrorDescription: err.Error(),
		}

		// Execute the template
		gores.JSON(w, http.StatusNotAcceptable, data)
		return
	}

	// Execute the template
	gores.JSON(w, http.StatusNoContent, nil)

}

// Close the coop
func (ctrl *CoopController) Close(w http.ResponseWriter, r *http.Request) {
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
}
