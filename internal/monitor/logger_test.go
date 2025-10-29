package monitor_test

import (
	"CLI_go_monitor/internal/monitor"
	"bytes"
	"log"
	"strings"
	"testing"
)

func NewTestLogger(buf *bytes.Buffer) *log.Logger {
	return log.New(buf, "", 0)
}

func TestLoggerLevels(t *testing.T) {
	var buf bytes.Buffer

	logger := &monitor.Logger{
		Logger:  NewTestLogger(&buf),
		Level:   monitor.INFO,
		Enabled: true,
	}

	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")

	output := buf.String()
	if !strings.Contains(output, "info message") ||
		!strings.Contains(output, "warn message") ||
		!strings.Contains(output, "error message") {
		t.Errorf("expected all log levels printed, got %q", output)
	}
}

func TestLoggerLevelFiltering(t *testing.T) {
	var buf bytes.Buffer

	logger := &monitor.Logger{
		Logger:  NewTestLogger(&buf),
		Level:   monitor.WARN,
		Enabled: true,
	}

	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")

	output := buf.String()
	if strings.Contains(output, "info message") {
		t.Errorf("info message should be filtered out at WARN level")
	}
	if !strings.Contains(output, "warn message") || !strings.Contains(output, "error message") {
		t.Errorf("warn and error messages should appear, got %q", output)
	}
}

func TestLoggerDisabled(t *testing.T) {
	var buf bytes.Buffer

	logger := &monitor.Logger{
		Logger:  NewTestLogger(&buf),
		Level:   monitor.INFO,
		Enabled: false,
	}

	logger.Info("info message")
	logger.Warn("warn message")

	output := buf.String()
	if output != "" {
		t.Errorf("expected no output when logger disabled, got %q", output)
	}
}

func TestWithMockLogger(t *testing.T) {
	mock := &monitor.MockLogger{}

	mock.Info("test info")
	mock.Warn("test warn")
	mock.Error("test error")

	if len(mock.Messages) != 3 {
		t.Errorf("expected 3 messages, got %d", len(mock.Messages))
	}
}

func TestNewLogger(t *testing.T) {
	logger := monitor.NewLogger("test.log", monitor.INFO, true)
	if logger == nil {
		t.Fatal("expected non-nil logger")
	}
	if logger.Enabled != true {
		t.Errorf("expected Enabled true")
	}
	if logger.Level != monitor.INFO {
		t.Errorf("expected Level INFO")
	}

	logger2 := monitor.NewLogger("", monitor.ERROR, false)
	if logger2 == nil {
		t.Fatal("expected non-nil logger")
	}
	if logger2.Enabled != false {
		t.Errorf("expected Enabled false")
	}
	if logger2.Level != monitor.ERROR {
		t.Errorf("expected Level ERROR")
	}
}

func TestMockLoggerReset(t *testing.T) {
	mockLog := &monitor.MockLogger{}
	mockLog.Info("message 1")
	mockLog.Warn("message 2")

	if len(mockLog.Messages) == 0 {
		t.Fatal("expected messages")
	}

	mockLog.Reset()

	if len(mockLog.Messages) != 0 {
		t.Errorf("expected messages to be cleared after reset")
	}
}
