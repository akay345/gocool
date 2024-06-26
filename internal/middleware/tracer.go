
package middleware

import (
    "log"
    "net/http"
    "time"
)

// TracerMiddleware logs the incoming HTTP request and measures its duration.
func TracerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Log the start of the request
        log.Printf("Started %s %s", r.Method, r.RequestURI)

        // Use a response writer to capture HTTP status code
        rec := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}

        // Serve the next handler
        next.ServeHTTP(rec, r)

        // Log the completion of the request
        duration := time.Since(start)
        log.Printf("Completed %s %s in %v with status %d", r.Method, r.RequestURI, duration, rec.statusCode)
    })
}

// statusRecorder is a wrapper around http.ResponseWriter that allows us to capture the status code
type statusRecorder struct {
    http.ResponseWriter
    statusCode int
}

// WriteHeader captures the status code and calls the original WriteHeader function
func (rec *statusRecorder) WriteHeader(code int) {
    rec.statusCode = code
    rec.ResponseWriter.WriteHeader(code)
}
