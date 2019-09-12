package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../utils"
)

// GetSubjects gets the subjects for each class from the database
func GetSubjects(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting subjects!")

	// Call the handler
	subjects, err := models.GetSubjects()

	if err != "" {
		NoContent(w, subjects)
		return
	}

	subjectsdetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotSubjectIDs,
		Data:       subjects,
	}

	// Return from the function
	ResponseJSON(w, subjectsdetails)
}

// AddSubject adds a new notice to the database
func AddSubject(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "POST" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the POST method for this route")
	}
	fmt.Println("Adding a subject!")

	// Get the user's input details from the POST body
	decoder := json.NewDecoder(r.Body)
	var subject models.SubjectAdd
	err := decoder.Decode(&subject)
	if err != nil {
		panic(err)
	}

	// Check the user's input and then call the handler
	if subject.SubjectID < 0 {
		NoContent(w, "Please give a valid ID")
	} else if subject.Title != "" {
		NoContent(w, "Please give a valid title")
	} else {
		models.AddSubject(w, r, subject)
	}
}
