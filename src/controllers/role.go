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
		response := helpers.ErrorResponse(http.StatusBadRequest, "error to get roles")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := gin.H{"data": data, "total": total}
	buildResponse := helpers.BuildResponse(http.StatusOK, "data has been received", response)

	c.JSON(http.StatusOK, buildResponse)

}

// Find By Id
func (h *roleController) FindById(c *gin.Context) {
	var input schema.RoleByIdSchema

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.ErrorResponse(http.StatusBadRequest, "failed to get role")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data, err := h.service.FindById(input)
	if err != nil {
		response := helpers.ErrorResponse(http.StatusBadRequest, "data not found or has been deleted")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := gin.H{"data": data}
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
		response := helpers.ErrorResponse(http.StatusBadRequest, "failed to create role")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(http.StatusCreated, "success to create role", data)
	c.JSON(http.StatusCreated, response)
}

// Update
func (h *roleController) Update(c *gin.Context) {
	var params schema.RoleByIdSchema
	var input schema.RoleSchema

	errUri := c.ShouldBindUri(&params)
	if errUri != nil {
		response := helpers.ErrorResponse(http.StatusBadRequest, "failed to get role")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.BuildResponse(http.StatusUnprocessableEntity, "failed to update role", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateData, err := h.service.Update(params, input)
	if err != nil {
		response := helpers.ErrorResponse(http.StatusBadRequest, "failed to update role")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(http.StatusOK, "success to update role", updateData)
	c.JSON(http.StatusOK, response)
}

// Delete
func (h *roleController) Delete(c *gin.Context) {
	var params schema.RoleByIdSchema

	errUri := c.ShouldBindUri(&params)
	if errUri != nil {
		response := helpers.ErrorResponse(http.StatusBadRequest, "failed to get role")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := h.service.Delete(params)
	if err != nil {
		response := helpers.ErrorResponse(http.StatusBadRequest, "data not found or has been deleted")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(http.StatusOK, "success to deleted role", nil)
	c.JSON(http.StatusOK, response)
}
