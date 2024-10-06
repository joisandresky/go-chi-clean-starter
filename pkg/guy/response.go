package guy

import (
	"net/http"

	"github.com/go-chi/render"
)

type API struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	Error      string `json:"error,omitempty"`
	ErrorCode  int    `json:"error_code,omitempty"`
}

// HandleError - is used to handle errors coming from usecases
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	if appErr, ok := err.(AppError); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(appErr.StatusCode)
		render.JSON(w, r, API{
			StatusCode: appErr.StatusCode,
			Success:    false,
			Error:      appErr.ErrorDetail,
			Message:    appErr.Message,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	render.JSON(w, r, API{
		StatusCode: http.StatusInternalServerError,
		Success:    false,
		Error:      err.Error(),
		Message:    "Something went wrong, please try again later.",
	})
}

func JSON(w http.ResponseWriter, r *http.Request, data API) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.StatusCode)
	render.JSON(w, r, data)
}

func Error(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	render.JSON(w, r, API{
		StatusCode: status,
		Success:    false,
		Message:    message,
	})
}

func Ok(w http.ResponseWriter, r *http.Request, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, API{
		StatusCode: http.StatusOK,
		Success:    true,
		Data:       data,
		Message:    message,
	})
}

func Created(w http.ResponseWriter, r *http.Request, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, API{
		StatusCode: http.StatusCreated,
		Success:    true,
		Data:       data,
		Message:    message,
	})
}

func NoContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func BadRequest(w http.ResponseWriter, r *http.Request, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	render.JSON(w, r, API{
		StatusCode: http.StatusBadRequest,
		Success:    false,
		Error:      err.Error(),
		Message:    message,
	})
}

func NotFound(w http.ResponseWriter, r *http.Request, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	render.JSON(w, r, API{
		StatusCode: http.StatusNotFound,
		Success:    false,
		Error:      err.Error(),
		Message:    message,
	})
}

func InternalServerError(w http.ResponseWriter, r *http.Request, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	render.JSON(w, r, API{
		StatusCode: http.StatusInternalServerError,
		Success:    false,
		Error:      err.Error(),
		Message:    message,
	})
}

func Unauthorized(w http.ResponseWriter, r *http.Request, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	render.JSON(w, r, API{
		StatusCode: http.StatusUnauthorized,
		Success:    false,
		Error:      err.Error(),
		Message:    message,
	})
}

func Forbidden(w http.ResponseWriter, r *http.Request, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	render.JSON(w, r, API{
		StatusCode: http.StatusForbidden,
		Success:    false,
		Error:      err.Error(),
		Message:    message,
	})
}

func UnprocessableEntity(w http.ResponseWriter, r *http.Request, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	render.JSON(w, r, API{
		StatusCode: http.StatusUnprocessableEntity,
		Success:    false,
		Error:      err.Error(),
		Message:    message,
	})
}
