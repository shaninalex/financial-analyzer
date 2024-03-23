package manager

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/shaninalex/financial-analyzer/internal/db"
	"github.com/shaninalex/financial-analyzer/internal/typedefs"
	"golang.org/x/net/context"
)

type ReportManager struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	db         db.IDatabaseRepository
	ctx        context.Context
}

func InitReportManager(
	connection *amqp091.Connection,
	channel *amqp091.Channel,
	db db.IDatabaseRepository,
) (*ReportManager, error) {

	return &ReportManager{
		connection: connection,
		channel:    channel,
		db:         db,
		ctx:        context.Background(),
	}, nil
}

func (rm *ReportManager) ConsumeMessages() {
	messages, err := rm.channel.Consume("q.report", "", true, false, false, false, nil)
	if err != nil {
		log.Printf("Failed to register consumer: %v", err)
	}

	log.Println("ReportManager: start consume messages...")
	for m := range messages {

		var action typedefs.Action
		if err := json.Unmarshal(m.Body, &action); err != nil {
			log.Printf("Error parsing message: %v", err)
			break
		}

		switch m.RoutingKey {
		case "new_report":
			// save initial empty report in db with status=false for now
			report, err := rm.SaveInitialReport(action, m)
			if err != nil {
				log.Println(err)
				continue
			} else {
				// Send message to datasource to initialize gathering information
				err := rm.InitGatheringInformation(m, report)
				if err != nil {
					log.Printf("failed to send gathering information request: %v", err)
				}
			}

		case "update_report":
			err := rm.CreateReportData(action)
			if err != nil {
				log.Printf("Unable to save report data: %v", err)
			}

		case "":
		default:
			log.Println("no routing key provided")
		}
	}
}
