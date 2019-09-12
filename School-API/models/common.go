package models

// import "database/sql"

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// DB is a struct for the database connection
type DB struct {
	DBCon *sql.DB
}

// Response in JSON
type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// ResponseJSON does the common stuff
func ResponseJSON(w http.ResponseWriter, ResponseData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseData)
}
