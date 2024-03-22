package typedefs

type ActionType string

var (
	ActionTypeReport       ActionType = "report"
	ActionTypeGeneratePdf  ActionType = "generate_pdf"
	ActionTypeSendEmail    ActionType = "send_email"
	ActionTypeUpdateReport ActionType = "update_report"
)

// Deprecated.
type Action struct {
	Ticker          string     `json:"ticker"`
	Action          ActionType `json:"action"`
	RequestId       string     `json:"request_id"`
	EmailRecepients []string   `json:"email_recepients"`
}

// // Example actions:
// ```go
//
//	var newReport NewAction = NewAction{
//		Action: ActionTypeReport,
//		Payload: map[string]interface{}{
//			"ticker": "AAPL",
//			"client_id": "<uuid>",
//			"user_id": "<uuid>",
//		},
//	}
//
//	var sendEmail NewAction = NewAction{
//		Action: ActionTypeSendEmail,
//		Payload: map[string]interface{}{
//			"report_id":  "<ReportId>",
//			"recepients": []string{"test1@test.com", "test@test.com"},
//		},
//	}
//
//	var generatePdf NewAction = NewAction{
//		Action: ActionTypeGeneratePdf,
//		Payload: map[string]interface{}{
//			"report_id": "<ReportId>",
//		},
//	}
//
//	var updateReport NewAction = NewAction{
//		Action: ActionTypeUpdateReport,
//		Payload: map[string]interface{}{
//			"report_id": "<ReportId>",
//			"type":      "type", // a type of update ( data ) we add to report
//		},
//	}
//
// ```
type NewAction struct {
	Action  ActionType             `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}
