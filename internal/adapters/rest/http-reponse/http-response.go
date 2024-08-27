package httpreponse

import (
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	"encoding/json"
	"net/http"
)

// @todo finish
func Success(w *http.ResponseWriter) {

}

// Deprecated:  Migrar para ErrorResponseModel
type InternalError struct {
	Message string
	Error   error
}
type ErrorResponseModel struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
	Code    string `json:"code"`
}

func BadRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	err, _ := json.Marshal(map[string]string{
		"message": message,
	})
	w.Write(err)
}

func InternalServerError(w http.ResponseWriter, internalErr InternalError) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	var resp = map[string]string{
		"message": internalErr.Message,
	}
	if internalErr.Error != nil {
		resp["error"] = internalErr.Error.Error()
	}
	err, _ := json.Marshal(resp)
	w.Write(err)
}

func ErrorResponse(w http.ResponseWriter, errResponse domain.DomainError) {

	statusCode := mapErrorToHTTPStatusCode(errResponse.Code)
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	var resp = map[string]string{
		"message": errResponse.Message,
		"code":    errResponse.Code,
	}
	if errResponse.Error != nil {
		resp["error"] = errResponse.Error.Error()
	}
	err, _ := json.Marshal(resp)
	w.Write(err)
}

// @todo finish
func NoContent() {

}

// Define the mapping function from error codes to HTTP status codes
func mapErrorToHTTPStatusCode(errCode string) int {

	switch errCode {
	case errorscode.ErrNotFound:
		return http.StatusNotFound // 404 Not Found
	case errorscode.ErrInvalidInput:
		return http.StatusBadRequest // 400 Bad Request
	case errorscode.ErrUnauthorized:
		return http.StatusUnauthorized // 401 Unauthorized
	case errorscode.ErrInternalError:
		return http.StatusInternalServerError // 500 Internal Server Error
	default:
		return http.StatusInternalServerError // Default to 500 for unknown errors
	}

}
