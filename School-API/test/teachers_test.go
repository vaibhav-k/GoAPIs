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

func RouterTeacher() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/teachers/{id}", controllers.GetTeacher).Methods("GET")
	router.HandleFunc("/teachers", controllers.GetTeachers).Methods("GET")

	return router
}

func TestTeachersEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/teachers", nil)
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
