package internal

import "errors"

var (
	errInvalidField     = errors.New("field is invalid or not supported")
	errInvalidOperation = errors.New("operation is invalid or not supported")
	errNumberValue      = errors.New("the value should be a number")
)
