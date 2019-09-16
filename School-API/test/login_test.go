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

func TestValidLoginUserEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.User2{
		Email:    "richestman@money.com",
		Password: "teacher",
	}
	jsonUser, _ := json.Marshal(teacher)

	request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
	response := httptest.NewRecorder()
	RouterLogin().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidLoginUserEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.User2{
		Email:    "richestman@money.com2",
		Password: "teacher",
	}
	jsonUser, _ := json.Marshal(teacher)

	request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
	response := httptest.NewRecorder()
	RouterLogin().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}
