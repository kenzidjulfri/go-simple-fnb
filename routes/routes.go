package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kenzidjulfri/handler"
)

func InitRoutes(router *gin.Engine) {
	route := router.Group("api")
	{
		h := handler.NewBaseHandler()
		route.POST("/order", h.Order)
		route.GET("/bill/:id", h.Bill)
	}
}
