package models

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() {
	var err error

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
}
