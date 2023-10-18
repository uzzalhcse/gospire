package repositories

import (
	"errors"
	"github.com/uzzalhcse/GoSpire/app/interfaces"
	"github.com/uzzalhcse/GoSpire/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository // Embed the BaseRepository by reference
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{BaseRepository: NewBaseRepository(db)}
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) UpdateProfile(user *models.User) error {
	result := repo.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no record was updated")
	}
	return nil
}
