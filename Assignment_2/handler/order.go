package handler 

import (
	"assignment_2/orders"
	"assignment_2/helper"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	service orders.Service
}

func NewOrderHandler(service orders.Service) *orderHandler {
	return &orderHandler{service}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	var input orders.OrderInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed creating order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newOrder, err := h.service.Create(input)
	if err != nil {
		response := helper.APIResponse("Failed to create order, customer id not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create order", http.StatusOK, "success", orders.OrderFormat(newOrder))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) GetOrders(c *gin.Context) {
	order, err := h.service.GetOrders()
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed getting all orders", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Orders", http.StatusOK, "success", orders.OrdersFormat(order))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) UpdateOrder(c *gin.Context) {
	var inputID orders.GetOrderByID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed updating order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData orders.OrderUpdate

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed updating order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedOrder, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed updating order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success updating order", http.StatusOK, "success", orders.OrderFormat(updatedOrder))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) GetOrderByID(c *gin.Context) {
	var inputID orders.GetOrderByID

	err := c.ShouldBindUri("id")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed getting order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	order, err := h.service.GetOrderByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed getting order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success getting order", http.StatusOK, "success", orders.OrderFormat(order))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) DeleteOrder(c *gin.Context) {

	idparam := c.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed deleting order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	order, err := h.service.Delete(id)
	if err != nil {
		response := helper.APIResponse("Failed deleting order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success deleting Order", http.StatusOK, "success", orders.OrderFormat(order))
	c.JSON(http.StatusOK, response)
}




