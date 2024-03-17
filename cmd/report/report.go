package main

import (
	"log"

	"github.com/google/uuid"
)

type Report struct {
	ID     string
	UserId string
}

func InitReport(userId string) {
	report := &Report{
		ID:     uuid.New().String(),
		UserId: userId,
	}
	log.Println(report)
}

func (r *Report) GetData(dataType string) {}
func (r *Report) CalculateCustomValues()  {}
func (r *Report) SetComplete()            {}
