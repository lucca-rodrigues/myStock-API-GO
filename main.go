package main

import "github.com/lucca-rodrigues/myStock-API-GO/src/database"

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}