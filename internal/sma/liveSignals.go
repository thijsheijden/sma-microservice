package sma

type liveSignals struct {
	Exchange string            `json:"exchange"`
	Data     map[string]signal `json:"data"`
	Time     int64             `json:"time"`
}

type signal struct {
	CurrencyPair  string `json:"currencyPair"`
	Last          string `json:"last"`
	LowestAsk     string `json:"lowestAsk"`
	HighestBid    string `json:"highestBid"`
	PercentChange string `json:"percentChange"`
	BaseVolume    string `json:"baseVolume"`
	QuoteVolume   string `json:"quoteVolume"`
	IsFrozen      string `json:"isFrozen"`
	The24HrHigh   string `json:"24hrHigh"`
	The24HrLow    string `json:"24hrLow"`
}
