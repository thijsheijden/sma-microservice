package sma

type candleHistory []candle

type candle struct {
	Close float64 `json:"Close"` // Only store the Close value as that is the only value we need
}
