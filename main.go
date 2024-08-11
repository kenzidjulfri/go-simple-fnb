package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kenzidjulfri/config"
	"github.com/kenzidjulfri/routes"
)

func main() {
	config.Init()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.InitRoutes(router)
	router.Run(":" + os.Getenv("APP_PORT"))
}
