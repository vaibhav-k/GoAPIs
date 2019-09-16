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
	router.HandleFunc("/sections", controllers.AddSection).Methods("POSt")

	return router
}

func TestSectionsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/sections", nil)
	response := httptest.NewRecorder()
	RouterSection().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestValidPostSectionEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	section := &models.Sections{
		ClassID:        9,
		SectionID:      2,
		ClassSectionID: 27,
	}
	jsonSection, _ := json.Marshal(section)

	request, _ := http.NewRequest("POST", "/sections", bytes.NewBuffer(jsonSection))
	response := httptest.NewRecorder()
	RouterSection().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidPostSectionEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	section := &models.Sections{
		ClassID:        8,
		SectionID:      3,
		ClassSectionID: 24,
	}
	jsonSection, _ := json.Marshal(section)

	request, _ := http.NewRequest("POST", "/sections", bytes.NewBuffer(jsonSection))
	response := httptest.NewRecorder()
	RouterSection().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}
