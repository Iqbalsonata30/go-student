package exception

type BadRequestError struct {
	Error string
}

func NewBadRequestError(err string) BadRequestError {
	return BadRequestError{Error: err}
}
