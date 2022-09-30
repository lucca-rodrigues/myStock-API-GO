package main

import (
	"github.com/lucca-rodrigues/myStock-API-GO/src/database"
	"github.com/lucca-rodrigues/myStock-API-GO/src/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}