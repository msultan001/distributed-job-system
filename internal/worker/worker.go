package worker

import (
	"context"

	"github.com/msultan001/distributed-job-system/internal/queue"
)

type Worker struct {
	ID       int
	Consumer *queue.Consumer
}

func NewWorker(id int, consumer *queue.Consumer) *Worker {
	return &Worker{
		ID:       id,
		Consumer: consumer,
	}
}

func (w *Worker) Start(ctx context.Context) {
	for {
		_, err:=w.Consumer.Read(ctx, "")
		if err!=nil{
			return
		}
	}
}
