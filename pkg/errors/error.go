package errors

type ServiceError struct {
	Code    int
	Message string
}

func (s ServiceError) Error() string {
	return s.Message
}
