package helper

import (
	"fmt"
	"log"
)
import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	response := Response{
		Meta: meta,
		Data: data,
	}
	return response
}
func FormatValidationError(err error) []string {

	// make error information more beautiful
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		log.Panic(err)
	}
}

func Constanta(name string) string {
	return fmt.Sprint(name)
}
