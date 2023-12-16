package common

import (
	"errors"
	"log"
)

func HandleError(err error, msg string) error {
	newError := errors.New(msg)
	log.Printf(`%s: %s`, msg, err.Error())

	return newError

}
