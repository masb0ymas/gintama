package config

import (
	"fmt"
	"strings"

	"gintama/pkg/constant"

	"github.com/gin-contrib/cors"
)

func Cors() cors.Config {
	// list allowed origins
	allowedOrigins := strings.Join(constant.AllowedOrigins(), ", ")
	fmt.Println("Cors", "Allowed Origins ( "+allowedOrigins+" )")

	config := cors.Config{
		AllowOrigins: constant.AllowedOrigins(),
		// AllowMethods:  "GET, POST, HEAD, PUT, DELETE, PATCH",
		// AllowHeaders:  "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token",
		// ExposeHeaders: "Content-Length",
		// MaxAge:        86400,
	}

	return config
}
