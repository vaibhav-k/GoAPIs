package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

// GetMarks gets the marks of a student from the database
func GetMarks(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting marks!")

	params := mux.Vars(r)

	// Call the handler
	marks, err := models.GetMarks(params["id"])

	if err != "" {
		marksdetails := models.Response{
			StatusCode: utils.WrongParam,
			Message:    utils.GetFailed,
			Data:       marks,
		}

		// Return the exam detail's
		NoContent(w, marksdetails)
		return
	}

	marksdetails := models.Response{
		StatusCode: utils.WrongParam,
		Message:    utils.GetFailed,
		Data:       marks,
	}

	// Return from the function
	ResponseJSON(w, marksdetails)
}

// AddMarks adds marks to the database
func AddMarks(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "POST" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the POST method for this route")
	}
	fmt.Println("Adding marks!")
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var mark models.Marks
	err := decoder.Decode(&mark)
	if err != nil {
		panic(err)
	}

	// User input validation and calling the handler
	if mark.Marks < 0 {
		ResponseJSON(w, "Please enter positive marks")
	} else {
		models.AddMarks(w, r, mark)
		ResponseJSON(w, "Marks added!")
	}
}

// UpdateMarks updates marks of an exam
func UpdateMarks(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the PUT method for this route")
	}
	fmt.Println("Updating marks!")
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var mark models.Marks
	err := decoder.Decode(&mark)
	if err != nil {
		panic(err)
	}

	// User input validation and calling the handler
	if mark.Marks < 0 {
		ResponseJSON(w, "Please enter positive marks")
	} else {
		params := mux.Vars(r)
		models.UpdateMarks(w, r, params["id"], mark)
		ResponseJSON(w, "Marks updated!")
	}
}
