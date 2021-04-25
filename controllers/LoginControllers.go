package controllers

import (
	"AmigoAutonomo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)


// GetRequestLogin godoc
// @Summary Request to login user
// @Description Authenticate and Provides User JWT Token
// @Produce json
// @Param input query models.GetRequestLoginInput true "input"
// @Success 200 {object} models.GetRequestLoginOutput
// @Failure 400 {string} error "Username or Password does not match"
// @Router /login [get]
func GetRequestLogin(ctx *gin.Context) {
	LoginService, err := services.NewLoginService(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		Output, err := LoginService.DoLogin()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, Output)
		}
	}
}
