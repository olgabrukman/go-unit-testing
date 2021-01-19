package example5_httptest

import (
	"io"
	"net/http"
)

// e.g. http.HandleFunc("/health-check", HealthCheckHandler)
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
