/*
This package used for getting information from financials APIs such as
GuruFocus, Alphavantage etc.
*/
package datasource

// Usage example: data, err: := app.Datasource.Gurufocus.summary("AAPL")
type Datasource struct {
	Gurufocus    *Gurufocus
	Alphavantage *Alphavantage
}

func InitializeDatasource(gfApiKey, alphApiKey string) *Datasource {
	return &Datasource{
		Gurufocus:    InitGurufocus(gfApiKey),
		Alphavantage: InitAlphavantage(alphApiKey),
	}
}
