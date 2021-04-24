package interfaces

import "AmigoAutonomo/entities"

type IUserRepository interface {
	GetUserLogin(UserLogin string, UserPass string) (*entities.Users, error)
}
