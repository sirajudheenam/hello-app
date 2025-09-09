package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// ResponseWriter wrapper to capture status and size
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.size += size
	return size, err
}

// LogEntry represents a single request log
type LogEntry struct {
	Timestamp    string `json:"timestamp"`
	RemoteAddr   string `json:"remote_addr"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	Status       int    `json:"status"`
	LatencyMs    int64  `json:"latency_ms"`
	ResponseSize int    `json:"response_size"`
}

// Logger that writes to both stdout and file
type dualLogger struct {
	file io.Writer
}

func newDualLogger(filename string) *dualLogger {
	file := &lumberjack.Logger{
		Filename:   os.Getenv("LOG_FILE"),
		MaxSize:    20, // MB
		MaxBackups: 7,
		MaxAge:     30,   // days
		Compress:   true, // compress old logs
	}
	return &dualLogger{file: io.MultiWriter(file, os.Stdout)}
}

func (l *dualLogger) Log(entry LogEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	_, err = l.file.Write(append(data, '\n'))
	return err
}

// Logging middleware
func loggingMiddleware(next http.Handler, logger *dualLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}

		next.ServeHTTP(lrw, r)

		duration := time.Since(start)

		entry := LogEntry{
			Timestamp:    start.Format(time.RFC3339),
			RemoteAddr:   r.RemoteAddr,
			Method:       r.Method,
			Path:         r.URL.Path,
			Status:       lrw.statusCode,
			LatencyMs:    duration.Milliseconds(),
			ResponseSize: lrw.size,
		}

		_ = logger.Log(entry)
	})
}

// Your handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	podName := os.Getenv("HOSTNAME") // Kubernetes pod name
	fmt.Fprintf(w, "Hello World - SAM - %s - Pod: %s\n", currentTime, podName)

}

func main() {
	logger := newDualLogger("logs/server.json.log")

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	loggedMux := loggingMiddleware(mux, logger)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
