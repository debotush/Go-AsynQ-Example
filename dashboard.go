package main

import (
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func asyncQDashboard() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Asynqmon handler
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/", // RootPath specifies the root for asynqmon app
		RedisConnOpt: asynq.RedisClientOpt{Addr: redisAddress},
	})

	// Register Asynqmon handler with Echo
	e.Any(h.RootPath()+"*", echo.WrapHandler(h))

	// Start the Echo server
	log.Fatal(e.Start(":8081"))

}
