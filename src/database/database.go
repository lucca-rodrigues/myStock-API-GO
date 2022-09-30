package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB(){
	connection := "host=innova-systems.tec.br port=49153 user=finances dbname=myStock-dev sslmode=disable password=financesprod"

	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	config.SetMaxIdleConns(64)

	config.SetMaxOpenConns(64)
	config.SetConnMaxLifetime(time.Minute)
	// migrations.RunMigrations(db)
}