package models

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// User2 is a struct which stores information for the user
type User2 struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// LoginUser2 logs a user in
func LoginUser2(user User2) (string, string) {
	id := 0
	utype := ""
	tokenString := ""
	var error error

	s := fmt.Sprintf("SELECT `student_id` FROM `school_students` WHERE `email_id` = '%s' AND `password` = '%s'", user.Email, user.Password)
	result, err := db.Query(s)
	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		err = result.Scan(&id)
		if err != nil {
			panic(err.Error())
		}
	}

	if id != 0 {
		utype = "student"
		fmt.Println("User is a", utype)
	} else {
		s = fmt.Sprintf("SELECT `teacher_id` FROM `school_teachers` WHERE `email_id` = '%s' AND `password` = '%s'", user.Email, user.Password)
		result, err = db.Query(s)
		if err != nil {
			fmt.Println(err)
		}

		for result.Next() {
			err = result.Scan(&id)
			if err != nil {
				panic(err.Error())
			}
		}

		if id != 0 {
			utype = "teacher"
			fmt.Println("User is a", utype)
		} else {
			s = fmt.Sprintf("SELECT `admin_id` FROM `school_admin` WHERE `email_id` = '%s' AND `password` = '%s'", user.Email, user.Password)
			result, err = db.Query(s)
			if err != nil {
				fmt.Println(err)
			}

			for result.Next() {
				err = result.Scan(&id)
				if err != nil {
					panic(err.Error())
				}
			}

			if id != 0 {
				utype = "admin"
				fmt.Println("User is a", utype)
			}
		}
	}
	result.Close()

	// Make the token and give it back to the user
	if id != 0 {
		fmt.Printf("Logged in as %s!\n", utype)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"type": utype,
		})
		tokenString, error = token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}
	} else {
		return tokenString, "Wrong credentials"
	}
	return tokenString, ""
}
