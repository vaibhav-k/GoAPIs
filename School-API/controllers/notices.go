package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

// GetNotice gets the notices for a student from the database
func GetNotice(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting the notice!")

	params := mux.Vars(r)

	// Call the handler
	notice, err := models.GetNotice(w, r, params["id"])

	if err != "" {
		noticedetails := models.Response{
			StatusCode: utils.WrongParam,
			Message:    utils.GetFailed,
			Data:       notice,
		}

		// Return the exam detail's
		ResponseJSON(w, noticedetails)
		return
	}

	noticedetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotNotice,
		Data:       notice,
	}

	// Return from the function
	ResponseJSON(w, noticedetails)
}

// GetNotices gets the notices for a student from the database
func GetNotices(w http.ResponseWriter, r *http.Request) {
	// CHeck if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting the notices!")

	// Call the handler
	notices, err := models.GetNotices()

	noticesdetails := models.Response{
		StatusCode: utils.WrongParam,
		Message:    utils.GetFailed,
		Data:       notices,
	}

	if err != "" {
		ResponseJSON(w, noticesdetails)
		return
	}

	noticesdetails = models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotNotice,
		Data:       notices,
	}

	// Return from the function
	ResponseJSON(w, noticesdetails)
}

// AddNotice adds a new notice to the database
func AddNotice(w http.ResponseWriter, r *http.Request) {
	// CHeck if the Method is correct
	if r.Method != "POST" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the POST method for this route")
	}
	fmt.Println("Adding a notice!")
	decoder := json.NewDecoder(r.Body)
	var notice models.Notice
	err := decoder.Decode(&notice)
	if err != nil {
		panic(err)
	}

	// User input validation and calling the handler
	if len(notice.Notice) < 5 {
		noticedetails := models.Response{
			StatusCode: utils.WrongParam,
			Message:    "The length notice should be >= 5",
			Data:       "The length notice should be >= 5",
		}
		ResponseJSON(w, noticedetails)
	} else {
		err := models.AddNotice(notice)
		if err != "" {
			noticedetails := models.Response{
				StatusCode: utils.WrongParam,
				Message:    utils.InsertionFailed,
				Data:       utils.InsertionFailed,
			}
			ResponseJSON(w, noticedetails)
			return
		}
		noticedetails := models.Response{
			StatusCode: utils.SuccessCode,
			Message:    "Notice added!",
			Data:       "Notice added!",
		}
		ResponseJSON(w, noticedetails)
		return
	}
}
