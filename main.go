package main

import (
	"net/http"
	"os"
	"strings"
	"test-case-backend/auth"
	"test-case-backend/connection"
	"test-case-backend/handler"
	"test-case-backend/helper"
	"test-case-backend/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	err := godotenv.Load("backend.env")
	helper.FailOnError(err, "Error when trying to env")
	db, err = connection.InitDBConnection()
	if err != nil || db == nil {
		helper.FailOnError(err, "Error to connect DB")
		os.Exit(1)
	}
}
func main() {

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewHandler(userService, authService)
	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("api/v0")
	api.POST("/users", userHandler.Register)
	api.GET("/users/:username", authMiddleware(userService, authService), userHandler.Profile)
	api.POST("/login", userHandler.Login)
	api.POST("/foto", authMiddleware(userService, authService), userHandler.UploadFoto)
	api.PUT("/user", authMiddleware(userService, authService), userHandler.Update)
	api.DELETE("/user", authMiddleware(userService, authService), userHandler.DeleteAccount)
	router.Run()
}
func authMiddleware(userService user.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// Check header has "Bearer"
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// Check header length = 2
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) != 2 {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := arrayToken[1]
		// validate token
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId := int(payload["user_id"].(float64))
		currentUser, err := userService.FindUserByID(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// Set context
		c.Set("currentUser", currentUser)
	}
}
