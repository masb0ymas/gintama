package handler

import (
	"context"
	"fmt"
	"net/http"

	"gintama/database/entity"
	"gintama/pkg/utils"
	"gintama/service"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	ctx     context.Context
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *productHandler {
	return &productHandler{
		ctx:     context.Background(),
		service: service,
	}
}

func toStoreProduct(p *entity.ProductReq) *entity.Product {
	return &entity.Product{
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
	}
}

func (h *productHandler) createProduct(c *gin.Context) {
	p := new(entity.ProductReq)

	if code, message, errors := utils.ParseFormDataAndValidate(c, p); errors != nil {
		response := utils.FailureResponse(code, message, errors)
		c.JSON(http.StatusOK, response)
		return
	}

	product, err := h.service.CreateProduct(h.ctx, toStoreProduct(p))
	if err != nil {
		fmt.Println(err)
		response := utils.FailureResponse(http.StatusInternalServerError, "error creating product", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.SuccessResponse(http.StatusOK, "data has been added", product)
	c.JSON(http.StatusOK, response)
}
