package main

import (
	"fmt"
	"os"
	"test-case-backend/connection"
	"test-case-backend/helper"
	"test-case-backend/user"

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

	userInput := user.RegisterInput{}
	userInput.Username = "galihsatriawan"
	userInput.Password = "OK"
	userInput.NamaLengkap = "galihsatriawan"
	newUser, err := userService.RegisterUser(userInput)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Success create user")
	fmt.Println(newUser)
}
