package middlewares

import (
	"fmt"
	"net/http"

	"gintama/internal/lib"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m Middlewares) PermissionAccess(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := lib.ContextGetUID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := m.app.Repositories.User.GetByID(uid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("Unauthorized, permission access failed: %s", err.Error()),
			})
			return
		}

		if user.ID != uuid.Nil && !lib.Contains(roles, user.RoleID.String()) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized, permission access failed: you are not allowed!",
			})
			return
		}

		c.Next()
	}
}
