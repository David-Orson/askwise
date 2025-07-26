package events

import "context"

type EventBus interface {
	Publish(ctx context.Context, topic string, payload any) error
	Subscribe(topic string, handler func(ctx context.Context, data []byte)) error
}
