package sma

type liveSignals struct {
	Data map[string]signal `json:"data"`
}

type signal struct {
	Last string `json:"last"`
}
