package utils

import (
	"B20_Backend/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectToDb() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	dbport := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, name, password, dbport)

	Db, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Succesfully Connected to Data Base on port: " + dbport)
	}

	Db.AutoMigrate(&models.Bug{})
	Db.AutoMigrate(&models.ContactInfo{})
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.BugPreference{})

	return Db

}

func GetDb() *gorm.DB {
	return Db
}
