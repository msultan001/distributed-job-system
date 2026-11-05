package queue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Consumer struct {
	redis *redis.Client
}

func NewConsumer(r *redis.Client) *Consumer {
	return &Consumer{
		redis: r,
	}
}

func (c *Consumer) Read(
	ctx context.Context,
	consumername string,
) ([]redis.XStream, error) {
	return c.redis.XReadGroup(
		ctx,
		&redis.XReadGroupArgs{
			Group:    "job-workers",
			Consumer: consumername,
			Streams:  []string{"jobs", ">"},

			Count: 1,
			Block: 0,
		},
	).Result()

}
