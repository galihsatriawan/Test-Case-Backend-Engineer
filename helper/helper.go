package helper

import (
	"fmt"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		log.Panic(err)
	}
}

func Constanta(name string) string {
	return fmt.Sprint(name)
}
