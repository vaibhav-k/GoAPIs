package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../utils"
	"github.com/dgrijalva/jwt-go"
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

// LoginUserCombined logs a user in
func LoginUserCombined(w http.ResponseWriter, r *http.Request) {
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

	utype := ""

	var db *sql.DB

	uname := "vaibhavk"
	pword := "V@ibhav23"
	ip := "www.db4free.net"
	dbname := "expressapi"

	// Open the DB connection
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", uname, pword, ip, dbname))

	// Check for errors
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to the database!")
	}

	s := fmt.Sprintf("SELECT `type` FROM `school_users` WHERE `email` = '%s' AND `password` = '%s'", user.Email, user.Password)
	result, err := db.Query(s)
	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		err = result.Scan(&utype)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("User is a", utype)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type": utype,
	})
	tokenString, _ := token.SignedString([]byte("secret"))

	logintoken := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.LoggedIn,
		Data:       tokenString,
	}
	ResponseJSON(w, logintoken)
}
