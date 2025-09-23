package internal

import (
	"log"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Level represents the severity of the log message.
type Level int

// constants for log levels
const (
	INFO Level = iota
	WARN
	ERROR
)

// Logger is a simple logging structure with log rotation.
type Logger struct {
	mu      sync.Mutex
	Logger  *log.Logger
	Level   Level
	Enabled bool
}

// LoggerInterface defines the methods for logging at different levels.
type LoggerInterface interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

// NewLogger creates a new Logger instance.
func NewLogger(path string, level Level, enabled bool) *Logger {
	var logger *log.Logger
	if enabled {
		lumberjackLogger := &lumberjack.Logger{
			Filename:   path,
			MaxSize:    5, // MB
			MaxBackups: 3,
			MaxAge:     28, // days
			Compress:   true,
		}
		logger = log.New(lumberjackLogger, "", log.LstdFlags)
	} else {
		// If disabled, just discard logs or log to os.Stdout as fallback
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}
	return &Logger{
		Logger:  logger,
		Level:   level,
		Enabled: enabled,
	}
}

func (l *Logger) log(lv Level, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.Enabled {
		return
	}
	if lv >= l.Level {
		l.Logger.Println(msg)
	}
}

// Info logs an info level message.
func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

// Warn logs a warning level message.
func (l *Logger) Warn(msg string) {
	l.log(WARN, msg)
}

// Error logs an error level message.
func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}
