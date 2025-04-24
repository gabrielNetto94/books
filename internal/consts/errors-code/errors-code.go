package errorscode

import "net/http"

type ErrorCode string

const (
	ErrNotFound      ErrorCode = "NOT_FOUND"
	ErrInvalidInput  ErrorCode = "INVALID_INPUT"
	ErrUnauthorized  ErrorCode = "UNAUTHORIZED"
	ErrInternalError ErrorCode = "INTERNAL_ERROR"
)

func (err ErrorCode) String() string {
	return (string)(err)
}

var errorCodeToHttpStatus = map[ErrorCode]int{
	ErrNotFound:      http.StatusNotFound,            // 404 Not Found
	ErrInvalidInput:  http.StatusBadRequest,          // 400 Bad Request
	ErrUnauthorized:  http.StatusUnauthorized,        // 401 Unauthorized
	ErrInternalError: http.StatusInternalServerError, // 500 Internal Server Error
}

func (err ErrorCode) ToHttpStatus() int {
	if status, exists := errorCodeToHttpStatus[err]; exists {
		return status
	}
	return http.StatusInternalServerError // Default to 500 for unknown errors
}
