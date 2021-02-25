package handler

import (
	"net/http"
	"test-case-backend/helper"
	"test-case-backend/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) Register(c *gin.Context) {
	var input user.RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}
		response := helper.APIResponse("Register user failed", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser)
	response := helper.APIResponse("User is successfully registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
