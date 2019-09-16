package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../utils"
)

// LoginUser logs a user in
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "POST" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the POST method for this route")
	}
	fmt.Println("Logging user in!")

	// Get the login details given in the POST
	decoder := json.NewDecoder(r.Body)
	var user models.User2
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	// Call the handler
	token, er := models.LoginUser2(user)

	if er != "" {
		logintoken := models.Response{
			StatusCode: utils.WrongParam,
			Message:    "Couldn't log in",
			Data:       "Wrong credentials",
		}
		fmt.Println(logintoken.StatusCode)
		ResponseJSON(w, logintoken)
		return
	}

	logintoken := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.LoggedIn,
		Data:       token,
	}

	// Return from the function
	ResponseJSON(w, logintoken)
}
