package exception

type UnauthorizedError struct {
	Error string
}

func NewUnauthorizedError(err string) UnauthorizedError {
	return UnauthorizedError{Error: err}
}
