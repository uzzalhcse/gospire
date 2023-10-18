package services

import (
	"errors"
	"github.com/uzzalhcse/GoSpire/app/helpers"
	"github.com/uzzalhcse/GoSpire/app/http/requests"
	"github.com/uzzalhcse/GoSpire/app/interfaces"
	"github.com/uzzalhcse/GoSpire/app/models"
	"github.com/uzzalhcse/GoSpire/app/repositories"
	"github.com/uzzalhcse/GoSpire/bootstrap/app"
	"github.com/uzzalhcse/GoSpire/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateService struct {
	userRepository interfaces.UserRepositoryInterface
}

func NewAuthenticateService() *AuthenticateService {
	return &AuthenticateService{
		userRepository: repositories.NewUserRepository(app.DB),
	}
}

func (as *AuthenticateService) Authenticate(username, password string) (string, error) {
	var user models.User
	err := as.userRepository.GetFirst(&user, "email = ? OR phone = ?", username, username)
	if err != nil {
		return "", err
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("incorrect password")
	}

	// Generate a new JWT token for the authenticated user
	token, err := helpers.GenerateToken(user, config.Jwt.SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (as *AuthenticateService) Register(request requests.RegisterRequest) (string, error) {
	// Check if the email or phone already exists
	var existingUser models.User
	err := as.userRepository.GetFirst(&existingUser, "email = ? OR phone = ?", request.Email, request.Phone)
	if err == nil {
		return "", errors.New("email or phone already exists")
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}
	// Create the user
	err = as.userRepository.Create(&user)
	if err != nil {
		return "", err
	}
	// Generate a new JWT token for the authenticated user
	token, err := helpers.GenerateToken(user, config.Jwt.SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
