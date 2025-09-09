package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// mock logger that also implements io.Writer
type mockLogger struct {
	buf bytes.Buffer
}

// implements io.Writer
func (m *mockLogger) Write(p []byte) (int, error) {
	return m.buf.Write(p)
}

// implements our Log method
func (m *mockLogger) Log(entry LogEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	m.buf.Write(append(data, '\n'))
	return nil
}

func TestHelloHandler(t *testing.T) {
	_ = os.Setenv("HOSTNAME", "test-pod")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "Hello Sami") {
		t.Errorf("expected greeting, got %s", body)
	}
	if !strings.Contains(string(body), "test-pod") {
		t.Errorf("expected pod name, got %s", body)
	}
}

func TestLoggingMiddleware(t *testing.T) {
	logger := &mockLogger{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok"))
	})

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	rr := httptest.NewRecorder()

	// wrap handler with middleware
	middleware := loggingMiddleware(handler, &dualLogger{file: logger})
	middleware.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", rr.Code)
	}

	lines := strings.Split(strings.TrimSpace(logger.buf.String()), "\n")
	if len(lines) == 0 {
		t.Fatal("expected log entry, got none")
	}

	var entry LogEntry
	if err := json.Unmarshal([]byte(lines[0]), &entry); err != nil {
		t.Fatalf("failed to parse log entry: %v", err)
	}

	if entry.Method != http.MethodPost {
		t.Errorf("expected method POST, got %s", entry.Method)
	}
	if entry.Path != "/test" {
		t.Errorf("expected path /test, got %s", entry.Path)
	}
	if entry.Status != http.StatusCreated {
		t.Errorf("expected status 201, got %d", entry.Status)
	}
	if entry.ResponseSize == 0 {
		t.Error("expected non-zero response size")
	}
}

// Add Ability to Bemnchmark with logging
// Test with:
// go test -bench . -benchmem
func BenchmarkHelloHandlerWithLogging(b *testing.B) {
	logger := &mockLogger{}

	handler := loggingMiddleware(http.HandlerFunc(helloHandler), &dualLogger{file: logger})

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
	}
}
