package repositories

import (
	"AmigoAutonomo/entities"
	"AmigoAutonomo/interfaces"
	"gorm.io/gorm"
)

type UserRepository struct {
	interfaces.IUserRepository
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) (interfaces.IUserRepository, error) {
	return &UserRepository{
		database: db,
	}, nil
}

func (repository *UserRepository) GetUserLogin(UserLogin string, UserPass string) (*entities.Users, error) {
	return nil, nil
}