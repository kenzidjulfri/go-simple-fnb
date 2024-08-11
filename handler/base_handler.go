package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenzidjulfri/request"
	"github.com/kenzidjulfri/services"
	"github.com/kenzidjulfri/utils"
)

type BaseHandler struct {
	baseService services.BaseService
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		baseService: services.GetBaseService(),
	}
}

func (h *BaseHandler) Order(c *gin.Context) {
	var req request.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationError := utils.GenerateErrorValidationResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input", "errors": validationError.Errors})
		return
	}

	resp, err := h.baseService.CreateOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order created successfully", "data": &resp})
}

func (h *BaseHandler) Bill(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	resp, err := h.baseService.GetBill(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data found", "data": &resp})
}
