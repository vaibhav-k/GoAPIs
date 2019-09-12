package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// IsLoggedIn checks if the user is logged in
func IsLoggedIn(f http.HandlerFunc) http.HandlerFunc {
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
		} else if typeofuser == "teacher" {
			f(w, r)
		} else if typeofuser == "admin" {
			f(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
