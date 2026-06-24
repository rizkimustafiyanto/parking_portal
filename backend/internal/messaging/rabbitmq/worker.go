package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type HandlerFunc func(ctx context.Context, routingKey string, body []byte) error

type WorkerConfig struct {
	URL      string
	Exchange string
	Queue    string
	Routes   []string
}

type Worker struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	queue    string
	exchange string
	routes   []string
}

func NewWorker(cfg WorkerConfig) (*Worker, error) {
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

	queue, err := ch.QueueDeclare(
		cfg.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, fmt.Errorf("declare rabbitmq queue: %w", err)
	}

	for _, route := range cfg.Routes {
		if err := ch.QueueBind(queue.Name, route, cfg.Exchange, false, nil); err != nil {
			_ = ch.Close()
			_ = conn.Close()
			return nil, fmt.Errorf("bind rabbitmq queue: %w", err)
		}
	}

	if err := ch.Qos(10, 0, false); err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, fmt.Errorf("set rabbitmq qos: %w", err)
	}

	return &Worker{
		conn:     conn,
		channel:  ch,
		queue:    queue.Name,
		exchange: cfg.Exchange,
		routes:   cfg.Routes,
	}, nil
}

func (w *Worker) Start(ctx context.Context, handler HandlerFunc) error {
	deliveries, err := w.channel.Consume(
		w.queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("start rabbitmq consumer: %w", err)
	}

	log.Printf("rabbitmq worker listening on queue=%s exchange=%s routes=%v", w.queue, w.exchange, w.routes)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg, ok := <-deliveries:
			if !ok {
				return nil
			}

			messageCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			err := handler(messageCtx, msg.RoutingKey, msg.Body)
			cancel()

			if err != nil {
				log.Printf("rabbitmq worker handler failed routing_key=%s: %v", msg.RoutingKey, err)
				_ = msg.Nack(false, true)
				continue
			}

			if err := msg.Ack(false); err != nil {
				log.Printf("rabbitmq worker ack failed routing_key=%s: %v", msg.RoutingKey, err)
			}
		}
	}
}

func (w *Worker) Close() error {
	if w == nil {
		return nil
	}

	if err := w.channel.Close(); err != nil {
		log.Printf("close rabbitmq worker channel: %v", err)
	}

	return w.conn.Close()
}

