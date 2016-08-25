package errors

import (
	"log"
	"os"

	"gopkg.in/inconshreveable/log15.v2"
)

// CheckFatal checks if the error is not nil, if it's not then executes `log.Fatal`.
func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CheckPanic checks if the error is not nil, if it's not, then panics.
func CheckPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckLog15Crit checks if the error is not nil, if it's not then executes prints the
// error using `log15.Crit` and exits with status code 1.
func CheckLog15Crit(err error) {
	if err != nil {
		log15.Crit(err.Error())
		os.Exit(1)
	}
}
