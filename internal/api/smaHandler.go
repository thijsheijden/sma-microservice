package api

import (
	"errors"
	"net/http"
	"net/url"
	"sma-microservice/internal/sma"
)

// smaHandler returns an HTTP handler for the 'simple moving average'
func (a *API) smaHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Grab the required URL parameters
		exchange, interval, market, err := getSMAURLParams(r.URL.Query())
		if err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}

		// Create the SMA Request
		smaRequest, err := sma.CreateSMARequest(exchange, market, interval)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		// Get the indicator
		indicator, err := sma.GetIndicator(smaRequest)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		w.Write([]byte(indicator.String()))
	})
}

// Get the URL parameters from the request url.Values
func getSMAURLParams(values url.Values) (exchange string, interval string, market string, err error) {
	if _exchange, ok := values["exchange"]; ok {
		exchange = _exchange[0]
	} else {
		err = errors.New("missing URL param exchange")
		return
	}

	if _interval, ok := values["interval"]; ok {
		interval = _interval[0]
	} else {
		err = errors.New("missing URL param interval")
		return
	}

	if _market, ok := values["market"]; ok {
		market = _market[0]
	} else {
		err = errors.New("missing URL param market")
		return
	}

	return
}
