package main

type AlphavantageOverview struct {
	Symbol                     string `json:"Symbol"`
	AssetType                  string `json:"AssetType"`
	Name                       string `json:"Name"`
	Description                string `json:"Description"`
	CIK                        string `json:"CIK"`
	Exchange                   string `json:"Exchange"`
	Currency                   string `json:"Currency"`
	Country                    string `json:"Country"`
	Sector                     string `json:"Sector"`
	Industry                   string `json:"Industry"`
	Address                    string `json:"Address"`
	FiscalYearEnd              string `json:"FiscalYearEnd"`
	LatestQuarter              string `json:"LatestQuarter"`
	MarketCapitalization       string `json:"MarketCapitalization"`
	EBITDA                     string `json:"EBITDA"`
	PERatio                    string `json:"PERatio"`
	PEGRatio                   string `json:"PEGRatio"`
	BookValue                  string `json:"BookValue"`
	DividendPerShare           string `json:"DividendPerShare"`
	DividendYield              string `json:"DividendYield"`
	EPS                        string `json:"EPS"`
	RevenuePerShareTTM         string `json:"RevenuePerShareTTM"`
	ProfitMargin               string `json:"ProfitMargin"`
	OperatingMarginTTM         string `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          string `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          string `json:"ReturnOnEquityTTM"`
	RevenueTTM                 string `json:"RevenueTTM"`
	GrossProfitTTM             string `json:"GrossProfitTTM"`
	DilutedEPSTTM              string `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY string `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  string `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         string `json:"AnalystTargetPrice"`
	TrailingPE                 string `json:"TrailingPE"`
	ForwardPE                  string `json:"ForwardPE"`
	PriceToSalesRatioTTM       string `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           string `json:"PriceToBookRatio"`
	EVToRevenue                string `json:"EVToRevenue"`
	EVToEBITDA                 string `json:"EVToEBITDA"`
	Beta                       string `json:"Beta"`
	Five2WeekHigh              string `json:"52WeekHigh"`
	Five2WeekLow               string `json:"52WeekLow"`
	Five0DayMovingAverage      string `json:"50DayMovingAverage"`
	Two00DayMovingAverage      string `json:"200DayMovingAverage"`
	SharesOutstanding          string `json:"SharesOutstanding"`
	DividendDate               string `json:"DividendDate"`
	ExDividendDate             string `json:"ExDividendDate"`
}

type AlphavantageEarnings struct {
	Symbol         string `json:"symbol"`
	AnnualEarnings []struct {
		FiscalDateEnding string `json:"fiscalDateEnding"`
		ReportedEPS      string `json:"reportedEPS"`
	} `json:"annualEarnings"`
	QuarterlyEarnings []struct {
		FiscalDateEnding   string `json:"fiscalDateEnding"`
		ReportedDate       string `json:"reportedDate"`
		ReportedEPS        string `json:"reportedEPS"`
		EstimatedEPS       string `json:"estimatedEPS"`
		Surprise           string `json:"surprise"`
		SurprisePercentage string `json:"surprisePercentage"`
	} `json:"quarterlyEarnings"`
}

type AlphavantageCashFlow struct {
	Symbol        string `json:"symbol"`
	AnnualReports []struct {
		FiscalDateEnding                                          string `json:"fiscalDateEnding"`
		ReportedCurrency                                          string `json:"reportedCurrency"`
		OperatingCashflow                                         string `json:"operatingCashflow"`
		PaymentsForOperatingActivities                            string `json:"paymentsForOperatingActivities"`
		ProceedsFromOperatingActivities                           string `json:"proceedsFromOperatingActivities"`
		ChangeInOperatingLiabilities                              string `json:"changeInOperatingLiabilities"`
		ChangeInOperatingAssets                                   string `json:"changeInOperatingAssets"`
		DepreciationDepletionAndAmortization                      string `json:"depreciationDepletionAndAmortization"`
		CapitalExpenditures                                       string `json:"capitalExpenditures"`
		ChangeInReceivables                                       string `json:"changeInReceivables"`
		ChangeInInventory                                         string `json:"changeInInventory"`
		ProfitLoss                                                string `json:"profitLoss"`
		CashflowFromInvestment                                    string `json:"cashflowFromInvestment"`
		CashflowFromFinancing                                     string `json:"cashflowFromFinancing"`
		ProceedsFromRepaymentsOfShortTermDebt                     string `json:"proceedsFromRepaymentsOfShortTermDebt"`
		PaymentsForRepurchaseOfCommonStock                        string `json:"paymentsForRepurchaseOfCommonStock"`
		PaymentsForRepurchaseOfEquity                             string `json:"paymentsForRepurchaseOfEquity"`
		PaymentsForRepurchaseOfPreferredStock                     string `json:"paymentsForRepurchaseOfPreferredStock"`
		DividendPayout                                            string `json:"dividendPayout"`
		DividendPayoutCommonStock                                 string `json:"dividendPayoutCommonStock"`
		DividendPayoutPreferredStock                              string `json:"dividendPayoutPreferredStock"`
		ProceedsFromIssuanceOfCommonStock                         string `json:"proceedsFromIssuanceOfCommonStock"`
		ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet string `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
		ProceedsFromIssuanceOfPreferredStock                      string `json:"proceedsFromIssuanceOfPreferredStock"`
		ProceedsFromRepurchaseOfEquity                            string `json:"proceedsFromRepurchaseOfEquity"`
		ProceedsFromSaleOfTreasuryStock                           string `json:"proceedsFromSaleOfTreasuryStock"`
		ChangeInCashAndCashEquivalents                            string `json:"changeInCashAndCashEquivalents"`
		ChangeInExchangeRate                                      string `json:"changeInExchangeRate"`
		NetIncome                                                 string `json:"netIncome"`
	} `json:"annualReports"`
	QuarterlyReports []struct {
		FiscalDateEnding                                          string `json:"fiscalDateEnding"`
		ReportedCurrency                                          string `json:"reportedCurrency"`
		OperatingCashflow                                         string `json:"operatingCashflow"`
		PaymentsForOperatingActivities                            string `json:"paymentsForOperatingActivities"`
		ProceedsFromOperatingActivities                           string `json:"proceedsFromOperatingActivities"`
		ChangeInOperatingLiabilities                              string `json:"changeInOperatingLiabilities"`
		ChangeInOperatingAssets                                   string `json:"changeInOperatingAssets"`
		DepreciationDepletionAndAmortization                      string `json:"depreciationDepletionAndAmortization"`
		CapitalExpenditures                                       string `json:"capitalExpenditures"`
		ChangeInReceivables                                       string `json:"changeInReceivables"`
		ChangeInInventory                                         string `json:"changeInInventory"`
		ProfitLoss                                                string `json:"profitLoss"`
		CashflowFromInvestment                                    string `json:"cashflowFromInvestment"`
		CashflowFromFinancing                                     string `json:"cashflowFromFinancing"`
		ProceedsFromRepaymentsOfShortTermDebt                     string `json:"proceedsFromRepaymentsOfShortTermDebt"`
		PaymentsForRepurchaseOfCommonStock                        string `json:"paymentsForRepurchaseOfCommonStock"`
		PaymentsForRepurchaseOfEquity                             string `json:"paymentsForRepurchaseOfEquity"`
		PaymentsForRepurchaseOfPreferredStock                     string `json:"paymentsForRepurchaseOfPreferredStock"`
		DividendPayout                                            string `json:"dividendPayout"`
		DividendPayoutCommonStock                                 string `json:"dividendPayoutCommonStock"`
		DividendPayoutPreferredStock                              string `json:"dividendPayoutPreferredStock"`
		ProceedsFromIssuanceOfCommonStock                         string `json:"proceedsFromIssuanceOfCommonStock"`
		ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet string `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
		ProceedsFromIssuanceOfPreferredStock                      string `json:"proceedsFromIssuanceOfPreferredStock"`
		ProceedsFromRepurchaseOfEquity                            string `json:"proceedsFromRepurchaseOfEquity"`
		ProceedsFromSaleOfTreasuryStock                           string `json:"proceedsFromSaleOfTreasuryStock"`
		ChangeInCashAndCashEquivalents                            string `json:"changeInCashAndCashEquivalents"`
		ChangeInExchangeRate                                      string `json:"changeInExchangeRate"`
		NetIncome                                                 string `json:"netIncome"`
	} `json:"quarterlyReports"`
}
