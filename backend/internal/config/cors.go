package config

import (
	"time"

	"github.com/gin-contrib/cors"
)

func NewCorsConfig(origins []string) cors.Config {
	cfg := cors.Config{
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	if len(origins) == 0 {
		cfg.AllowAllOrigins = true
		return cfg
	}

	cfg.AllowOrigins = origins

	return cfg
}