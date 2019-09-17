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

func RouterMark() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/mark/{studentID}", controllers.GetMarks).Methods("GET")
	router.HandleFunc("/marks", controllers.AddMarks).Methods("POST")
	router.HandleFunc("/mark/{marksID}", controllers.UpdateMarks).Methods("PUT")

	return router
}

func TestValidMarksEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/mark/59", nil)
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidMarksEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/mark/1234", nil)
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")

}

func TestValidPostMarkEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	marks := &models.Marks{
		ExamTypeID: 2,
		Subject:    "Arts",
		StudentID:  5,
		Marks:      25,
	}
	jsonMarks, _ := json.Marshal(marks)

	request, _ := http.NewRequest("POST", "/marks", bytes.NewBuffer(jsonMarks))
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidStudentIDPostMarkEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	marks := &models.Marks{
		ExamTypeID: 2,
		Subject:    "Arts",
		StudentID:  512345,
		Marks:      27,
	}
	jsonMarks, _ := json.Marshal(marks)

	request, _ := http.NewRequest("POST", "/marks", bytes.NewBuffer(jsonMarks))
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestInvalidSubjectPostMarkEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	marks := &models.Marks{
		ExamTypeID: 2,
		Subject:    "Nice Subject",
		StudentID:  5,
		Marks:      27,
	}
	jsonMarks, _ := json.Marshal(marks)

	request, _ := http.NewRequest("POST", "/marks", bytes.NewBuffer(jsonMarks))
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestInvalidExamtypeIDPostMarkEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	marks := &models.Marks{
		ExamTypeID: 20,
		Subject:    "Arts",
		StudentID:  5,
		Marks:      27,
	}
	jsonMarks, _ := json.Marshal(marks)

	request, _ := http.NewRequest("POST", "/marks", bytes.NewBuffer(jsonMarks))
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestValidUpdateMarkEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	marks := &models.Marks{
		ExamTypeID: 2,
		Subject:    "Arts",
		StudentID:  5,
		Marks:      27,
	}
	jsonMarks, _ := json.Marshal(marks)

	request, _ := http.NewRequest("PUT", "/mark/6", bytes.NewBuffer(jsonMarks))
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidStudentIDUpdateMarkEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	marks := &models.Marks{
		ExamTypeID: 2,
		Subject:    "Arts",
		StudentID:  512345,
		Marks:      27,
	}
	jsonMarks, _ := json.Marshal(marks)

	request, _ := http.NewRequest("PUT", "/mark/2", bytes.NewBuffer(jsonMarks))
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestInvalidSubjectUpdateMarkEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	marks := &models.Marks{
		ExamTypeID: 2,
		Subject:    "Nice Subject",
		StudentID:  5,
		Marks:      27,
	}
	jsonMarks, _ := json.Marshal(marks)

	request, _ := http.NewRequest("PUT", "/mark/3", bytes.NewBuffer(jsonMarks))
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestInvalidExamtypeIDUpdateMarkEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	marks := &models.Marks{
		ExamTypeID: 20,
		Subject:    "Arts",
		StudentID:  5,
		Marks:      27,
	}
	jsonMarks, _ := json.Marshal(marks)

	request, _ := http.NewRequest("PUT", "/mark/4", bytes.NewBuffer(jsonMarks))
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}
