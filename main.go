package main

import (
	"os"
	"test-case-backend/connection"
	"test-case-backend/handler"
	"test-case-backend/helper"
	"test-case-backend/user"

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

	userHandler := handler.NewHandler(userService)
	router := gin.Default()

	api := router.Group("api/v0")
	api.POST("/users", userHandler.Register)

	router.Run()
}
