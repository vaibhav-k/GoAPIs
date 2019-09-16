package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

// GetTeacher gets the teacher's detail from the database
func GetTeacher(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting teacher's details!")

	// params := mux.Vars(r)
	URISegments := strings.Split(r.URL.Path, "/")

	if URISegments[2] != "" {
		// Call the handler
		teacher, err := models.GetTeacher(URISegments[2])

		teacherdetails := models.Response{
			StatusCode: utils.SuccessCode,
			Message:    utils.GotTeacher,
			Data:       teacher,
		}

		if err != "" {
			teacherdetails := models.Response{
				StatusCode: utils.WrongParam,
				Message:    utils.GetFailed,
				Data:       teacher,
			}
			ResponseJSON(w, teacherdetails)
			return
		}

		// Return from the function
		ResponseJSON(w, teacherdetails)
	} else {
		// Return from the function
		ResponseJSON(w, "You have to supply a teacher ID at the end of your URL")
	}
}

// GetTeachers gets all teachers detail from the database
func GetTeachers(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting all teachers details!")

	// Call the handler
	teachers, err := models.GetTeachers()

	teachersdetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotTeachers,
		Data:       teachers,
	}

	if err != "" {
		teachersdetails := models.Response{
			StatusCode: utils.WrongParam,
			Message:    utils.GetFailed,
			Data:       teachers,
		}
		ResponseJSON(w, teachersdetails)
		return
	}

	// Return from the function
	ResponseJSON(w, teachersdetails)
}

// AddTeacher adds marks to the database
func AddTeacher(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "POST" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the POST method for this route")
	}
	fmt.Println("Adding a teacher!")
	w.Header().Set("Content-Type", "application/json")

	// Get the user's input details from the POST body
	decoder := json.NewDecoder(r.Body)
	var teacher models.Teachers
	err := decoder.Decode(&teacher)
	if err != nil {
		panic(err)
	}

	// Check the user's input and then call the handler
	if teacher.TeacherID < 0 {
		ResponseJSON(w, "Please give a valid ID")
	} else if teacher.FirstName == "" {
		ResponseJSON(w, "Please give a first name")
	} else if teacher.Password == "" {
		ResponseJSON(w, "Please give a password")
	} else {
		err := models.AddTeacher(w, r, teacher)

		if err != "" {
			teacherdetails := models.Response{
				StatusCode: utils.WrongParam,
				Message:    utils.InsertionFailed,
				Data:       err,
			}
			ResponseJSON(w, teacherdetails)
			return
		}

		teacherdetails := models.Response{
			StatusCode: utils.SuccessCode,
			Message:    "Adding successful",
			Data:       "Adding successful",
		}
		ResponseJSON(w, teacherdetails)
		return
	}
}

// DeleteTeacher deletes a teacher from the database
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the DELETE method for this route")
	}
	fmt.Println("Deleting a teacher!")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Call the handler
	er := models.DeleteTeacher(w, r, params["id"])

	if er != "" {
		teacherdetails := models.Response{
			StatusCode: utils.WrongParam,
			Message:    utils.DeletionFailed,
			Data:       er,
		}
		ResponseJSON(w, teacherdetails)
		return
	}

	teacherdetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    "Deleting successful",
		Data:       "Deleting successful",
	}
	ResponseJSON(w, teacherdetails)
	return
}

// UpdateTeacher updates details of a teacher
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the PUT method for this route")
	}
	fmt.Println("Updating a teacher's details!")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var teacher models.Teachers
	err := decoder.Decode(&teacher)
	if err != nil {
		panic(err)
	}

	if teacher.TeacherID < 0 {
		ResponseJSON(w, "Please give a valid ID")
	} else if teacher.FirstName == "" {
		ResponseJSON(w, "Please give a first name")
	} else if teacher.Password == "" {
		ResponseJSON(w, "Please give a password")
	} else {
		er := models.UpdateTeacher(w, r, params["id"], teacher)

		if er != "" {
			teacherdetails := models.Response{
				StatusCode: utils.WrongParam,
				Message:    utils.UpdatingFailed,
				Data:       utils.UpdatingFailed,
			}
			ResponseJSON(w, teacherdetails)
			return
		}

		teacherdetails := models.Response{
			StatusCode: utils.SuccessCode,
			Message:    "Updating successful",
			Data:       "Updating successful",
		}
		ResponseJSON(w, teacherdetails)
		return
	}
}
