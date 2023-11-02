
export interface ITickerAction {
    ticker: string
    action: "search"|"process"|"report"
}

export interface IResponseData {
    Ticker: string
    Datasource: string
    Payload: any
}

