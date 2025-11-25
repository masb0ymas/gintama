package jwt

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func (j *JWT) ExtractToken(c *gin.Context) (string, error) {
	contextQuery := c.Query("token")
	if contextQuery != "" {
		return contextQuery, nil
	}

	contextCookie, _ := c.Cookie("token")
	if contextCookie != "" {
		return contextCookie, nil
	}

	contextHeader := c.GetHeader("Authorization")
	if contextHeader != "" {
		// Bearer <token>
		parts := strings.Split(contextHeader, " ")
		if len(parts) != 2 {
			return "", errors.New("invalid token format")
		}

		if parts[0] != "Bearer" || parts[1] == "" {
			return "", errors.New("invalid token format")
		}

		return parts[1], nil
	}

	return "", errors.New("token not found")
}
