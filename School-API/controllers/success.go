package controllers

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON does the common stuff
func ResponseJSON(w http.ResponseWriter, ResponseData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseData)
}
