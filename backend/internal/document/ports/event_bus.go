package ports

import "context"

type EventBus interface {
	Publish(ctx context.Context, topic string, payload any) error
}
