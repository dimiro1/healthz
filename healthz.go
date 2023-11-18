package healthz

import (
	"fmt"
	"net/http"
)

// AlwaysUp is an HTTP handler that always responds with an HTTP 200 status code,
// indicating that the application is healthy and operational. It is a convenience
// wrapper around the Handler function, pre-configured to always return a healthy status.
// This can be useful in scenarios where a simple, always-positive health check is needed,
// such as in development environments or for applications with no specific health criteria.
//
// The function internally uses Handler, passing a function that always returns true,
// to create a health check endpoint that invariably reports the application as up and running.
//
// Example usage:
//
//	http.HandleFunc("/healthz", healthz.AlwaysUp)
func AlwaysUp(w http.ResponseWriter, r *http.Request) {
	Check(func() bool { return true }).ServeHTTP(w, r)
}

// Check creates an HTTP handler for health check endpoints. It uses okFn,
// a custom function provided by the user, to determine the application's health.
// okFn should return true if the application is healthy, false otherwise.
//
// The handler responds with HTTP 200 (OK) if okFn returns true, indicating
// a healthy state, and with HTTP 503 (Service Unavailable) if okFn returns false,
// indicating an unhealthy state. The response includes an HTML body with a
// background color: green for healthy (200) and red for unhealthy (503).
//
// Example usage:
//
//	http.Handle("/healthz", Check(func() bool {
//	  // Custom health check logic
//	  return true // or false based on health status
//	}))
//
// Parameters:
//   - okFn: Function returning a bool to indicate health status.
//
// Returns:
//
//	An http.Handler for the health check endpoint.
func Check(okFn func() bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			color  = "green"
			status = http.StatusOK
		)

		if !okFn() {
			color = "red"
			status = http.StatusServiceUnavailable
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(status)

		_, _ = w.Write([]byte(fmt.Sprintf(`<!DOCTYPE html><html><body style="background-color: %s"></body></html>`, color)))
	})
}
