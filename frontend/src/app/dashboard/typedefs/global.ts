export const ActionTypeReport = "report"
export const ActionTypeGeneratePdf = "generate_pdf"
export const ActionTypeSendEmail = "send_email"

export type ActionType = typeof ActionTypeReport | typeof ActionTypeGeneratePdf | typeof ActionTypeSendEmail

export interface ISocketAction {
  action: ActionType
  payload: any
}

export interface IResponseData {
  data: any
  ticker: string
  type: string
}

export interface INotification {
  level: "success" | "error" | "warning" | "info"
  message: string
  read: boolean
  datetime: Date
}

export interface IBaseWebsocket {
  action: "notification" | "data_result"
  payload: IResponseData | INotification
}

export interface IFinancials {
  amount: string
  currency: string
  ex_date: string
  pay_date: string
  record_date: string
  type: string
}
