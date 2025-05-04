package httputils

import (
	"books/internal/core/domain"
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	} else {
		w.WriteHeader(http.StatusNoContent) // No content
	}
}

func HandleError(w http.ResponseWriter, errResponse domain.DomainError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errResponse.Code.ToHttpStatus())

	var resp = map[string]string{
		"message": errResponse.Message,
		"code":    errResponse.Code.String(),
	}
	if errResponse.Error != nil {
		resp["error"] = errResponse.Error.Error()
	}

	json.NewEncoder(w).Encode(resp)

}

func BindJson(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unknown fields
	return decoder.Decode(v)
}
