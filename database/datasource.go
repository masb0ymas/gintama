package database

import (
	"fmt"
	"strconv"

	"gintama/config"
	"gintama/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Database struct {
	db *gorm.DB
}

var (
	host     = config.Env("DB_HOST", "127.0.01")
	port, _  = strconv.Atoi(config.Env("DB_PORT", "5432"))
	dbname   = config.Env("DB_DATABASE", "db_example")
	username = config.Env("DB_USERNAME", "postgres")
	password = config.Env("DB_PASSWORD", "postgres")
)

func NewDatabase() (*Database, error) {
	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	// Connect to the DB and initialize the DB variable
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// initial migrations
	migrations.Initial(db)

	return &Database{db: db}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}
