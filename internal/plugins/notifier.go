package plugins

import "time"

// Notice represents the result of a monitoring check
type Notice struct {
	URL        string
	Status     string
	StatusCode int
	Message    string
	Timestamp  time.Time
	Error      error
}

// Notifier defines the interface for notification plugins
// Any alert system (Slack, Email, Telegram, Webhook) must implement this interface
type Notifier interface {
	// Notify sends an alert based on the check result
	Notify(result Notice) error

	// Name returns the notifier's identifier
	Name() string
}
