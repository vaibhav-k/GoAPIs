package test

import (
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

	return router
}

func TestSectionsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/sections", nil)
	response := httptest.NewRecorder()
	RouterSection().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}
