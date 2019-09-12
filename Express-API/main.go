package main

import (
	"database/sql"

	"./service"
)

var db *sql.DB
var err error

func main() {
	db := service.Database()
	service.Route(db)
}
