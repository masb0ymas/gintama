package handlers

import (
	"gintama/internal/app"
	"gintama/internal/dto"
	"gintama/internal/lib"
	"gintama/internal/models"
	"gintama/internal/repositories"
	"gintama/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	app *app.Application
}

func (h *sessionHandler) Index(c *gin.Context) {
	var dto dto.SessionPagination

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

	sessions, meta, err := h.app.Repositories.Session.List(opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseMultiData[*models.Session]{
		Message: "list data has been retrieved successfully",
		Data:    sessions,
		Meta: gin.H{
			"total": meta.Total,
		},
	})
}
