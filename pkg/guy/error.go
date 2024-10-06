package guy

import (
	"net/http"

	"gorm.io/gorm"
)

type AppError struct {
	StatusCode  int
	Message     string
	ErrorDetail string
}

func (e AppError) Error() string {
	return e.Message
}

func NewAppError(statusCode int, message string, errorDetail string) AppError {
	return AppError{
		Message:     message,
		StatusCode:  statusCode,
		ErrorDetail: errorDetail,
	}
}

func NewRepoError(err error, message string) AppError {
	statusCode := http.StatusInternalServerError
	if err == gorm.ErrRecordNotFound {
		statusCode = http.StatusNotFound
	}

	return AppError{
		Message:     message,
		StatusCode:  statusCode,
		ErrorDetail: err.Error(),
	}
}
