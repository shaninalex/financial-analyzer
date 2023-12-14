package main

import (
	"os"
	"strconv"
)

var (
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	APP_PORT     = os.Getenv("APP_PORT")
)

func main() {
	app, err := InitApplication()
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(APP_PORT)
	if err != nil {
		panic(err)
	}
	app.Run(port)
}
