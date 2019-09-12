package controllers

import (
	"fmt"
	"net/http"

	"../models"
	"../utils"
)

// GetAttendance gets the attendance of a date from the database
func GetAttendance(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting attendance!")

	// Get the date parameter frrm the URL
	key := r.URL.Query().Get("date")

	// Call the handler
	attds, err := models.GetAttendance(key)

	if err != nil {
		http.Error(w, http.StatusText(utils.ErrorCode), utils.ErrorCode)
		DidNotComplete(w, http.StatusText(utils.ErrorCode))
	}

	attendances := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotAttendances,
		Data:       attds,
	}
	w.WriteHeader(http.StatusOK)

	// Return the attendances
	ResponseJSON(w, attendances)
}

// GetAttendances gets the attendance of a date from the database
func GetAttendances(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting attendance!")

	// Call the handler
	attds, err := models.GetAttendances()

	if err != nil {
		http.Error(w, http.StatusText(utils.ErrorCode), utils.ErrorCode)
		DidNotComplete(w, http.StatusText(utils.ErrorCode))
	}

	attendancess := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotAttendances,
		Data:       attds,
	}
	w.WriteHeader(http.StatusOK)

	// Return the attendances
	ResponseJSON(w, attendancess)
}
