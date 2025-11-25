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

type userHandler struct {
	app *app.Application
}

func (h *userHandler) Index(c *gin.Context) {
	var dto dto.UserPagination

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

	users, meta, err := h.app.Repositories.User.List(opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseMultiData[*models.User]{
		Message: "list data has been retrieved successfully",
		Data:    users,
		Meta: gin.H{
			"total": meta.Total,
		},
	})
}

func (h *userHandler) Show(c *gin.Context) {
	userID, err := lib.ContextParamUUID(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	user, err := h.app.Repositories.User.Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.User]{
		Message: "get data has been retrieved successfully",
		Data:    user,
	})
}

func (h *userHandler) Create(c *gin.Context) {
	var dto dto.UserCreate

	if err := lib.ValidateRequestBody(c, &dto); err != nil {
		switch e := err.(type) {
		case *lib.ErrValidationFailed:
			c.JSON(http.StatusBadRequest, lib.WrapValidationError(e.MessageRecord))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	userID, err := uuid.NewV7()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user := &models.User{
		Base: models.Base{
			ID: userID,
		},
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Phone:     dto.Phone,
		Password:  dto.Password,
		RoleID:    dto.RoleID,
		UploadID:  dto.UploadID,
	}

	err = h.app.Repositories.User.Insert(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.User]{
		Message: "data has been created successfully",
		Data:    user,
	})
}

func (h *userHandler) Update(c *gin.Context) {
	userID, err := lib.ContextParamUUID(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	var dto dto.UserUpdate

	if err := lib.ValidateRequestBody(c, &dto); err != nil {
		switch e := err.(type) {
		case *lib.ErrValidationFailed:
			c.JSON(http.StatusBadRequest, lib.WrapValidationError(e.MessageRecord))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	user, err := h.app.Repositories.User.Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if dto.FirstName != "" {
		user.FirstName = dto.FirstName
	}

	if dto.LastName != nil {
		user.LastName = dto.LastName
	}

	if dto.Phone != nil {
		user.Phone = dto.Phone
	}

	if dto.UploadID != nil {
		user.UploadID = dto.UploadID
	}

	err = h.app.Repositories.User.Update(userID, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.User]{
		Message: "data has been updated successfully",
		Data:    user,
	})
}

func (h *userHandler) Delete(c *gin.Context) {
	userID, err := lib.ContextParamUUID(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	err = h.app.Repositories.User.Delete(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.User]{
		Message: "data has been deleted successfully",
	})
}

func (h *userHandler) SoftDelete(c *gin.Context) {
	userID, err := lib.ContextParamUUID(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	err = h.app.Repositories.User.SoftDelete(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.User]{
		Message: "data has been soft deleted successfully",
	})
}

func (h *userHandler) Restore(c *gin.Context) {
	userID, err := lib.ContextParamUUID(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id must be uuid format",
			"error":   err.Error(),
		})
		return
	}

	err = h.app.Repositories.User.Restore(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.User]{
		Message: "data has been restored successfully",
	})
}
