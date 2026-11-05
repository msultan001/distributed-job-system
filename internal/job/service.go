package job

import (
	"context"
	"fmt"

	"github.com/msultan001/distributed-job-system/internal/queue"
	"github.com/redis/go-redis/v9"
)

type JobService struct {
	r        *redis.Client
	producer *queue.Producer
}

func NewJobService(redisClient *redis.Client, p *queue.Producer) *JobService {
	return &JobService{
		r:        redisClient,
		producer: p,
	}
}

func (j *JobService) CreateJob(ctx context.Context, job *Job) error {
	fmt.Println("Creating job:", job.ID)
	id, err := j.producer.Publish(ctx, job.ID, job.Payload)
	if err != nil {
		fmt.Println("Publish failed: ", err)
		return err
	}
	fmt.Println("Published Redis message", id)
	return nil
}
