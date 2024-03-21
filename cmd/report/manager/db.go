package manager

import (
	"log"

	"github.com/shaninalex/financial-analyzer/internal/typedefs"

	"github.com/rabbitmq/amqp091-go"
)

func (rm *ReportManager) SaveInitialReport(action typedefs.Action, m amqp091.Delivery) {
	go func() {
		if err := rm.db.ReportCreate(&typedefs.Report{
			UserId:    m.Headers["user_id"].(string),
			RequestId: m.Headers["request_id"].(string),
			Ticker:    action.Ticker,
		}); err != nil {
			log.Printf("Error saving new request to db: %v", err)
		}
	}()
}
