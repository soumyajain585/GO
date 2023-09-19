package main

import (
	"fmt"
	"root/Controllers"

	"github.com/hibiken/asynq"
)

func main() {
	//Push Date Into Redis
	Controllers.PushDateIntoRedis()

	//call method
	rect := Controllers.Rectangle{Width: 2, Height: 5}
	fmt.Println("area ", rect.Area())

	//call concurrency
	Controllers.Concurrency()

	//Syncronization
	Controllers.Counter()

	//redis conn
	redisCon := asynq.RedisClientOpt{
		DB:       0,
		Addr:     "localhost:6379",
		Password: "",
	}
	worker := asynq.NewServer(redisCon, asynq.Config{
		Concurrency: 12,
		Queues:      map[string]int{"my-data": 12},
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(Controllers.TaskType, Controllers.PullDataFromRedis)
	worker.Run(mux)
}
