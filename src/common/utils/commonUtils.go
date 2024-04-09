package utils

import "errors"

func MillisecondsToSeconds(milliseconds int) float32 {
	return float32(milliseconds) / 1000.0
}

func NewError(message string) error {
	return errors.New(message)
}
