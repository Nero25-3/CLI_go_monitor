package internal

import (
	"fmt"
	"net/http"
	"time"
)

// MonitorURL checks if the given URL is reachable within the specified timeout.
func MonitorURL(url string, timeout time.Duration) error {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		return nil
	}
	return fmt.Errorf("HTTP status %d", resp.StatusCode)
}
