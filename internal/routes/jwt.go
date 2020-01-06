package routes

import (
	"fmt"
	"net/http"
	"strings"

	"gocoop/internal/protocols"
	"gocoop/internal/services"
	"gocoop/internal/utils"

	"github.com/alioygur/gores"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// JwtController is the structure of the JWT controller
type JwtController struct {
	jwtService services.JwtService
	username   string
	password   string
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewJwtController returns a JWT controller
func NewJwtController(jwtService services.JwtService, username, password string) *JwtController {
	return &JwtController{
		jwtService: jwtService,
		username:   username,
		password:   password,
	}
}

//------------------------------------------------------------------------------
// Function
//------------------------------------------------------------------------------

// Login a user to the API
func (ctrl *JwtController) Login(w http.ResponseWriter, r *http.Request) {
	var input protocols.JWTokenRequest

	// Parse the request
	err := utils.ParseRequest(r, &input)
	if err != nil {
		response := &protocols.APIControllerResponse{
			ErrorID:      "JWT/LOGIN/001",
			ErrorMessage: "Unable to parse the request",
		}

		// Publish the respons
		gores.JSON(w, http.StatusNotAcceptable, response)
		return
	}

	// Check the email
	if len(strings.TrimSpace(input.Username)) == 0 {
		response := &protocols.APIControllerResponse{
			ErrorID:      "JWT/LOGIN/002",
			ErrorMessage: "Username cannot be empty",
		}

		// Publish the respons
		gores.JSON(w, http.StatusPreconditionFailed, response)
		return
	}

	// Check the password
	if len(strings.TrimSpace(input.Password)) == 0 {
		response := &protocols.APIControllerResponse{
			ErrorID:      "JWT/LOGIN/003",
			ErrorMessage: "Password cannot be empty",
		}

		// Publish the respons
		gores.JSON(w, http.StatusPreconditionFailed, response)
		return
	}

	// Authenticate the user
	if input.Username != ctrl.username || input.Password != ctrl.password {
		response := &protocols.APIControllerResponse{
			ErrorID:      "JWT/LOGIN/004",
			ErrorMessage: fmt.Sprintf("Unable to authenticate the user"),
		}

		// Publish the respons
		gores.JSON(w, http.StatusNotAcceptable, response)
		return
	}

	// Create the token
	token, err := ctrl.jwtService.Create(input.Username)
	if err != nil {
		response := &protocols.APIControllerResponse{
			ErrorID:      "JWT/LOGIN/005",
			ErrorMessage: fmt.Sprintf("Unable to create the token : %s", err),
		}

		// Publish the respons
		gores.JSON(w, http.StatusNotAcceptable, response)
		return
	}

	// Publish the response
	gores.JSON(w, http.StatusOK, token)
}

// Refresh a JWT token
func (ctrl *JwtController) Refresh(w http.ResponseWriter, r *http.Request) {
}

// Logout a user of the API
func (ctrl *JwtController) Logout(w http.ResponseWriter, r *http.Request) {
}
