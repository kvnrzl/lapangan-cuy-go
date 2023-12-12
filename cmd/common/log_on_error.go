package common

import "log"

func LogOnError(err error, msg string) {
	if err != nil {
		log.Printf(`%s: %s`, msg, err.Error())
	}
}
