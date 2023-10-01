package main

import (
	"os"
	"strconv"
)

var (
	APP_PORT     = os.Getenv("APP_PORT")
	GURU_API_KEY = os.Getenv("GURU_API_KEY")
	ALPH_API_KEY = os.Getenv("ALPH_API_KEY")
)

func main() {
	api, err := InitializeAPI(GURU_API_KEY, ALPH_API_KEY)
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(APP_PORT)
	if err != nil {
		panic(err)
	}

	api.Run(port)
}
