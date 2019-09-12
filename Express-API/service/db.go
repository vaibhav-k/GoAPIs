package service

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// Database initializes the database connection
func Database() *sql.DB {
	db, err = sql.Open("mysql", "phpmyadmin:V@ibhav23@tcp(127.0.0.1:3306)/phpmyadmin")

	if err != nil {
		panic(err.Error())
	}

	// defer db.Close()
	return db
}
