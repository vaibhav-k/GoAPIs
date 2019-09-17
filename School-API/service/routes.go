package service

import (
	"fmt"
	"net/http"

	"../controllers"
	"../middlewares"

	// Package for the db connection
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// InitServer gives the routes for the API and then starts the server
func InitServer() {
	// Create a new router and print the URL
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	// Make the routes with the handler functions
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/notices", middlewares.IsLoggedIn(controllers.GetNotices)).Methods("GET")
	router.HandleFunc("/notice/{id}", middlewares.IsLoggedIn(controllers.GetNotice)).Methods("GET")
	router.HandleFunc("/mark/{id}", middlewares.IsLoggedIn(controllers.GetMarks)).Methods("GET")
	router.HandleFunc("/attendance", middlewares.IsLoggedIn(controllers.GetAttendance)).Methods("GET")
	router.HandleFunc("/attendances", middlewares.IsLoggedIn(controllers.GetAttendances)).Methods("GET")
	router.HandleFunc("/subjects", middlewares.IsLoggedIn(controllers.GetSubjects)).Methods("GET")
	router.HandleFunc("/teacher/{id}", middlewares.IsLoggedIn(controllers.GetTeacher)).Methods("GET")
	router.HandleFunc("/teachers", middlewares.IsLoggedIn(controllers.GetTeachers)).Methods("GET")
	router.HandleFunc("/exam/{id}", middlewares.IsStudent(controllers.GetExam)).Methods("GET")
	router.HandleFunc("/exams", middlewares.IsStudent(controllers.GetExams)).Methods("GET")
	router.HandleFunc("/sections", middlewares.IsLoggedIn(controllers.GetSections)).Methods("GET")
	router.HandleFunc("/students", middlewares.IsLoggedIn(controllers.GetStudents)).Methods("GET")
	router.HandleFunc("/student/{id}", middlewares.IsLoggedIn(controllers.GetStudent)).Methods("GET")
	router.HandleFunc("/notices", middlewares.IsTeacher(controllers.AddNotice)).Methods("POST")
	router.HandleFunc("/marks", middlewares.IsTeacher(controllers.AddMarks)).Methods("POST")
	router.HandleFunc("/subjects", middlewares.IsAdmin(controllers.AddSubject)).Methods("POST")
	router.HandleFunc("/sections", middlewares.IsAdmin(controllers.AddSection)).Methods("POST")
	router.HandleFunc("/teachers", middlewares.IsAdmin(controllers.AddTeacher)).Methods("POST")
	router.HandleFunc("/student/{id}", middlewares.IsAdmin(controllers.DeleteStudent)).Methods("DELETE")
	router.HandleFunc("/teacher/{id}", middlewares.IsAdmin(controllers.DeleteTeacher)).Methods("DELETE")
	router.HandleFunc("/student/{id}", middlewares.IsAdmin(controllers.UpdateStudent)).Methods("PUT")
	router.HandleFunc("/teacher/{id}", middlewares.IsAdmin(controllers.UpdateTeacher)).Methods("PUT")
	router.HandleFunc("/mark/{id}", middlewares.IsTeacher(controllers.UpdateMarks)).Methods("PUT")

	// Start the server and handle CORS also
	port := ":8080"
	http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT"}), handlers.AllowedOrigins([]string{"*"}))(router))
}
