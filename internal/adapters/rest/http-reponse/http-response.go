package httpreponse

import (
	"encoding/json"
	"net/http"
)

// @todo finish
func Success(w *http.ResponseWriter) {

}

type InternalError struct {
	Message string
	Error   error
}

// @todo finish
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

// @todo finish
func NoContent() {

}
