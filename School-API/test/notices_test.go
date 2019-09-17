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

func RouterNotice() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/notices/{noticeID}", controllers.GetNotice).Methods("GET")
	router.HandleFunc("/notices", controllers.GetNotices).Methods("GET")
	router.HandleFunc("/notices", controllers.AddNotice).Methods("POST")

	return router
}

func TestValidNoticesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/notices/1", nil)
	response := httptest.NewRecorder()
	RouterNotice().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidNoticesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/notices/123456", nil)
	response := httptest.NewRecorder()
	RouterNotice().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No response is expected")
}

func TestNoticesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/notices", nil)
	response := httptest.NewRecorder()
	RouterNotice().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestValidTeacherIDPostNoticeEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	notice := &models.Notice{
		TeacherID: 2,
		Notice:    "Hello world",
	}
	jsonNotice, _ := json.Marshal(notice)

	request, _ := http.NewRequest("POST", "/notices", bytes.NewBuffer(jsonNotice))
	response := httptest.NewRecorder()
	RouterNotice().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidTeacherIDPostNoticeEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	notice := &models.Notice{
		TeacherID: 200,
		Notice:    "Hello world",
	}
	jsonNotice, _ := json.Marshal(notice)

	request, _ := http.NewRequest("POST", "/notices", bytes.NewBuffer(jsonNotice))
	response := httptest.NewRecorder()
	RouterNotice().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestInvalidNoticePostNoticeEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	notice := &models.Notice{
		TeacherID: 2,
		Notice:    "Hi",
	}
	jsonNotice, _ := json.Marshal(notice)

	request, _ := http.NewRequest("POST", "/notices", bytes.NewBuffer(jsonNotice))
	response := httptest.NewRecorder()
	RouterNotice().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}
