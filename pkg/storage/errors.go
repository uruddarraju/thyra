package storage

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) error {
	return NotFoundError{message: message}
}

func (e *NotFoundError) Error() string {
	return e.message
}
