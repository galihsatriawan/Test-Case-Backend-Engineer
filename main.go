package main

import (
	"os"
	"test-case-backend/connection"
	"test-case-backend/helper"

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

}
