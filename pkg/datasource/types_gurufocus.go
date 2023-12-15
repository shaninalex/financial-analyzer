package datasource

// This object includes increadible amount of details and for not I will not
// define them all. More details can be found in the official documentation:
// https://www.gurufocus.com/api/stock-data#summary
type GuruFocusSummary struct {
	Summary struct {
		General struct {
			Company string  `json:"company"`
			Price   float64 `json:"price"`
		} `json:"general"`
		Chart struct {
		} `json:"chart"`
		Ratio struct {
		} `json:"ratio"`
		Guru struct {
		} `json:"guru"`
		Insider struct {
		} `json:"insider"`
		CompanyData struct {
			Symbol string `json:"symbol"`
		} `json:"company_data"`
		Estimate struct {
		} `json:"estimate"`
	} `json:"summary"`
}
