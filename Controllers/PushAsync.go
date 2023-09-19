package Controllers

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

var TaskType = "user:data"

func CreateJob() *asynq.Task {
	var mydata = map[string]interface{}{
		"name": "soumya",
	}
	payload, err := json.Marshal(mydata)
	if err != nil {
		fmt.Println("@CreateJob Marshal error : ", err)
		return nil
	}
	return asynq.NewTask(TaskType, payload)
}

func PushDateIntoRedis() {
	task := CreateJob()
	redis := asynq.NewClient(asynq.RedisClientOpt{DB: 0, Addr: "localhost:6379", Password: ""})
	_, err := redis.Enqueue(task, asynq.Queue("my-data"), asynq.ProcessIn(1))
	if err != nil {
		fmt.Println("@PushDateIntoRedis error while pushing data into queue: ", err)
		return
	}
	fmt.Println("@PushDateIntoRedis push data into queue sucessfully!!")
}
