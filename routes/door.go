package routes

import (
	"fmt"
	"net/http"

	"gocoop-api/protocols"
	"gocoop-api/services"
	"gocoop-api/utils"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// DoorController is the controller of Door
type DoorController struct {
	doorService services.DoorService
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewDoorController returns a new DoorController
func NewDoorController(doorService services.DoorService) *DoorController {
	return &DoorController{
		doorService: doorService,
	}
}

//------------------------------------------------------------------------------
// Routes
//------------------------------------------------------------------------------

// Status of the door
func (ctrl *DoorController) Status(w http.ResponseWriter, r *http.Request) {
	// Get the status of the door
	status := ctrl.doorService.Status()

	// Response
	utils.JSONResponse(w, http.StatusOK, status)
}

// Use the door
func (ctrl *DoorController) Use(w http.ResponseWriter, r *http.Request) {
	// Retreive the parameters
	action := r.URL.Query().Get("action")

	// Process the action
	switch action {
	case "open":
		err := ctrl.doorService.Open()
		if err != nil {
			logrus.Errorln(err)

			// Prepare the data
			data := protocols.APIControllerResponse{
				ErrorID:          "DOOR/USE/001",
				ErrorMessage:     "Error when opening the door",
				ErrorDescription: fmt.Sprint(err),
			}

			// Execute the template
			utils.JSONResponse(w, http.StatusNotAcceptable, data)
			return
		}

		// Execute the template
		utils.JSONResponse(w, http.StatusNoContent, nil)
		break

	case "close":
		err := ctrl.doorService.Close()
		if err != nil {
			logrus.Errorln(err)

			// Prepare the data
			data := protocols.APIControllerResponse{
				ErrorID:          "DOOR/USE/002",
				ErrorMessage:     "Error when closing the door",
				ErrorDescription: fmt.Sprint(err),
			}

			// Execute the template
			utils.JSONResponse(w, http.StatusNotAcceptable, data)
			return
		}

		// Execute the template
		utils.JSONResponse(w, http.StatusNoContent, nil)
		break

	case "stop":
		err := ctrl.doorService.Stop()
		if err != nil {
			logrus.Errorln(err)

			// Prepare the data
			data := protocols.APIControllerResponse{
				ErrorID:          "DOOR/USE/003",
				ErrorMessage:     "Error when stopping the door",
				ErrorDescription: fmt.Sprint(err),
			}

			// Execute the template
			utils.JSONResponse(w, http.StatusNotAcceptable, data)
			return
		}

		// Execute the template
		utils.JSONResponse(w, http.StatusNoContent, nil)
		break

	default:
		logrus.Errorln("The action is not correct")

		// Prepare the data
		data := protocols.APIControllerResponse{
			ErrorID:      "DOOR/USE/003",
			ErrorMessage: "The action is not correct",
		}

		// Execute the template
		utils.JSONResponse(w, http.StatusNotAcceptable, data)
		return
	}
}
