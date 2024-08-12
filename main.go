package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kenzidjulfri/config"
	"github.com/kenzidjulfri/docs"
	"github.com/kenzidjulfri/routes"
)

// @title           Simple F&B API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api
func main() {
	config.Init()

	router := gin.New()
	docs.SwaggerInfo.BasePath = "/api"
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.InitRoutes(router)
	router.Run(":" + os.Getenv("APP_PORT"))
}
