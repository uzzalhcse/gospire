package services

import (
	"github.com/uzzalhcse/GoSpire/app/interfaces"
	"github.com/uzzalhcse/GoSpire/app/models"
)

type UserService struct {
	userRepository interfaces.UserRepositoryInterface
}

func NewUserService(userRepo interfaces.UserRepositoryInterface) *UserService {
	return &UserService{userRepository: userRepo}
}

func (us *UserService) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	err := us.userRepository.GetAll(&users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (us *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := us.userRepository.GetByID(&user, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) GetUserBySponsor(SponsorID uint) (*models.User, error) {
	var user models.User
	err := us.userRepository.GetFirst(&user, "sponsor_id = ?", SponsorID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	return us.userRepository.GetUserByEmail(email)
}

func (us *UserService) UpdateProfile(user *models.User) error {
	return us.userRepository.UpdateProfile(user)
}

func (us *UserService) Create(data interface{}) error {
	return us.userRepository.Create(data)
}

func (us *UserService) Delete(data interface{}) error {
	return us.userRepository.Delete(data)
}
