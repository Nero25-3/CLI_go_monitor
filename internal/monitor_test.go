package internal_test

import (
	"CLI_go_monitor/internal"
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
