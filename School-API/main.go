package main

import (
	"./models"
	"./service"
	_ "github.com/go-sql-driver/mysql"
)

var err error

func main() {
	// Initialize the database connection
	models.InitDB()

	// Initialize the router
	service.InitServer()
}
