package sma

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Request describes the request to get an SMA value
type Request struct {
	exchange       string
	market         string
	intervalString string
	interval       time.Duration
}

// Indicator indicates whether to BUY, SELL or NEUTRAL
type Indicator int

const (
	Buy Indicator = iota
	Sell
	Neutral
)

func (i Indicator) String() string {
	switch i {
	case 0:
		return "Buy"
	case 1:
		return "Sell"
	case 2:
		return "Neutral"
	default:
		return "Neutral"
	}
}

// CreateSMARequest creates an SMA request from the URL params
func CreateSMARequest(exchange string, market string, interval string) (request *Request, err error) {

	request = &Request{
		exchange:       exchange,
		market:         market,
		intervalString: interval,
	}

	// Check if the interval contains the day unit, as Go cannot parse this
	if strings.ContainsRune(interval, 'd') {
		// Remove the d from the string
		numberOfDays, convertErr := strconv.Atoi(interval[:len(interval)-1])
		if convertErr != nil {
			err = errors.New("invalid interval format")
			return
		}

		// Change into hours format
		interval = fmt.Sprintf("%dh", numberOfDays*24)
	}

	// Parse the interval string into the Go time.Duration type
	_interval, err := time.ParseDuration(interval)
	if err != nil {
		return nil, err
	}

	request.interval = _interval

	return
}
