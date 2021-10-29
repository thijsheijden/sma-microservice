package sma

type candleHistory []candles

type candles struct {
	Open        float64 `json:"Open"`
	High        float64 `json:"High"`
	Low         float64 `json:"Low"`
	Close       float64 `json:"Close"`
	BaseVolume  float64 `json:"BaseVolume"`
	QuoteVolume float64 `json:"QuoteVolume"`
	OpenTime    string  `json:"OpenTime"`
}
