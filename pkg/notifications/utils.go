package main

import (
	"log"
)

func failOnError(err error) {
	if err != nil {
		log.Panicf("%s", err)
	}
}
