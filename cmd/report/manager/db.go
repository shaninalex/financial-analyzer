package manager

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/shaninalex/financial-analyzer/internal/typedefs"

	"github.com/rabbitmq/amqp091-go"
)

func (rm *ReportManager) SaveInitialReport(action typedefs.Action, m amqp091.Delivery) (*typedefs.Report, error) {

	ticker, ok := action.Payload["ticker"].(string)
	if !ok {
		return nil, fmt.Errorf("unable to create report. Payload is corrapted: %v", action.Payload)
	}

	report := &typedefs.Report{
		UserId:    m.Headers["user_id"].(string),
		RequestId: m.Headers["request_id"].(string),
		Ticker:    ticker,
		Status:    false,
	}
	if err := rm.db.ReportCreate(report); err != nil {
		return nil, fmt.Errorf("Error saving new request to db: %v", err)
	}

	return report, nil
}

func (rm *ReportManager) CreateReportData(action typedefs.Action) error {
	report_id := uint(action.Payload["report_id"].(float64))
	data_type := action.Payload["type"].(string)
	data, err := json.Marshal(action.Payload["data"].(interface{}))
	if err != nil {
		log.Println(err)
	}
	err = rm.db.ReportDataCreate(&typedefs.ReportData{
		Type:     data_type,
		Data:     string(data),
		ReportID: report_id,
	})

	if err != nil {
		log.Println(err)
	}

	// TODO: run in goroutine
	rm.db.CheckReportStatus(report_id, data_type)

	return nil
}
