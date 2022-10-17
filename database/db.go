package database

import (
	"fmt"
	"log"
	"mygram-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	dbPort   = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "mygram_api"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, dbPort, user, dbname, password)
	con := config

	db, err = gorm.Open(postgres.Open(con), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Successfully connected to database")
	db.Debug().AutoMigrate(&models.User{}, &models.Photo{}, &models.SocialMedia{}, &models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
