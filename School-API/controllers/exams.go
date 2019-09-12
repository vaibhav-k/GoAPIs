package controllers

import (
	"fmt"
	"net/http"

	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

// GetExam gets the exam's detail from the database
func GetExam(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		fmt.Println("Please use the GET method for this route")
	}
	fmt.Println("Getting exam's details!")

	params := mux.Vars(r)

	// Call the handler
	exam, err := models.GetExam(params["id"])

	if err != nil {
		http.Error(w, http.StatusText(utils.ErrorCode), utils.ErrorCode)
		DidNotComplete(w, http.StatusText(utils.ErrorCode))
	}

	examdetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotExams,
		Data:       exam,
	}

	// Return from the function
	ResponseJSON(w, examdetails)
}

// GetExams gets all the exams to be held from the database
func GetExams(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		fmt.Println("Please use the GET method for this route")
	}
	fmt.Println("Getting all exams details!")

	// Call the handler
	exams, err := models.GetExams()

	if err != nil {
		http.Error(w, http.StatusText(utils.ErrorCode), utils.ErrorCode)
		DidNotComplete(w, http.StatusText(utils.ErrorCode))
	}

	examsdetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotExams,
		Data:       exams,
	}

	// Return from the function
	ResponseJSON(w, examsdetails)
}
