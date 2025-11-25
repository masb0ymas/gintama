package handlers

import (
	"net/http"

	"gintama/internal/app"
	"gintama/internal/dto"
	"gintama/internal/lib"
	"gintama/internal/models"
	"gintama/internal/repositories"
	"gintama/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type roleHandler struct {
	app *app.Application
}

func (h *roleHandler) Index(c *gin.Context) {
	var dto dto.RolePagination

	if err := lib.ValidateRequestQuery(c, &dto); err != nil {
		switch e := err.(type) {
		case *lib.ErrValidationFailed:
			c.JSON(http.StatusBadRequest, lib.WrapValidationError(e.MessageRecord))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	opts := &repositories.QueryOptions{
		Offset: dto.Offset,
		Limit:  dto.Limit,
	}

	roles, meta, err := h.app.Repositories.Role.List(opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseMultiData[*models.Role]{
		Message: "list data has been retrieved successfully",
		Data:    roles,
		Meta: gin.H{
			"total": meta.Total,
		},
	})
}

func (h *roleHandler) Show(c *gin.Context) {
	roleID, err := lib.ContextParamUUID(c, "roleID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid role id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	role, err := h.app.Repositories.Role.Get(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.Role]{
		Message: "get data has been retrieved successfully",
		Data:    role,
	})
}

func (h *roleHandler) Create(c *gin.Context) {
	var dto dto.RoleCreate

	if err := lib.ValidateRequestBody(c, &dto); err != nil {
		switch e := err.(type) {
		case *lib.ErrValidationFailed:
			c.JSON(http.StatusBadRequest, lib.WrapValidationError(e.MessageRecord))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	roleID, err := uuid.NewV7()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	role := &models.Role{
		Base: models.Base{
			ID: roleID,
		},
		Name: dto.Name,
	}

	err = h.app.Repositories.Role.Insert(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.Role]{
		Message: "data has been created successfully",
		Data:    role,
	})
}

func (h *roleHandler) Update(c *gin.Context) {
	roleID, err := lib.ContextParamUUID(c, "roleID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid role id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	var dto dto.RoleUpdate

	if err := lib.ValidateRequestBody(c, &dto); err != nil {
		switch e := err.(type) {
		case *lib.ErrValidationFailed:
			c.JSON(http.StatusBadRequest, lib.WrapValidationError(e.MessageRecord))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	role, err := h.app.Repositories.Role.Get(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if dto.Name != "" {
		role.Name = dto.Name
	}

	err = h.app.Repositories.Role.Update(roleID, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.Role]{
		Message: "data has been updated successfully",
		Data:    role,
	})
}

func (h *roleHandler) Delete(c *gin.Context) {
	roleID, err := lib.ContextParamUUID(c, "roleID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid role id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	err = h.app.Repositories.Role.Delete(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.Role]{
		Message: "data has been deleted successfully",
	})
}

func (h *roleHandler) SoftDelete(c *gin.Context) {
	roleID, err := lib.ContextParamUUID(c, "roleID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid role id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	err = h.app.Repositories.Role.SoftDelete(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.Role]{
		Message: "data has been soft deleted successfully",
	})
}

func (h *roleHandler) Restore(c *gin.Context) {
	roleID, err := lib.ContextParamUUID(c, "roleID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid role id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	err = h.app.Repositories.Role.Restore(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.Role]{
		Message: "data has been restored successfully",
	})
}
