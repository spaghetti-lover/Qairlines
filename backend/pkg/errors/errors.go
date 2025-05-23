package errors

type AppError struct {
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}
