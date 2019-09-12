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

func RouterSubject() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/subjects", controllers.GetSubjects).Methods("GET")

	return router
}

func TestSubjectsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/subjects", nil)
	response := httptest.NewRecorder()
	RouterSubject().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
