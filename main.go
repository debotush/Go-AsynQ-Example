package main

import (
	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	pauseChannel := make(chan bool)
	//resumeChannel <- true

	qClient := initQueue(initRedisBackedAsyncQ, redisAddress)
	defer func(qClient *asynq.Client) {
		err := qClient.Close()
		if err != nil {

		}
	}(qClient)

	go taskGenerator(qClient)

	go echoServer(pauseChannel)

	go asyncQDashboard()

	initWorkers(pauseChannel)

}

func echoServer(pauseChannel chan bool) {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.POST("/pause", func(c echo.Context) error {
		pauseChannel <- true
		return c.JSON(http.StatusOK, "paused")
	})

	e.POST("/resume", func(c echo.Context) error {
		pauseChannel <- false

		return c.JSON(http.StatusOK, "resume")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
