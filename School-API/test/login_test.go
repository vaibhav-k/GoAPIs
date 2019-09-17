package test

import (
	"bytes"
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

func RouterLogin() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	return router
}

// Tests if a user can login with valid credentials
func TestValidLoginUserEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.User2{
		Email:    "richestman@money.com",
		Password: "teacher",
	}
	jsonUser, _ := json.Marshal(teacher)

	// Make the request
	request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
	response := httptest.NewRecorder()
	RouterLogin().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if a user can login with invalid credentials
func TestInvalidEmailIDLoginUserEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.User2{
		Email:    "richestman@money.com2",
		Password: "teacher",
	}
	jsonUser, _ := json.Marshal(teacher)

	// Make the request
	request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
	response := httptest.NewRecorder()
	RouterLogin().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestInvalidPasswordLoginUserEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.User2{
		Email:    "richestman@money.com",
		Password: "teacher2",
	}
	jsonUser, _ := json.Marshal(teacher)

	// Make the request
	request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
	response := httptest.NewRecorder()
	RouterLogin().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}
