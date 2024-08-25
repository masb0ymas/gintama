package main

import (
	"log"
	"strconv"

	"gintama/config"
	"gintama/database"
	"gintama/pkg/utils"
	"gintama/routes"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

var (
	dbname = config.Env("DB_DATABASE", "db_example")
	rl, _  = strconv.Atoi(config.Env("APP_RATE_LIMIT", "10"))
	port   = config.Env("APP_PORT", "8000")
)

func main() {
	limit := ratelimit.New(rl)

	// database instance
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	log.Printf("successfully connected to database %v", dbname)

	app := gin.New()

	// default middleware
	app.Use(cors.New(config.Cors()))
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(helmet.Default())
	app.Use(utils.RateLimit(limit))
	app.Use(requestid.New())

	// static file
	app.Static("/", "./public")

	// register routes
	routes.RegisterRoutes(app, db.GetDB())

	// listening app
	log.Fatal(app.Run(":" + port))
}
