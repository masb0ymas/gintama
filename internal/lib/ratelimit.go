package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	var (
		limit = rate.Limit(2) // 2 tokens per second
		burst = 4             // 4 requests
	)

	limiter := rate.NewLimiter(limit, burst)
	return func(c *gin.Context) {

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests",
			})
			return
		}

		c.Next()
	}
}
