package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"../utils"
	"github.com/dgrijalva/jwt-go"
)

// IsStudent checks if the user is a student
func IsStudent(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, _, err := new(jwt.Parser).ParseUnverified(r.Header.Get("token"), jwt.MapClaims{})
		if err != nil {
			fmt.Println(err)
			return
		}
		typeofuser := token.Claims.(jwt.MapClaims)["type"]

		fmt.Println("The user is a", typeofuser)
		log.Println(r.URL.Path)

		if typeofuser == "student" {
			f(w, r)
		} else {
			fmt.Println(utils.ErrorStudent)
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
