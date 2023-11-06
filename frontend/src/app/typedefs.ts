
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

