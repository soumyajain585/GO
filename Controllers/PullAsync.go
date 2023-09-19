package Controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

func PullDataFromRedis(ctx context.Context, t *asynq.Task) error {
	var data map[string]interface{}
	err := json.Unmarshal(t.Payload(), &data)
	if err != nil {
		fmt.Println("@PullDataFromRedis unmarshl err ", err)
		return err
	}
	fmt.Println("data ", data)
	return nil
}
