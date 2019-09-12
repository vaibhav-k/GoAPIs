package models

import (
	"fmt"

	"../utils"
	"github.com/dgrijalva/jwt-go"
)

// User is a struct which stores information for the user
type User struct {
	Type     string `json:"type,omitempty"`
	Password string `json:"password,omitempty"`
}

// LoginUser logs a user in
func LoginUser(user User) (string, error) {
	id := 0
	s := ""
	tokenString := ""
	var error error

	// Query the DB according to the type specified
	if user.Type == "student" {
		s = fmt.Sprintf("SELECT student_id FROM `school_students` WHERE `password` = '%s'", user.Password)
	} else if user.Type == "teacher" {
		s = fmt.Sprintf("SELECT teacher_id FROM `school_teachers` WHERE `password` = '%s'", user.Password)
	} else if user.Type == "admin" {
		s = fmt.Sprintf("SELECT admin_id FROM `school_admin` WHERE `password` = '%s'", user.Password)
	} else {
		return utils.NoType, nil
	}
	result, err := db.Query(s)

	if err != nil {
		panic(err.Error())
	} else if result != nil {
		for result.Next() {
			err := result.Scan(&id)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	result.Close()

	// Make the token and give it back to the user
	if id != 0 {
		fmt.Printf("Logged in as %s!\n", user.Type)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"type": user.Type,
		})
		tokenString, error = token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}
	}
	return tokenString, nil
}
