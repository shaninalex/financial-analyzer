package main

import (
	"os"
)

var (
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
)

func main() {

	app, err := InitApplication(RABBITMQ_URL)
	if err != nil {
		panic(err)
	}
	app.Run()
}
