package httpreponse

import (
	"books/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, errResponse domain.DomainError) {

	var resp = map[string]string{
		"message": errResponse.Message,
		"code":    errResponse.Code.String(),
	}
	if errResponse.Error != nil {
		resp["error"] = errResponse.Error.Error()
	}
	ctx.JSON(errResponse.Code.ToHttpStatus(), resp)
}
