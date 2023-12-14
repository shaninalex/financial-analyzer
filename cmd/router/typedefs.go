package main

type TickerActionType string

var (
	TickerActionTypeSearch  TickerActionType = "search"
	TickerActionTypeProcess TickerActionType = "process"
	TickerActionTypeReport  TickerActionType = "report"
)

type ITickerAction struct {
	Ticker string           `json:"ticker"`
	Action TickerActionType `json:"action"`
}
