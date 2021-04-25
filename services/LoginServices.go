package services

import (
	"AmigoAutonomo/config"
	"AmigoAutonomo/infrastructures"
	"AmigoAutonomo/interfaces"
	"AmigoAutonomo/models"
	"AmigoAutonomo/repositories"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginServices struct {
	interfaces.ILoginServices
	repository interfaces.IUserRepository
	Context *gin.Context
	Request *http.Request
}

func NewLoginService(ctx *gin.Context) (interfaces.ILoginServices, error) {
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
		Context: ctx,
		Request: ctx.Request,
	}, nil
}

func (service *LoginServices) DoLogin() (*models.GetRequestLoginOutput, error) {
	input, err := newGetRequestLoginInput(service.Context)
	if err != nil {
		return nil, err
	}
	resultLogin, err := service.repository.GetUserByLogin(input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	return &models.GetRequestLoginOutput{
		ID:        resultLogin.ID,
		Username:  resultLogin.Username,
		FirstName: resultLogin.FirstName,
		LastName:  resultLogin.LastName,
		Email:     resultLogin.Email,
	}, nil
}

func newGetRequestLoginInput(ctx *gin.Context) (*models.GetRequestLoginInput, error) {
	returnedInput := &models.GetRequestLoginInput{}

	if Username, ok := ctx.GetQuery("username"); ok && Username != "/" {
		returnedInput.Username = Username
	} else if Username, ok := ctx.GetQuery("Username"); ok && Username != "/" {
		returnedInput.Username = Username
	}

	if Password, ok := ctx.GetQuery("password"); ok && Password != "/" {
		returnedInput.Password = Password
	} else if Password, ok := ctx.GetQuery("Password"); ok && Password != "/" {
		returnedInput.Password = Password
	}

	return returnedInput, nil
}
