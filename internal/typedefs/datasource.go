package typedefs

type GurufocusRequestType string

var (
	GurufocusRequestSummary    GurufocusRequestType = "summary"
	GurufocusRequestFinancials GurufocusRequestType = "financials"
	GurufocusRequestDividend   GurufocusRequestType = "dividend"
)

type AlphavantageRequestType string

var (
	AlphavantageRequestOverview AlphavantageRequestType = "overview"
	AlphavantageRequestEarnings AlphavantageRequestType = "earnings"
	AlphavantageRequestCashFlow AlphavantageRequestType = "cashflow"
)
