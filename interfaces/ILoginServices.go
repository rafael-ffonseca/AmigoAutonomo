package interfaces

import (
	"AmigoAutonomo/models"
)

type ILoginServices interface {
	DoLogin() (*models.GetRequestLoginOutput, error)
}