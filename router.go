package main

import (
	"AmigoAutonomo/controllers"
	"github.com/gin-gonic/gin"
)

const GetRequestLoginRoute = "/login"


func InitRoutes(route *gin.Engine) {
	route.GET(GetRequestLoginRoute, controllers.GetRequestLogin)
}
