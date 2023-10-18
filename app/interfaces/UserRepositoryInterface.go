package interfaces

import "github.com/uzzalhcse/GoSpire/app/models"

type UserRepositoryInterface interface {
	BaseRepositoryInterface // Embed the base repository interface
	GetUserByEmail(email string) (*models.User, error)
	UpdateProfile(user *models.User) error
}
