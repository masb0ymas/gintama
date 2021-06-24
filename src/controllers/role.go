package controllers

import (
	"gintama/src/helpers"
	"gintama/src/models"
	"gintama/src/schema"
	"gintama/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type roleController struct {
	service services.IService
}

func RoleController(service services.IService) *roleController {
	return &roleController{service}
}

// Get All
func (h *roleController) GetAll(c *gin.Context) {
	var queryFiltered models.QueryFiltered

	queryFiltered.Page = c.Query("page")
	queryFiltered.PageSize = c.Query("pageSize")

	data, total, err := h.service.GetAll(queryFiltered)
	if err != nil {
		response := helpers.BuildResponse(http.StatusBadRequest, "error to get roles", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := gin.H{"data": data, "total": total}
	buildResponse := helpers.BuildResponse(http.StatusOK, "data has been received", response)

	c.JSON(http.StatusOK, buildResponse)

}

// Create
func (h *roleController) CreateRole(c *gin.Context) {
	var input schema.RoleSchema

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.BuildResponse(http.StatusUnprocessableEntity, "failed to create role", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.service.CreateRole(input)
	if err != nil {
		response := helpers.BuildResponse(http.StatusBadRequest, "failed to create role", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(http.StatusCreated, "success to create role", data)
	c.JSON(http.StatusCreated, response)
}
