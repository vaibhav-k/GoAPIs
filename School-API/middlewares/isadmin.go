package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"../utils"
	"github.com/dgrijalva/jwt-go"
)

// IsAdmin checks if the user is an admin
func IsAdmin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, _, err := new(jwt.Parser).ParseUnverified(r.Header.Get("token"), jwt.MapClaims{})
		if err != nil {
			fmt.Println(err)
			return
		}
		typeofuser := token.Claims.(jwt.MapClaims)["type"]

		fmt.Println("The user is a", typeofuser)
		log.Println(r.URL.Path)

		if typeofuser == "admin" {
			f(w, r)
		} else {
			fmt.Println(utils.ErrorAdmin)
		}
	}
}
