package typedefs

// Deprecated: Use ActionType
type TickerActionType string

var (
	TickerActionTypeSearch  TickerActionType = "search"
	TickerActionTypeProcess TickerActionType = "process"
	TickerActionTypeReport  TickerActionType = "report"
)

// Deprecated: Use ActionType
type ITickerAction struct {
	Ticker string           `json:"ticker"`
	Action TickerActionType `json:"action"`
}

type ActionType string

var (
	ActionTypeReport      ActionType = "report"
	ActionTypeGeneratePdf ActionType = "generate_pdf"
	ActionTypeSendEmail   ActionType = "send_email"
)

type Action struct {
	Ticker         string     `json:"ticker"`
	Action         ActionType `json:"action"`
	RequestId      string     `json:"request_id"`
	EmailRecepient string     `json:"email_recepient"`
}
