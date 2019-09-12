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
	// CHeck if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting the notice!")

	params := mux.Vars(r)

	// Call the handler
	notice, err := models.GetNotice(w, r, params["id"])

	if err != nil {
		http.Error(w, http.StatusText(utils.ErrorCode), utils.ErrorCode)
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

	if err != nil {
		http.Error(w, http.StatusText(utils.ErrorCode), utils.ErrorCode)
		return
	}

	noticesdetails := models.Response{
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
		ResponseJSON(w, "The length notice should be >= 5")
	} else if notice.NoticeID < 1 {
		ResponseJSON(w, "The notice ID should be >= 1")
	} else {
		err := models.AddNotice(notice)
		if err != nil {
			ResponseJSON(w, err)
		} else {
			ResponseJSON(w, "Notice added!")
		}
	}
}
