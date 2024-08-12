package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kenzidjulfri/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route := router.Group("api")
	{
		h := handler.NewBaseHandler()
		route.POST("/order", h.Order)
		route.GET("/bill/:id", h.Bill)
	}
}
