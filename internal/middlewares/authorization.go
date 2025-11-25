package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"gintama/internal/lib"
	"gintama/internal/lib/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m Middlewares) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := jwt.New(&m.app.Config.App)

		extractToken, err := jwt.ExtractToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("Unauthorized, %s", err.Error()),
			})
			return
		}

		session, err := m.app.Repositories.Session.GetByToken(extractToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("Unauthorized, %s", err.Error()),
			})
			return
		}

		if session.ID != uuid.Nil {
			claims, err := jwt.Verify(extractToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": fmt.Sprintf("Unauthorized, %s", err.Error()),
				})
				return
			}

			if claims.UID != session.UserID.String() {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized, invalid session",
				})
				return
			}

			if claims.Exp < time.Now().Unix() {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized, expired session",
				})
				return
			}

			lib.ContextSetUID(c, uuid.MustParse(claims.UID))
		}

		c.Next()
	}
}
