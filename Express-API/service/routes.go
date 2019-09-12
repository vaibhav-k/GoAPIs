package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"../model"
	// "github.com/auth0-community/auth0"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Route handles all possible routes for the API
func Route(db *sql.DB) {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	dbCon := model.DB{db}
	router.HandleFunc("/register", dbCon.RegisterUser).Methods("POST")
	router.HandleFunc("/login", dbCon.LoginUser).Methods("POST")
	router.HandleFunc("/cart/{username}", dbCon.AddToCart).Methods("POST")
	router.HandleFunc("/cart/{username}", dbCon.ViewCart).Methods("GET")
	router.HandleFunc("/products", dbCon.GetProducts).Methods("GET")
	router.HandleFunc("/products", dbCon.CreateProducts).Methods("POST")
	router.HandleFunc("/products/{id}", dbCon.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", dbCon.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", dbCon.DeleteProduct).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
