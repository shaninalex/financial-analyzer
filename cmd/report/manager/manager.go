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
		log.Println("Failed to register consumer:")
		log.Println(err)
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
			log.Println("Client id:", m.Headers["client_id"].(string))
			// save initial empty report in db with status=false for now
			rm.SaveInitialReport(action, m)
			// Send message to datasource to initialize gathering information
			rm.InitGatheringInformation(m)
		case "update_report":
			// TODO:
			// This event needs to update report according to the documentation
			// https://github.com/shaninalex/financial-analyzer/issues/53
		case "":
		default:
			log.Println("no routing key provided")
		}
	}
}
