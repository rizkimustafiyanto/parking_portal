package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"backend/internal/messaging"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PublisherConfig struct {
	URL      string
	Exchange string
}

type publisher struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
}

func NewPublisher(cfg PublisherConfig) (messaging.Publisher, error) {
	conn, err := amqp.Dial(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("connect rabbitmq: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("open rabbitmq channel: %w", err)
	}

	if err := ch.ExchangeDeclare(
		cfg.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, fmt.Errorf("declare rabbitmq exchange: %w", err)
	}

	return &publisher{
		conn:     conn,
		channel:  ch,
		exchange: cfg.Exchange,
	}, nil
}

func (p *publisher) PublishJSON(ctx context.Context, routingKey string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal rabbitmq payload: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return p.channel.PublishWithContext(ctx, p.exchange, routingKey, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		Body:         body,
	})
}

func (p *publisher) Close() error {
	if p == nil {
		return nil
	}

	if err := p.channel.Close(); err != nil {
		log.Printf("close rabbitmq channel: %v", err)
	}

	return p.conn.Close()
}
