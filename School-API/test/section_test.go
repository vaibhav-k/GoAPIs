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

func RouterSection() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/sections", controllers.GetSections).Methods("GET")
	router.HandleFunc("/sections", controllers.AddSection).Methods("POST")

	return router
}

// Tests if all the sections are viewable
func TestSectionsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	// Make the request
	request, _ := http.NewRequest("GET", "/sections", nil)
	response := httptest.NewRecorder()
	RouterSection().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if all a section can be added with valid details
func TestPostSectionEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	section := &models.Sections{
		ClassID:   9,
		SectionID: 4,
	}
	jsonSection, _ := json.Marshal(section)

	// Make the request
	request, _ := http.NewRequest("POST", "/sections", bytes.NewBuffer(jsonSection))
	response := httptest.NewRecorder()
	RouterSection().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}
