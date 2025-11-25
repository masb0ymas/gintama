package main

import (
	"net/http"

	"gintama/internal/app"
	"gintama/internal/handlers"
	"gintama/internal/lib/constant"
	"gintama/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine, app *app.Application) {
	h := handlers.New(app)
	m := middlewares.New(app)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/health-check", h.Health.Check)

	authRoutes := r.Group("/v1/auth")
	authRoutes.POST("/sign-up", h.Auth.SignUp)
	authRoutes.POST("/sign-in", h.Auth.SignIn)
	authRoutes.POST("/verify-registration", h.Auth.VerifyRegistration)
	authRoutes.GET("/verify-session", m.Authorization(), h.Auth.VerifySession)
	authRoutes.POST("/sign-out", m.Authorization(), h.Auth.SignOut)

	adminOnly := []string{constant.RoleAdmin}

	sessionRoutes := r.Group("/v1/sessions")
	sessionRoutes.Use(m.Authorization(), m.PermissionAccess(adminOnly))
	sessionRoutes.GET("", h.Session.Index)

	roleRoutes := r.Group("/v1/roles")
	roleRoutes.Use(m.Authorization())
	roleRoutes.GET("", h.Role.Index)
	roleRoutes.GET("/:roleID", h.Role.Show)
	roleRoutes.POST("", m.PermissionAccess(adminOnly), h.Role.Create)
	roleRoutes.PUT("/:roleID", m.PermissionAccess(adminOnly), h.Role.Update)
	roleRoutes.DELETE("/:roleID", m.PermissionAccess(adminOnly), h.Role.Delete)
	roleRoutes.DELETE("/:roleID/soft-delete", m.PermissionAccess(adminOnly), h.Role.SoftDelete)
	roleRoutes.PATCH("/:roleID/restore", m.PermissionAccess(adminOnly), h.Role.Restore)

	userRoutes := r.Group("/v1/users")
	userRoutes.Use(m.Authorization())
	userRoutes.GET("", h.User.Index)
	userRoutes.GET("/:userID", h.User.Show)
	userRoutes.POST("", m.PermissionAccess(adminOnly), h.User.Create)
	userRoutes.PUT("/:userID", m.PermissionAccess(adminOnly), h.User.Update)
	userRoutes.DELETE("/:userID", m.PermissionAccess(adminOnly), h.User.Delete)
	userRoutes.DELETE("/:userID/soft-delete", m.PermissionAccess(adminOnly), h.User.SoftDelete)
	userRoutes.PATCH("/:userID/restore", m.PermissionAccess(adminOnly), h.User.Restore)

	// Not found handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Sorry, HTTP resource you are looking for was not found.",
		})
	})
}
