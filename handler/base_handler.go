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

// @BasePath /api/v1

// Order godoc
// @Summary Create a new order
// @Description create a new order by providing the table id and the details of the order, if success returns the newly created order id and a list of printers
// @Param order body request.OrderRequest true "order request"
// @Accept json
// @Produce json
// @Success 200 {object} response.OrderResponse
// @Router /order [post]
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

// Bill godoc
// @Summary Retrieves an order by its order id
// @Description get order by its id, returning the order details and the printers
// @Param id path int true "Order ID"
// @Produce json
// @Success 200 {object} response.BillResponse
// @Router /bill/{id} [get]
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

// Menu godoc
// @Summary Retrieves menu
// @Description get menu and the item's details
// @Produce json
// @Success 200 {object} response.MenuResponse
// @Router /menu [get]
func (h *BaseHandler) Menu(c *gin.Context) {
	resp, err := h.baseService.GetMenu()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data found", "data": &resp})
}
