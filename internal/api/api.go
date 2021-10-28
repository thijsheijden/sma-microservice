package api

import "net/http"

// API describes the API object
type API struct {
}

// New creates a new API object
func New() *API {
	return &API{}
}

// Start starts the API
func (a *API) Start() {
	http.HandleFunc("/sma", a.smaHandler().ServeHTTP)
	http.ListenAndServe(":3000", nil)
}
