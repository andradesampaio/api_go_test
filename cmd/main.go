package main

import (
	"api-go-test/internal/database"
	"api-go-test/internal/router"
)

func main() {
	database.Connect()
	router.HandleRequests()
}

