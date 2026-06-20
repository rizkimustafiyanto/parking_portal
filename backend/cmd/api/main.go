package main

import (
	"log"

	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	if err := database.EnsureDatabase(cfg); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	if cfg.AutoMigrate {
		if err := database.Migrate(db, "migrations"); err != nil {
			log.Fatal(err)
		}
	}

	if err := database.Seed(db, cfg); err != nil {
		log.Fatal(err)
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(
		cors.New(
			config.NewCorsConfig(cfg.CORSOrigins),
		),
	)

	routes.Register(router, db, cfg.JWTSecret)

	log.Printf(
		"server started on port %s",
		cfg.AppPort,
	)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
