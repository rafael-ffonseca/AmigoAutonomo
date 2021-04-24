package services

import (
	"AmigoAutonomo/config"
	"AmigoAutonomo/infrastructures"
	"AmigoAutonomo/interfaces"
	"AmigoAutonomo/repositories"
	"errors"
	"fmt"
)

type LoginServices struct {
	interfaces.ILoginServices
	repository interfaces.IUserRepository
}

func NewLoginService() (interfaces.ILoginServices, error) {
	db, err := infrastructures.GetDatabase(config.GetPropFromServiceSection().Homologation)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error trying get DB: %s", err.Error()))
	}
	userRepo, err := repositories.NewUserRepository(db)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error trying get User repository: %s", err.Error()))
	}

	return &LoginServices{
		repository: userRepo,
	}, nil
}

func (service *LoginServices) DoLogin() error {
	return nil
}
