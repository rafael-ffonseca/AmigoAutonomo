package interfaces

import "AmigoAutonomo/entities"

type IUserRepository interface {
	GetUserByLogin(username string, password string) (*entities.Users, error)
	CreateUser(*entities.Users) error
}
