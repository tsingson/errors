// +build go1.13

package errors

import (
	"fmt"

	"errors"
)

// New  wrap errors.New
func New(text string) error {
	return errors.New(text)
}

// Errorf  wrap fmt.Errorf
func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// As wrap errors.As
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Is wrap errors.As
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// Unwrap wrap xerrors.Unwrap
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
