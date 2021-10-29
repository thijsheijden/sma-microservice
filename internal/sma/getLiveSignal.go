package sma

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// getLiveSignal gets the live signal price for the given exchange and market
func getLiveSignal(request *Request) (signal float64, err error) {

	// Create request URL with the given exchange
	requestURL := fmt.Sprintf("http://cryptohopper-ticker-frontend.us-east-1.elasticbeanstalk.com/v1/%s/ticker", request.exchange)

	// Make the request
	response, err := http.Get(requestURL)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = errors.New("error while getting signal price")
		return
	}

	var signals liveSignals
	err = json.NewDecoder(response.Body).Decode(&signals)
	if err != nil {
		return
	}

	// Grab the signal we are interested in
	if s, ok := signals.Data[request.market]; ok {
		signal, err = strconv.ParseFloat(s.Last, 64)
	} else {
		err = errors.New("last price not float")
		return
	}

	return
}
