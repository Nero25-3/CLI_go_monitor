package core_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"CLI_go_monitor/internal/core"
)

type mockLogger struct {
	Messages []string
}

func (m *mockLogger) Info(msg string)  { m.Messages = append(m.Messages, msg) }
func (m *mockLogger) Warn(msg string)  { m.Messages = append(m.Messages, msg) }
func (m *mockLogger) Error(msg string) { m.Messages = append(m.Messages, msg) }

func TestRunCheck(t *testing.T) {
	original := core.MonitorURLFunc
	defer func() { core.MonitorURLFunc = original }()

	core.MonitorURLFunc = func(url string, timeout time.Duration) error {
		if strings.Contains(url, "fail") {
			return fmt.Errorf("fail error")
		}
		return nil
	}

	err := core.RunCheck([]string{"https://ok.com", "https://fail.com"}, 5, "", "", "")
	if err != nil {
		t.Fatalf("RunCheck failed: %v", err)
	}

}
