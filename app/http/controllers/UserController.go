package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/GoSpire/app/http/responses"
	"github.com/uzzalhcse/GoSpire/app/models"
	"github.com/uzzalhcse/GoSpire/app/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		responses.Error(c, err.Error(), nil)
		return
	}
	responses.Success(c, "All Users", users)
}

func (uc *UserController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := uc.userService.GetUserByEmail(email)
	if err != nil {
		responses.Error(c, err.Error(), nil)
		return
	}
	responses.Success(c, "User", user)
}

func (uc *UserController) UpdateProfile(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		responses.Error(c, err.Error(), nil)
		return
	}
	if err := uc.userService.UpdateProfile(&user); err != nil {
		responses.Error(c, err.Error(), nil)
		return
	}
	responses.Success(c, "Profile Updated", nil)
}
