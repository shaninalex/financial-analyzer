package typedefs

type GurufocusRequestType string

var (
	GurufocusRequestSummary    GurufocusRequestType = "summary"
	GurufocusRequestFinancials GurufocusRequestType = "financials"
	GurufocusRequestDividend   GurufocusRequestType = "dividend"
	GurufocusRequestPrice      GurufocusRequestType = "price"
	GurufocusRequestKeyratios  GurufocusRequestType = "keyratios"
)
