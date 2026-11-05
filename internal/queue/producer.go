package queue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Producer struct {
	redisClient *redis.Client
}

func NewProducer(r *redis.Client) *Producer {
	return &Producer{
		redisClient: r,
	}
}

func (p *Producer) Publish(ctx context.Context, jobID int, payload string) (string, error) {
	return p.redisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: "jobs",
		Values: map[string]interface{}{
			"job_id":  jobID,
			"payload": payload,
		},
	}).Result()

}
