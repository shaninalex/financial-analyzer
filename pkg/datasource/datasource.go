/*
This package used for getting information from financials APIs such as
GuruFocus, Alphavantage etc.
*/
package datasource

import "github.com/shaninalex/financial-analyzer/pkg/datasource/data"

// Usage example: data, err: := app.Datasource.Gurufocus.summary("AAPL")
type Datasource struct {
	Gurufocus    *data.GuruFocus
	Alphavantage *data.Alphavantage
}

func InitializeDatasource(gfApiKey, alphApiKey string, debug bool) *Datasource {
	return &Datasource{
		Gurufocus:    data.InitGurufocus(gfApiKey, debug),
		Alphavantage: data.InitAlphavantage(alphApiKey, debug),
	}
}
