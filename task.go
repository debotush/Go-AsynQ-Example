package main

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log/slog"
	"strconv"
)

type TaskPayload struct {
	ID      int
	Message string
}

func NewTaskPayload(id int, message string) *asynq.Task {
	payload := TaskPayload{
		ID:      id,
		Message: message,
	}
	jsonPayload, _ := json.Marshal(payload)
	return asynq.NewTask(task, jsonPayload)
}

func taskGenerator(client *asynq.Client) {
	for i := 1; ; i++ {
		message := "task: " + strconv.Itoa(i)
		task := NewTaskPayload(i, message)
		enqueue(client, task)
	}
}

func taskProcessor(ctx context.Context, t *asynq.Task) error {
	payload := TaskPayload{}
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		slog.Error("failed to unmarshal payload: %s", err)
		return err
	}
	slog.Info("task processed id=%d message=%s\n", payload.ID, payload.Message)
	return nil
}
