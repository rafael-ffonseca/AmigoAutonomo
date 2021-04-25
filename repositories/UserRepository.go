package repositories

import (
	"AmigoAutonomo/entities"
	"AmigoAutonomo/interfaces"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository struct {
	interfaces.IUserRepository
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) (interfaces.IUserRepository, error) {
	return &UserRepository{
		DB: db,
	}, nil
}

func (repository *UserRepository) GetUserByLogin(username string, password string) (*entities.Users, error) {
	if username == "" || password == "" {
		return nil, errors.New(fmt.Sprintf("Username or Password does not match"))
	}
	var User *entities.Users
	repository.DB.Where(&entities.Users{Username: username, Password: password}).First(&User)
	if User.ID > 0 {
		return User, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Username or Password does not match"))
	}
}

func (repository *UserRepository) CreateUser(*entities.Users) error {
	return nil
}