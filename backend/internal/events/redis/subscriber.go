package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func (bus *RedisEventBus) Subscribe(topic string, handler func(context.Context, []byte)) error {
	group := "askwise-consumers"
	consumer := "askwise-" + topic

	ctx := context.Background()

	_ = bus.Client.XGroupCreateMkStream(ctx, topic, group, "0")

	go func() {
		for {
			streams, err := bus.Client.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    group,
				Consumer: consumer,
				Streams:  []string{topic, ">"},
				Count:    1,
				Block:    5 * time.Second,
			}).Result()

			if err != nil && err != redis.Nil {
				log.Printf("error reading from stream %s: %v", topic, err)
				continue
			}

			for _, stream := range streams {
				for _, msg := range stream.Messages {
					data, ok := msg.Values["data"].(string)
					if !ok {
						log.Printf("malformed event in stream %s", topic)
						continue
					}
					handler(ctx, []byte(data))

					_ = bus.Client.XAck(ctx, topic, group, msg.ID)
				}
			}
		}
	}()

	return nil
}
