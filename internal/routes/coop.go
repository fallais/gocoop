package routes

import (
	"net/http"

	"gocoop/internal/protocols"
	"gocoop/internal/services"
	"gocoop/internal/utils"

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
		Latitude:    coop.Latitude(),
		Longitude:   coop.Longitude(),
		Status:      string(coop.Status()),
		IsAutomatic: coop.IsAutomatic(),
	}

	// Response
	gores.JSON(w, http.StatusOK, response)
}

// GetStatus returns the status of the coop
func (ctrl *CoopController) GetStatus(w http.ResponseWriter, r *http.Request) {
	// Get the status of the coop
	status := ctrl.coopService.GetStatus()

	// Response
	gores.JSON(w, http.StatusOK, status)
}

// UpdateStatus updates of the coop
func (ctrl *CoopController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var input protocols.Status

	// Parse the request
	err := utils.ParseRequest(r, &input)
	if err != nil {
		response := &protocols.APIControllerResponse{
			ErrorMessage: "Unable to parse the request",
		}

		// Publish the respons
		gores.JSON(w, http.StatusNotAcceptable, response)
		return
	}

	// Update the status of the coop
	err = ctrl.coopService.UpdateStatus(input.Status)
	if err != nil {
		response := &protocols.APIControllerResponse{
			ErrorMessage: "Unable to update the status",
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
