package jwt

import "gintama/internal/config"

func New(config *config.ConfigApp) *JWT {
	return &JWT{
		config: config,
	}
}
