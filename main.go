package main

import (
	"os"
	"strconv"

	"github.com/shaninalex/financial-analyzer/pkg/datasource"
	"github.com/shaninalex/financial-analyzer/web"
)

var (
	DEBUG        = os.Getenv("DEBUG") // "0" or "1"
	GURU_API_KEY = os.Getenv("GURU_API_KEY")
	ALPH_API_KEY = os.Getenv("ALPH_API_KEY")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	APP_PORT     = os.Getenv("APP_PORT")
)

func main() {

	port, err := strconv.Atoi(APP_PORT)
	if err != nil {
		panic(err)
	}

	// initialize websocket connection
	web.Websocket(port)

	// initialize datasource
	datasource.Init(GURU_API_KEY, ALPH_API_KEY, RABBITMQ_URL)

}
