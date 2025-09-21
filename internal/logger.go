package internal

import (
	"log"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Level int

const (
	INFO Level = iota
	WARN
	ERROR
)

type Logger struct {
	mu      sync.Mutex
	logger  *log.Logger
	level   Level
	enabled bool
}

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
		logger:  logger,
		level:   level,
		enabled: enabled,
	}
}

func (l *Logger) log(lv Level, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.enabled {
		return
	}
	if lv >= l.level {
		l.logger.Println(msg)
	}
}

func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

func (l *Logger) Warn(msg string) {
	l.log(WARN, msg)
}

func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}
