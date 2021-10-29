package sma

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetIndicator gets the SMA with the params specified in the request
func GetIndicator(request *Request) (indicator Indicator, err error) {
	// Get the live signal price
	signal, err := getLiveSignal(request)
	if err != nil {
		// This error could be ignored, and instead the last close price could be used
		return
	}

	// Get start and end times in unix format
	endTime := time.Now().Unix()

	// Get the start time by taking the current time and removing 55 intervals from it
	startTime := time.Now().Add(-55 * request.interval).Unix()

	// Create the request url
	requestURL := fmt.Sprintf("http://cryptohopper-ticker-frontend.us-east-1.elasticbeanstalk.com/v1/%s/candles?pair=%s&start=%d&end=%d&period=%s",
		request.exchange,
		request.market,
		startTime,
		endTime,
		request.intervalString)

	// Make the request
	response, err := http.Get(requestURL)
	if err != nil {
		return
	}

	// Read the body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var candles candleHistory
	err = json.Unmarshal(body, &candles)
	if err != nil {
		return
	}

	// Make sure we have at least 55 candles
	if len(candles) < 55 {
		err = errors.New("not enough candles to calculate SMA indicator")
		return
	}

	// Replace the last candle with a new candle containing the signal price as close value
	candles = candles[:len(candles)-1]
	candles = append(candles, candle{Close: signal})

	// Get the SMA(8)
	var sum float64
	for i := len(candles) - 1; i > len(candles)-9; i-- {
		sum += candles[i].Close
	}
	sma8 := sum / 8

	// Get the SMA(55)
	for i := len(candles) - 9; i > len(candles)-56; i-- {
		sum += candles[i].Close
	}
	sma55 := sum / 55

	// Grab the previous close value from the candle before the SMA(8)
	prevValue := candles[len(candles)-9].Close

	if sma8 < sma55 && prevValue >= sma8 {
		indicator = Sell
	} else if sma8 > sma55 && prevValue <= sma8 {
		indicator = Buy
	} else {
		indicator = Neutral
	}

	return
}
