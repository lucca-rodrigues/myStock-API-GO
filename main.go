package main

import (
	server "github.com/lucca-rodrigues/myStock-API-GO/src/Server"
	database "github.com/lucca-rodrigues/myStock-API-GO/src/database"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}