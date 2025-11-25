package app

import (
	"log/slog"

	"gintama/internal/config"
	"gintama/internal/repositories"
	"gintama/internal/services"
)

type Application struct {
	Config       config.Config
	Logger       *slog.Logger
	Repositories repositories.Repositories
	Services     services.Services
}
