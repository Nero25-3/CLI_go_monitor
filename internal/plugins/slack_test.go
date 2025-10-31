package plugins_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"CLI_go_monitor/internal/plugins"

	"github.com/stretchr/testify/assert"
)

func TestSlackNotifier_Name(t *testing.T) {
	notifier := plugins.NewSlackNotifier("https://hooks.slack.com/test")
	assert.Equal(t, "Slack", notifier.Name())
}

func TestSlackNotifier_Notify_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	notifier := plugins.NewSlackNotifier(server.URL)
	result := plugins.Notice{
		URL:    "https://example.com",
		Status: "FAILED",
	}

	err := notifier.Notify(result)
	assert.NoError(t, err)
}
