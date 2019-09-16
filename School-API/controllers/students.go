package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

// GetStudents gets all of all students from the database
func GetStudents(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting all students details!")

	// Call the handler
	students, err := models.GetStudents()

	if err != "" {
		studentsdetails := models.Response{
			StatusCode: utils.WrongParam,
			Message:    utils.GetFailed,
			Data:       students,
		}
		ResponseJSON(w, studentsdetails)
		return
	}

	studentsdetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotStudent,
		Data:       students,
	}

	// Return from the function
	ResponseJSON(w, studentsdetails)
}

// GetStudent gets all the details of a student from the database
func GetStudent(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting student's details!")

	params := mux.Vars(r)

	// Call the handler
	student, err := models.GetStudent(params["id"])

	studentdetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotStudent,
		Data:       student,
	}

	if err != "" {
		studentdetails = models.Response{
			StatusCode: utils.WrongParam,
			Message:    utils.GetFailed,
			Data:       student,
		}
		ResponseJSON(w, studentdetails)
		return
	}

	// Return from the function
	ResponseJSON(w, studentdetails)
}

// DeleteStudent deletes a student from the database
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the DELETE method for this route")
	}
	fmt.Println("Deleting student's details!")
	w.Header().Set("Content-Type", "application/json")

	// Call the handler
	params := mux.Vars(r)
	err := models.DeleteStudent(w, r, params["id"])

	studentdetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    "Deleted the student!",
		Data:       "Deleted the student!",
	}

	if err != "" {
		studentdetails = models.Response{
			StatusCode: utils.WrongParam,
			Message:    utils.DeletionFailed,
			Data:       utils.DeletionFailed,
		}
		ResponseJSON(w, studentdetails)
		return
	}

	// Return from the function
	ResponseJSON(w, studentdetails)
}

// UpdateStudent updates details of a student
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the PUT method for this route")
	}
	fmt.Println("Updating student's details!")
	w.Header().Set("Content-Type", "application/json")

	// Get the user's input details from the POST body
	decoder := json.NewDecoder(r.Body)
	var student models.Students
	err := decoder.Decode(&student)
	if err != nil {
		panic(err)
	}

	// Check the user's input and then call the handler
	if student.StudentID < 0 {
		ResponseJSON(w, "Please give a valid ID")
	} else if student.FirstName == "" {
		ResponseJSON(w, "Please give a first name")
	} else if student.Password == "" {
		ResponseJSON(w, "Please give a password")
	} else {
		params := mux.Vars(r)
		err := models.UpdateStudent(w, r, params["id"], student)

		studentdetails := models.Response{
			StatusCode: utils.SuccessCode,
			Message:    "Updated the student!",
			Data:       "Updated the student!",
		}

		if err != "" {
			studentdetails = models.Response{
				StatusCode: utils.WrongParam,
				Message:    utils.UpdatingFailed,
				Data:       utils.UpdatingFailed,
			}
			ResponseJSON(w, studentdetails)
			return
		}

		// Return from the function
		ResponseJSON(w, studentdetails)
	}
}
