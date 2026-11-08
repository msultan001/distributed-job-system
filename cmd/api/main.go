package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/msultan001/distributed-job-system/internal/api"
	"github.com/msultan001/distributed-job-system/internal/config"
	"github.com/msultan001/distributed-job-system/internal/databasse"
	"github.com/msultan001/distributed-job-system/internal/job"
	"github.com/msultan001/distributed-job-system/internal/queue"
	"github.com/msultan001/distributed-job-system/internal/worker"
	"github.com/redis/go-redis/v9"
)

func main() {
	exit := make(chan os.Signal, 1)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*5),
	)

	defer cancel()
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6739",
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis connection failed: ", err)
	}

	log.Println("Connected to Redis")
	config := config.NewConfig()

	config.Load()

	confi, err := pgxpool.ParseConfig(config.DB_URL)
	if err != nil {
		log.Fatal(err)
		return
	}

	pg, err := pgxpool.NewWithConfig(ctx, confi)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer pg.Close()

	if err := pg.Ping(ctx); err != nil {
		log.Fatal(err)
		return
	}

	db := databasse.NewDb(pg)
	if err := db.DB.Ping(
		ctx,
	); err != nil {
		exit <- syscall.SIGTERM
		log.Fatal(err)
	}
	producer := queue.NewProducer(client)
	consumer := queue.NewConsumer(client)
	worker := worker.NewWorker(1, consumer)
	consumer.CreateGroup(context.Background())
	go worker.Start(context.Background())

	jobService := job.NewJobService(client, producer)
	handler := api.NewHandler(jobService)
	router := api.NewRouter(handler)
	router.RegisterRoutes()
	api := api.NewAPI("localhost:8000", router)

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}

	<-exit
	fmt.Println("shutting down server")

}
