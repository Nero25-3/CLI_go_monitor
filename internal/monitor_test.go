package internal_test

import (
	"CLI_go_monitor/internal"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMonitorURL(t *testing.T) {
	tests := []struct {
		url     string
		timeout time.Duration
	}{
		{"https://www.google.com", 5 * time.Second},
		{"https://www.nonexistenturl123456.com", 5 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			errTest := internal.MonitorURL(tt.url, tt.timeout)
			// Check if the error is nil for reachable URLs and not nil for unreachable URLs
			if tt.url == "https://www.google.com" {
				assert.NoError(t, errTest)
			} else {
				assert.Error(t, errTest)
				if errTest != nil {
					t.Logf("Expected error for unreachable URL: %v", errTest)
				}
			}
		})
	}
}

func TestMonitorURLServer(t *testing.T) {
	// Server that returns 200 OK
	okServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer okServer.Close()

	// Server that returns 500 Internal Server Error
	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer errorServer.Close()

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{"OK status", okServer.URL, false},
		{"Internal Server Error", errorServer.URL, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := internal.MonitorURL(tt.url, 2*time.Second)
			if (err != nil) != tt.wantErr {
				t.Errorf("MonitorURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMonitorMultipleURLsConcurrently(t *testing.T) {
	okServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer okServer.Close()

	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer errorServer.Close()

	timeoutURL := "http://10.255.255.1" // Should timeout
	urls := []struct {
		url     string
		wantErr bool
	}{
		{okServer.URL, false},
		{errorServer.URL, true},
		{timeoutURL, true},
	}

	var wg sync.WaitGroup
	results := make([]error, len(urls))

	for i, tt := range urls {
		wg.Add(1)
		go func(i int, tt struct {
			url     string
			wantErr bool
		}) {
			defer wg.Done()
			results[i] = internal.MonitorURL(tt.url, 500*time.Millisecond)
		}(i, tt)
	}

	wg.Wait()

	// Assert results
	for i, tt := range urls {
		gotErr := results[i] != nil
		if gotErr != tt.wantErr {
			t.Errorf("MonitorURL(%q) error: %v, wantErr: %v", tt.url, gotErr, tt.wantErr)
		}
	}
}
