package main

import (
	"context"
	"log"

	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/messaging"
	"backend/internal/messaging/rabbitmq"
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

	var publisher messaging.Publisher
	var worker *rabbitmq.Worker
	if cfg.RabbitMQEnabled {
		publisher, err = rabbitmq.NewPublisher(rabbitmq.PublisherConfig{
			URL:      cfg.RabbitMQURL,
			Exchange: cfg.RabbitMQExchange,
		})
		if err != nil {
			log.Printf("rabbitmq disabled after connection failure: %v", err)
		} else {
			defer publisher.Close()
		}

		worker, err = rabbitmq.NewWorker(rabbitmq.WorkerConfig{
			URL:      cfg.RabbitMQURL,
			Exchange: cfg.RabbitMQExchange,
			Queue:    "portal_digital.events.worker",
			Routes: []string{
				"invoice.created",
				"payment.completed",
			},
		})
		if err != nil {
			log.Printf("rabbitmq worker disabled after connection failure: %v", err)
		} else {
			defer worker.Close()
		}
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

	routes.Register(router, db, cfg.JWTSecret, publisher)

	if worker != nil {
		go func() {
			if err := worker.Start(context.Background(), messaging.NewLoggingHandler()); err != nil && err != context.Canceled {
				log.Printf("rabbitmq worker stopped: %v", err)
			}
		}()
	}

	log.Printf("server started on port %s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
