package routes

import (
	"gintama/src/controllers"
	"gintama/src/helpers"
	"gintama/src/middlewares"
	"gintama/src/models"
	"gintama/src/repository"
	"gintama/src/services"
	"log"
	"net/http"
	"os"
	"runtime"

	helmet "github.com/danielkov/gin-helmet"
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
	r.Use(helmet.Default())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.RequestIDMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":      http.StatusOK,
			"message":   "Gintama ( Rest API with Gin Golang )",
			"maintaner": "masb0ymas, <n.fajri@outlook.com>",
			"source":    "https://github.com/masb0ymas/gintama",
			"goVersion": runtime.Version(),
		})
	})

	// Load static file
	r.Static("/public", "./public")

	r.GET("/v1", func(c *gin.Context) {
		response := helpers.ErrorResponse(http.StatusForbidden, "forbidden, wrong access endpoint")
		c.JSON(http.StatusForbidden, response)
	})

	db := models.GetDB()

	// List Repository
	roleRepository := repository.RoleRepository(db)

	// List Service
	roleService := services.RoleService(roleRepository)

	// List Controller
	roleController := controllers.RoleController(roleService)

	// Grouping Routes
	v1 := r.Group("/v1")
	{
		v1.GET("/role", roleController.GetAll)
		v1.GET("/role/:id", roleController.FindById)
		v1.POST("/role", roleController.CreateRole)
		v1.PUT("/role/:id", roleController.Update)
		v1.DELETE("/role/:id", roleController.Delete)
	}

	log.Printf("\n\n PORT: %s \n ENV: %s", os.Getenv("PORT"), os.Getenv("ENV"))

	return r
}
