package services

import "main/internal/models"

type UserService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *UserService {
	return &UserService{userStorage}
}

type UserStorage interface {
	CreateUser(user *models.User) error
	GetUser(id int) (*models.User, error)
}

func (userService *UserService) CreateUser(user *models.User) error {
	return userService.userStorage.CreateUser(user)
}

func (userService *UserService) GetUser(id int) (*models.User, error) {
	user, err := userService.userStorage.GetUser(id)
	return user, err
}
