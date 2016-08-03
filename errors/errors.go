package errors

import (
	"fmt"
	"log"
)

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func Formatted(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}
