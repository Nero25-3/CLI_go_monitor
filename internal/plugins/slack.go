package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SlackNotifier sends alerts to Slack via webhook
type SlackNotifier struct {
	WebhookURL string
	Client     *http.Client
}

// NewSlackNotifier creates a new Slack notifier instance
func NewSlackNotifier(webhookURL string) *SlackNotifier {
	return &SlackNotifier{
		WebhookURL: webhookURL,
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// Name returns the notifier identifier
func (s *SlackNotifier) Name() string {
	return "Slack"
}

// Notify sends a message to Slack webhook
func (s *SlackNotifier) Notify(result Notice) error {
	if s.WebhookURL == "" {
		return fmt.Errorf("slack webhook URL is empty")
	}

	// Build Slack message payload
	payload := map[string]interface{}{
		"text": "CLI_go_monitor Alert",
		"blocks": []map[string]interface{}{
			{
				"type": "section",
				"text": map[string]interface{}{
					"type": "mrkdwn",
					"text": fmt.Sprintf(
						"*Status:* %s\n*URL:* %s\n*Message:* %s\n*Time:* %s",
						result.Status,
						result.URL,
						result.Message,
						result.Timestamp.Format("2006-01-02 15:04:05"),
					),
				},
			},
		},
	}

	// Marshal to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal Slack payload: %w", err)
	}

	// Send POST request to webhook
	resp, err := s.Client.Post(
		s.WebhookURL,
		"application/json",
		bytes.NewBuffer(payloadBytes),
	)
	if err != nil {
		return fmt.Errorf("failed to send Slack notification: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("slack webhook returned status %d", resp.StatusCode)
	}

	return nil
}
