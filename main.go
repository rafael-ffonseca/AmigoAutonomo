package main

import (
	"AmigoAutonomo/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
	}))
	
	InitRoutes(router)
	InitSwagger(router)

	servicePort := config.GetPropFromServiceSection().ServicePort

	log.Fatal(router.Run(":" + servicePort))
}