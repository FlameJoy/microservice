package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelMap = map[int]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARN",
	3: "ERROR",
	4: "FATAL",
}

type CustomLogger struct {
	logger *log.Logger
	level  int
	mu     sync.Mutex
}

func NewCustomLogger(logger *log.Logger, level int) *CustomLogger {
	return &CustomLogger{
		logger: logger,
		level:  level,
	}
}

func (l *CustomLogger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *CustomLogger) Debug(format string, v ...interface{}) {
	l.logMessage(DEBUG, format, v...)
}

func (l *CustomLogger) Info(format string, v ...interface{}) {
	l.logMessage(INFO, format, v...)
}

func (l *CustomLogger) Warn(format string, v ...interface{}) {
	l.logMessage(WARN, format, v...)
}

func (l *CustomLogger) Error(format string, v ...interface{}) {
	l.logMessage(ERROR, format, v...)
}

func (l *CustomLogger) Fatal(format string, v ...interface{}) {
	l.logMessage(FATAL, format, v...)
	os.Exit(1)
}

func (l *CustomLogger) logMessage(level int, format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if level < l.level {
		return
	}

	prefix := fmt.Sprintf("[%s] | %v ", levelMap[level], time.Now().Format("2006-01-02 15:04:05"))
	msg := fmt.Sprintf(format, v...)
	l.logger.Println(prefix, msg)
}
