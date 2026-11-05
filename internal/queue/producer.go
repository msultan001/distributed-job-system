package queue

import (
	"context"
	"encoding/json"

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

func (p *Producer) Publish(ctx context.Context, jobID int, payload interface{}) (string, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return p.redisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: "jobs",
		Values: map[string]interface{}{
			"job_id":  jobID,
			"payload": string(data),
		},
	}).Result()

}
