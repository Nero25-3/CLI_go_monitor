package internal

// MockLogger is a mock implementation of LoggerInterface for testing purposes.
type MockLogger struct {
	Messages []string
}

// Info logs an informational message.
func (m *MockLogger) Info(msg string) {
	m.Messages = append(m.Messages, "INFO: "+msg)
}

// Warn logs a warning message.
func (m *MockLogger) Warn(msg string) {
	m.Messages = append(m.Messages, "WARN: "+msg)
}

// Error logs an error message.
func (m *MockLogger) Error(msg string) {
	m.Messages = append(m.Messages, "ERROR: "+msg)
}

// Reset clears all logged messages.
func (m *MockLogger) Reset() {
	m.Messages = []string{}
}
