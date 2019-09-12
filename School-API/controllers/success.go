package controllers

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON does the common stuff
func ResponseJSON(w http.ResponseWriter, ResponseData interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseData)
}

// NoContent sends an error message as a JSON
func NoContent(w http.ResponseWriter, ResponseData interface{}) {
	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseData)
}
