package internal_test

import (
	"CLI_go_monitor/internal"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	const yamlContent = `
urls:
  - https://google.com
  - https://example.com
interval: 30
timeout: 5
logfile: monitor.log
`
	tmpfile, err := os.CreateTemp("", "config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(yamlContent)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	cfg, err := internal.ReadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("ReadConfig failed: %v", err)
	}
	if len(cfg.URLs) != 2 {
		t.Errorf("expected 2 URLs, got %d", len(cfg.URLs))
	}
	if cfg.Interval != 30 {
		t.Errorf("expected interval 30, got %d", cfg.Interval)
	}
	if cfg.Timeout != 5 {
		t.Errorf("expected timeout 5, got %d", cfg.Timeout)
	}
	if cfg.Logfile != "monitor.log" {
		t.Errorf("expected logfile monitor.log, got %s", cfg.Logfile)
	}
}
