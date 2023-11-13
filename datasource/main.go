package main

import (
	"os"
	"strconv"
)

var (
	DEBUG        = os.Getenv("DEBUG") // "0" or "1"
	APP_PORT     = os.Getenv("APP_PORT")
	GURU_API_KEY = os.Getenv("GURU_API_KEY")
	ALPH_API_KEY = os.Getenv("ALPH_API_KEY")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
)

func main() {

	s_debug, err := strconv.Atoi(APP_PORT)
	if err != nil {
		panic(err)
	}
	b_debug := false

	if s_debug > 0 {
		b_debug = true
	}

	api, err := InitializeAPI(GURU_API_KEY, ALPH_API_KEY, b_debug)
	if err != nil {
		panic(err)
	}

	defer api.MQConnection.Close()
	defer api.MQChannel.Close()

	go api.ConsumeRabbitMessages()

	port, err := strconv.Atoi(APP_PORT)
	if err != nil {
		panic(err)
	}
	api.Run(port)
}
