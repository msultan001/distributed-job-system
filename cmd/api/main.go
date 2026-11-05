package main

import (
	"context"
	"log"

	"github.com/msultan001/distributed-job-system/internal/api"
	"github.com/msultan001/distributed-job-system/internal/job"
	"github.com/msultan001/distributed-job-system/internal/queue"
	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6739",
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Redis connection failed: ", err)
	}

	log.Println("Connected to Redis")
	producer:= queue.NewProducer(client)
	_= queue.NewConsumer(client)
	jobService:=job.NewJobService(client, producer)
	handler := api.NewHandler(jobService)
	router := api.NewRouter(handler)
	router.RegisterRoutes()
	api := api.NewAPI("localhost:8000", router)

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
