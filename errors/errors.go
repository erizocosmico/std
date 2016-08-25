package errors

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/inconshreveable/log15.v2"
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

func CheckLog15Crit(err error) {
	if err != nil {
		log15.Crit(err.Error())
		os.Exit(1)
	}
}

func Formatted(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}
