package messaging

import (
	"context"
	"encoding/json"
	"log"
)

type EventLogger interface {
	Start(ctx context.Context, handler func(ctx context.Context, routingKey string, body []byte) error) error
	Close() error
}

func NewLoggingHandler() func(ctx context.Context, routingKey string, body []byte) error {
	return func(ctx context.Context, routingKey string, body []byte) error {
		var payload map[string]any
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Printf("rabbitmq event received routing_key=%s raw=%s", routingKey, string(body))
			return nil
		}

		log.Printf("rabbitmq event received routing_key=%s payload=%v", routingKey, payload)
		return nil
	}
}

