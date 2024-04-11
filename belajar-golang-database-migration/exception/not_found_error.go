package exception

type NotFoundError struct {
	error string
}

func (e *NotFoundError) Error() string {
	return e.error
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{error}
}
