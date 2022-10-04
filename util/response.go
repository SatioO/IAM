package util

import (
	"encoding/json"
	"net/http"
)

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")

	response, _ := json.Marshal(map[string]interface{}{
		"data":       nil,
		"error":      msg,
		"statusCode": code,
	})

	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithJSON ...
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, _ := json.Marshal(map[string]interface{}{
		"data":       payload,
		"error":      nil,
		"statusCode": code,
	})

	w.WriteHeader(code)
	w.Write(response)
}
