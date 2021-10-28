package sma

import "time"

type SMARequest struct {
	exchange        string
	market          string
	interval        time.Duration
	numberOfCandles int
}

// CreateSMARequest creates an SMA request from the URL params
func CreateSMARequest(exchange string, market string, interval string, numberOfCandles int) (*SMARequest, error) {
	// Parse the interval string into the Go time.Duration type
	_interval, err := time.ParseDuration(interval)
	if err != nil {
		return nil, err
	}

	return &SMARequest{
		exchange:        exchange,
		market:          market,
		interval:        _interval,
		numberOfCandles: numberOfCandles,
	}, nil
}
