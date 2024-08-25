package routes

import (
	"gintama/database/repository"
	"gintama/handler"
	"gintama/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(app *gin.Engine, db *gorm.DB) {
	// product endpoint
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	handler.ProductHandler(app, productHandler)
}
