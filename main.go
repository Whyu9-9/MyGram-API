package main

import (
	"mygram-api/database"
	"mygram-api/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
