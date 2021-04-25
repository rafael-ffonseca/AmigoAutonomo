package interfaces

import "AmigoAutonomo/models"

type IRegisterServices interface {
	RegisterUser() (*models.PostRegisterUserOutput, error)
}
