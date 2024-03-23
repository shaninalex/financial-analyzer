package typedefs

type ActionType string

var (
	ActionTypeReport       ActionType = "report"
	ActionTypeGeneratePdf  ActionType = "generate_pdf"
	ActionTypeSendEmail    ActionType = "send_email"
	ActionTypeUpdateReport ActionType = "update_report"
)

// // Example actions:
// ```go
//
//	var newReport Action = NewAction{
//		Action: ActionTypeReport,
//		Payload: map[string]interface{}{
//			"ticker": "AAPL",
//			"client_id": "<uuid>",
//			"user_id": "<uuid>",
//		},
//	}
//
//	var sendEmail Action = NewAction{
//		Action: ActionTypeSendEmail,
//		Payload: map[string]interface{}{
//			"report_id":  "<ReportId>",
//			"recepients": []string{"test1@test.com", "test@test.com"},
//		},
//	}
//
//	var generatePdf Action = NewAction{
//		Action: ActionTypeGeneratePdf,
//		Payload: map[string]interface{}{
//			"report_id": "<ReportId>",
//		},
//	}
//
//	var updateReport Action = NewAction{
//		Action: ActionTypeUpdateReport,
//		Payload: map[string]interface{}{
//			"report_id": "<ReportId>",
//			"type":      "type", // a type of update ( data ) we add to report
//		},
//	}
//
// ```
type Action struct {
	Action  ActionType             `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}
