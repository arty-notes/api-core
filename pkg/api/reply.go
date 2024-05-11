package api

import (
	"encoding/json"
	"net/http"
)

func JsonReply(w http.ResponseWriter, statusCode int, body any) {

	jsonBytes, _ := json.Marshal(body)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}
