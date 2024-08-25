package handler

import (
	"github.com/gin-gonic/gin"
)

func ProductHandler(app *gin.Engine, handler *productHandler) {
	r := app.Group("/product")

	// r.Get("/", handler.listProducts)
	r.POST("/", handler.createProduct)

	// r_id := r.Group("/:id")
	// r_id.Get("/", handler.getProduct)
	// r_id.Put("/", handler.updateProduct)
	// r_id.Delete("/", handler.deleteProduct)
}
