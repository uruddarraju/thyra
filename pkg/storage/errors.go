package storage

import "fmt"

type NotFoundError struct {
	message string
}

type DuplicateEntryError struct {
	message string
}

func NewNotFoundError(message string) error {
	return NotFoundError{message: fmt.Sprintf("Object not found: %s", message)}
}

func NewDuplicateEntryError(message string) error {
	return DuplicateEntryError{message: fmt.Sprintf("Duplicate entry for the object: %s", message)}
}

func (e NotFoundError) Error() string {
	return e.message
}

func (e DuplicateEntryError) Error() string {
	return e.message
}

func IsNotFoundError(err error) bool {
	if _, ok := err.(NotFoundError); ok {
		return true
	}
	return false
}
