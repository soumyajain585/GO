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
	love := Controllers.Love{Partner1: "soum", Partner2: "yash", LovePer: 100}
	fmt.Println(love.LoveCalculator())

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
