package handlers

import (
	"gintama/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthHandler struct {
	app *app.Application
}

func (h *healthHandler) Check(c *gin.Context) {
	v := gin.H{
		"machineID": h.app.Config.App.MachineID,
		"status":    "ok",
		"systemInfo": map[string]interface{}{
			"debug": h.app.Config.App.Debug,
		},
	}

	c.JSON(http.StatusOK, v)
}
