package datasource

type Gurufocus struct {
	API_KEY string
}

func InitGurufocus(apikey string) *Gurufocus {
	return &Gurufocus{
		API_KEY: apikey,
	}
}
