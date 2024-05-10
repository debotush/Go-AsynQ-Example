package main

import (
	"github.com/hibiken/asynq"
	"log/slog"
)

type AsyncQ func(address string) *asynq.Client

func initQueue(q AsyncQ, address string) *asynq.Client {
	return q(address)
}

func initRedisBackedAsyncQ(address string) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: address})
}

func enqueue(client *asynq.Client, task *asynq.Task) {
	info, err := client.Enqueue(task)
	if err != nil {
		slog.Error("Error enqueueing task: %v", err)
	}
	slog.Info("enqueued task: id="+info.ID+" queue=", info.Payload)
}
