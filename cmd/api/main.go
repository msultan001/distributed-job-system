package main

import (
	"log"

	"github.com/msultan001/distributed-job-system/internal/api"
	"github.com/msultan001/distributed-job-system/internal/queue"
	"github.com/redis/go-redis/v9"
)

func main() {
	client:=redis.NewClient(&redis.Options{

	})
	_=queue.NewProducer(client)
	_ =queue.NewConsumer(client)
	handler := api.NewHandler()
	router := api.NewRouter(handler)
	router.RegisterRoutes()
	api := api.NewAPI("localhost:8000", router)

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
