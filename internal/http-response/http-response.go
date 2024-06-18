package httpresponse

import (
	"encoding/json"
	"net/http"
)

func HandleResponseNoContent(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func HandleResponse(w http.ResponseWriter, statusCode int, body interface{}) {

	respBytes, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respBytes)
}

func HandleError(w http.ResponseWriter, statusCode int, message string, err error) {

	jsonResponse := map[string]string{
		"message": message,
	}
	if err != nil {
		jsonResponse["error"] = err.Error()
	}

	respBytes, _ := json.Marshal(jsonResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respBytes)
}
