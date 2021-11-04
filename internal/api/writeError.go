package api

import "net/http"

// writeError writes an error to the given http.ResponseWriter
func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
	return
}
