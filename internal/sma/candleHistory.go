package sma

type candleHistory []candle

type candle struct {
	Close float64 `json:"Close"`
}
