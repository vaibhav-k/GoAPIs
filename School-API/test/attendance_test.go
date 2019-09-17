package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"../controllers"
	"../models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func RouterAttd() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/attendance", controllers.GetAttendance).Methods("GET")
	router.HandleFunc("/attendances", controllers.GetAttendances).Methods("GET")

	return router
}

// Tests attendances for a date whose attendances are taken
func TestValidAttendancesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	// MAke the request
	request, _ := http.NewRequest("GET", "/attendance?date=02-Jan-20", nil)
	response := httptest.NewRecorder()
	RouterAttd().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests attendances for a date whose attendances are not taken
func TestInvalidAttendancesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	// Make the request
	request, _ := http.NewRequest("GET", "/attendance?date=200-Jan-20", nil)
	response := httptest.NewRecorder()
	RouterAttd().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests all attendances
func TestAttendancesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	// Make the request
	request, _ := http.NewRequest("GET", "/attendances", nil)
	response := httptest.NewRecorder()
	RouterAttd().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}
