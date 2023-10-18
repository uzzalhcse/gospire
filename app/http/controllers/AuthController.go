package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/GoSpire/app/http/requests"
	"github.com/uzzalhcse/GoSpire/app/http/responses"
	"github.com/uzzalhcse/GoSpire/app/services"
)

type AuthController struct {
	authService services.AuthenticateService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: *services.NewAuthenticateService(),
	}
}

func (ctl AuthController) Login(c *gin.Context) {
	var request requests.LoginRequest
	if err := requests.Validate(c, &request); err != nil {
		responses.Error(c, err.Error, nil)
		return
	}
	// Check if the credentials are valid
	token, err := ctl.authService.Authenticate(request.Username, request.Password)
	if err != nil {
		responses.UnAuthorized(c, err.Error())
		return
	}
	responses.Success(c, "Login Successful", gin.H{
		"token": token,
		"type":  "Bearer",
	})
}
func (ctl AuthController) Register(c *gin.Context) {
	var request requests.RegisterRequest
	if err := requests.Validate(c, &request); err != nil {
		responses.Error(c, err.Error, nil)
		return
	}
	// Create the user
	token, err := ctl.authService.Register(request)
	if err != nil {
		responses.Error(c, err.Error(), nil)
		return
	}
	responses.Success(c, "Register Successful", gin.H{
		"token": token,
		"type":  "Bearer",
	})
}
