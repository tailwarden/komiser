package internal

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	errInvalidField     = errors.New("field is invalid or not supported")
	errInvalidOperation = errors.New("operation is invalid or not supported")
	errNumberValue      = errors.New("the value should be a number")
)

func handleError(err error, msg string) {
	if err != nil {
		log.WithError(err).Error(msg)
	}
}
