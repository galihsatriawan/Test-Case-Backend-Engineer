package handler

import (
	"net/http"
	"test-case-backend/auth"
	"test-case-backend/helper"
	"test-case-backend/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
}
func (h *userHandler) Update(c *gin.Context) {
	var input user.UpdateInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}
		response := helper.APIResponse("Register user failed", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	updatedUser, err := h.userService.Update(currentUser, input)

	if err != nil {
		response := helper.APIResponse("Update was failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(updatedUser, "")
	response := helper.APIResponse("User is successfully Updated", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
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
	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		response := helper.APIResponse("Register user success but failed to create token", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("User is successfully registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
func (h *userHandler) Login(c *gin.Context) {
	var inputLogin user.LoginInput
	err := c.ShouldBindJSON(&inputLogin)
	// Check error validation
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// Check user

	loggedInUser, err := h.userService.Login(inputLogin)
	if err != nil {
		errorMessages := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnauthorized, "error", errorMessages)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusUnauthorized, "error", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	formatter := user.FormatUser(loggedInUser, token)
	response := helper.APIResponse("Login successfully", http.StatusAccepted, "success", formatter)

	c.JSON(http.StatusOK, response)
}
