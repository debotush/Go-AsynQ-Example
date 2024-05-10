# Example: Asynq. Simple, reliable & efficient distributed task queue for your next Go project
[Asynq](https://github.com/hibiken/asynq) is a Go library for distributing tasks and processing them asynchronously with multiple workers. It's backed by Redis and is designed to scale and distribute easily. Asynq has many features for tasks like scheduling, timeout, retry, etc.

# Example use-case
High-level overview of use-case:
1. Task generator generates and enqueues the tasks into Asynq.
2. The worker server is set for processing multiple tasks concurrently from asynq.
3. The echo server is set up for handling signals. Such as pausing and resuming the task processing workers of asynq.

# Prerequisite 
+ Install docker engine
+ Install docker compose

# Quick start
1. Clone this repository
   ```
   git clone github.com/debotush/Go-AsynQ-Example
   
   ```
2. Start redis
   ```
   docker compose up -d
   
   ```
3. Start the application
   ```
   go run main.go
   
   ```
4. For dashboard UI use this link after the application start.
   [web-brawser](https://localhost:8081)
5. For Resume task processing use this curl command
   ```
   curl -X POST http://localhost:8080/resume
   
   ```
6. For Pause the task processing use this curl command
   ```
   curl -X POST http://localhost:8080/pause
   
   ```
