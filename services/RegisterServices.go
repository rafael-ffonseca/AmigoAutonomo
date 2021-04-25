package services

import (
	"AmigoAutonomo/config"
	"AmigoAutonomo/entities"
	"AmigoAutonomo/infrastructures"
	"AmigoAutonomo/interfaces"
	"AmigoAutonomo/models"
	"AmigoAutonomo/repositories"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type RegisterServices struct {
	interfaces.IRegisterServices
	repository interfaces.IUserRepository
	Context *gin.Context
	Request *http.Request
}

func NewRegisterServices(ctx *gin.Context) (interfaces.IRegisterServices, error) {
	db, err := infrastructures.GetDatabase(config.GetPropFromServiceSection().Homologation)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error trying get DB: %s", err.Error()))
	}
	userRepo, err := repositories.NewUserRepository(db)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error trying get User repository: %s", err.Error()))
	}

	return &RegisterServices{
		repository: userRepo,
		Context: ctx,
		Request: ctx.Request,
	}, nil
}

func (service *RegisterServices) RegisterUser() (*models.PostRegisterUserOutput, error) {
	input, err := newPostRegisterUserInput(service.Context)
	if err != nil {
		return nil, err
	}

	err = ValidateRegisterUserInput(input)
	if err != nil {
		return nil, err
	}

	newUser := &entities.Users{

	}
	err = service.repository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return &models.PostRegisterUserOutput{

	}, nil
}

func newPostRegisterUserInput(ctx *gin.Context) (*models.PostRegisterUserInput, error) {
	var returnedInput *models.PostRegisterUserInput
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error trying read request body on method GetRequestLoginInput: %s", err.Error()))
	}

	err = json.Unmarshal(body, &returnedInput)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error trying read request body on method GetRequestLoginInput: %s. Request: %s", err.Error(), string(body)))
	}

	return returnedInput, nil
}

func ValidateRegisterUserInput(input *models.PostRegisterUserInput) error {

	return nil
}
