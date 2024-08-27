package httpreponse

import (
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, errResponse domain.DomainError) {

	statusCode := mapErrorToHTTPStatusCode(errResponse.Code)
	var resp = map[string]string{
		"message": errResponse.Message,
		"code":    errResponse.Code,
	}
	if errResponse.Error != nil {
		resp["error"] = errResponse.Error.Error()
	}
	ctx.JSON(statusCode, resp)
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
