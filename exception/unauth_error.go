package exception

type UnauthError struct {
	Error string
}

func NewUnauthError(error string) UnauthError {
	return UnauthError{
		Error: error,
	}
}
