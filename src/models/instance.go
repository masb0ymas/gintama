package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	*gorm.DB
}

type QueryFiltered struct {
	Page     string
	PageSize string
	Order    string
}

func ConnectDatabase() {
	// Load the .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_database := os.Getenv("DB_DATABASE")
	db_user := os.Getenv("DB_USERNAME")
	db_pass := os.Getenv("DB_PASSWORD")

	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_database + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, errSql := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errSql != nil {
		log.Fatal(errSql.Error())
	}

	fmt.Println("Connection has been established successfully.")

	// List Auto Migrate Table from struct model
	database.AutoMigrate(&Role{})

	DB = database
}

func GetDB() *gorm.DB {
	return DB
}
