package main

import (
	"fmt"
	"net/http"
	"time"
)

// ResponseWriter wrapper to capture status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Your original handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Hello World - SAM - %s\n", currentTime)
}

// Middleware to log request info, status, and latency
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap ResponseWriter to capture status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}

		// Call the next handler
		next.ServeHTTP(lrw, r)

		duration := time.Since(start)
		currentTime := start.Format("2006-01-02 15:04:05")

		fmt.Printf("[%s] %s %s %s %d %s\n",
			currentTime,
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			lrw.statusCode,
			duration,
		)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	// Wrap all requests with logging middleware
	loggedMux := loggingMiddleware(mux)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
