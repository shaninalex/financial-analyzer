
export interface ITickerAction {
    ticker: string
    action: "search"|"process"|"report"
}

export interface IResponseData {
    Ticker: string
    Datasource: string
    Payload: any
}

export interface INotification {
    level: "success"|"error"|"warning"|"info"
    message: string
    read: boolean
    datetime: Date
}

// example return object
// {
// 	"action": "results",
// 	"ticker": "IBM",
// 	"type":   "alph_overview",
// 	"data":   {
// 		"value1": "value1",
// 		"value2": "value2",
// 		"value3": "value3"
// 	}
// }