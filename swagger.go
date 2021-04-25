package main

import (
	"AmigoAutonomo/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwagger(route *gin.Engine) {

	docs.SwaggerInfo.Title = "Amigo Autônomo App"
	docs.SwaggerInfo.Description = "Amigo Autônomo - Backend API"
	docs.SwaggerInfo.Host = "localhost:1900"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = ""

	route.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
}
