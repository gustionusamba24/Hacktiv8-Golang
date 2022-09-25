package handler 

import (
	"assignment_2/customers"
	"assignment_2/helper"
	"net/http"
	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	service customers.Service
}

func NewCustomerHandler(service customers.Service) *customerHandler {
	return &customerHandler{service}
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	var input customers.CustomerInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed creating customer", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	newCustomer, err := h.service.CreateCustomer(input)
	if err != nil {
		response := helper.APIResponse("Failed to create customer", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create customer", http.StatusOK, "success", customers.FormatCustomer(newCustomer))
	c.JSON(http.StatusOK, response)
}