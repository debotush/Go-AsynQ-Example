package main

import (
	"github.com/hibiken/asynq"
	"log/slog"
	"sync"
)

type WorkerServer struct {
	server          *asynq.Server
	isPaused        bool
	numberOfWorkers int
}

func initWorkers(pauseChan chan bool) {

	numberOfWorkers := 10
	var wg sync.WaitGroup
	wg.Add(1)
	workerServer := NewWorkerServer(redisAddress, numberOfWorkers)
	go workerServer.workerProcessor(&wg, pauseChan)
	wg.Wait()
}

func NewWorkerServer(address string, numberOfWorkers int) *WorkerServer {
	return &WorkerServer{
		server: asynq.NewServer(
			asynq.RedisClientOpt{Addr: address},
			asynq.Config{
				Concurrency: numberOfWorkers,
			}),
		isPaused:        true,
		numberOfWorkers: numberOfWorkers,
	}
}

func (s *WorkerServer) workerProcessor(wg *sync.WaitGroup, pauseChan chan bool) {
	defer wg.Done()

	for {
		select {
		case val := <-pauseChan:
			if val {
				if s.isPaused {
					slog.Info("Worker server is already paused")
				} else {
					s.stop()
				}
			} else {
				if s.isPaused {
					s.start()
				} else {
					slog.Info("Worker server is already running")
				}
			}
		}
	}
}

func (s *WorkerServer) start() {
	if s.isPaused {
		if s.server == nil {
			s.server = asynq.NewServer(
				asynq.RedisClientOpt{Addr: redisAddress},
				asynq.Config{
					Concurrency: 10,
				})
		}
		s.isPaused = false
		mux := asynq.NewServeMux()
		mux.HandleFunc(task, taskProcessor)
		go func() {
			if err := s.server.Run(mux); err != nil {
				slog.Error("could not start asyncQ worker server: %v", err)
			}
		}()
		slog.Info("started asyncQ worker server")
	}
}

func (s *WorkerServer) stop() {
	if !s.isPaused {
		s.server.Stop()
		s.server = nil
		s.isPaused = true
		slog.Info("Worker server is paused")
	}
}
