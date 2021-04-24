package controllers

import (
	"AmigoAutonomo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRequestLogin(ctx *gin.Context) {
	LoginService, err := services.NewLoginService()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, LoginService)
	}
}
