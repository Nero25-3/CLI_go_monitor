package cmd_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"CLI_go_monitor/cmd"
)

type mockLogger struct {
	Messages []string
}

func (m *mockLogger) Info(msg string)  { m.Messages = append(m.Messages, msg) }
func (m *mockLogger) Warn(msg string)  { m.Messages = append(m.Messages, msg) }
func (m *mockLogger) Error(msg string) { m.Messages = append(m.Messages, msg) }

func TestRunCheck(t *testing.T) {
	original := cmd.MonitorURLFunc
	defer func() { cmd.MonitorURLFunc = original }()

	cmd.MonitorURLFunc = func(url string, timeout time.Duration) error {
		if strings.Contains(url, "fail") {
			return fmt.Errorf("fail error")
		}
		return nil
	}

	err := cmd.RunCheck([]string{"https://ok.com", "https://fail.com"}, 5, "", "", "")
	if err != nil {
		t.Fatalf("RunCheck failed: %v", err)
	}

}
