package routes

import (
	"gintama/src/controllers"
	"gintama/src/middlewares"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRoutes() *gin.Engine {
	// Load the .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// set release mode
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Using Middleware
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.RequestIDMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Gintama ( Rest API with Gin Golang )",
			"maintaner": "masb0ymas, <n.fajri@outlook.com>",
			"source":    "https://github.com/masb0ymas/gintama",
			"goVersion": runtime.Version(),
		})
	})

	// Load static file
	r.Static("/public", "./public")

	r.GET("/v1", func(c *gin.Context) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "forbidden, wrong access endpoint",
		})
	})

	// Grouping Routes
	v1 := r.Group("/v1")
	{
		role := new(controllers.RoleController)

		v1.GET("/role", role.GetAll)
	}

	log.Printf("\n\n PORT: %s \n ENV: %s", os.Getenv("PORT"), os.Getenv("ENV"))

	return r
}
