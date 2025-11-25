package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gintama/internal/app"
	"gintama/internal/lib"
	"gintama/internal/lib/constant"

	helmet "github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func serve(app *app.Application) error {
	server := gin.Default()

	// Middleware
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	server.Use(helmet.Default())
	server.Use(requestid.New())
	server.Use(lib.RateLimiter())

	// Cors
	server.Use(cors.New(cors.Config{
		AllowOrigins: constant.AllowedOrigins(app),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		MaxAge:       3600,
	}))

	// static file
	server.Static("/static", "./public/static")

	// Initial routes
	routes(server, app)

	// Create HTTP server with Gin handler
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.App.Port),
		Handler:      server,
		IdleTimeout:  time.Minute,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 3 * time.Minute,
	}

	// Start server in a goroutine
	go func() {
		app.Logger.Info("server started on port", "port", app.Config.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Error("failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	// Create channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal
	<-quit
	app.Logger.Info("Received interrupt signal, shutting down...")

	// Graceful shutdown with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		app.Logger.Error("server forced to shutdown", "error", err)
		return err
	}

	app.Logger.Info("server exited gracefully")
	return nil
}
