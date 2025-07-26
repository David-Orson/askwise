package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisEventBus struct {
	Client *redis.Client
}

func NewRedisEventBus(client *redis.Client) *RedisEventBus {
	return &RedisEventBus{Client: client}
}

func (bus *RedisEventBus) Publish(ctx context.Context, topic string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal event payload: %w", err)
	}

	return bus.Client.XAdd(ctx, &redis.XAddArgs{
		Stream: topic,
		Values: map[string]interface{}{
			"data": data,
		},
	}).Err()
}
