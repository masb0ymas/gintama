package controllers

import (
	"gintama/src/helpers"
	"gintama/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func (ctrl RoleController) GetAll(c *gin.Context) {
	var data []models.Role
	var count int64

	queryPage, _ := strconv.Atoi(c.Query("page"))
	queryPageSize, _ := strconv.Atoi(c.Query("pageSize"))

	page := queryPage | 1
	pageSize := queryPageSize | 10

	models.DB.Model(models.Role{}).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&data).
		Count(&count)

	buildResponse := helpers.BuildResponse("data has been received", http.StatusOK, data, count)

	c.JSON(http.StatusOK, buildResponse)
}
