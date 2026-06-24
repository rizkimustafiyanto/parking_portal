package messaging

import (
	"context"
)

type Publisher interface {
	PublishJSON(ctx context.Context, routingKey string, payload any) error
	Close() error
}

