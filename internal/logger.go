package internal

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger wraps a standard logger with log rotation capabilities.
type Logger struct {
	logger *log.Logger
}

// NewLogger creates a new Logger that writes to the specified file path with log rotation.
func NewLogger(path string) *Logger {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    5, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}
	logger := log.New(lumberjackLogger, "", log.LstdFlags)
	return &Logger{logger: logger}
}

// Info logs an informational message.
func (l *Logger) Info(msg string) {
	l.logger.Println(msg)
}
