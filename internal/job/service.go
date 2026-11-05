package job

import (
	"context"

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

	return j.CreateJob(ctx, job)
}
