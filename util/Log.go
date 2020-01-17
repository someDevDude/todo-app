package util

import (
	"os"

	"github.com/op/go-logging"
)

var log *logging.Logger
var logFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// InitialiseLogger creates the logger
func InitialiseLogger() {
	log = logging.MustGetLogger("todo-server")
	loggingBackend := logging.NewLogBackend(os.Stderr, "", 0)
	loggingBackendFormatter := logging.NewBackendFormatter(loggingBackend, logFormat)
	logging.SetBackend(loggingBackend, loggingBackendFormatter)

	Infof("Logger initialised")
}

// Debug logs message with DEBUG level
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debugf logs message with DEBUG level
func Debugf(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

// Info logs message with INFO level
func Info(args ...interface{}) {
	log.Debug(args...)
}

// Infof logs message with INFO level
func Infof(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

// Notice logs message with NOTICE level
func Notice(args ...interface{}) {
	log.Debug(args...)
}

// Noticef logs message with NOTICE level
func Noticef(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

// Warning logs message with WARNING level
func Warning(args ...interface{}) {
	log.Debug(args...)
}

// Warningf logs message with WARNING level
func Warningf(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

// Error logs message with ERROR level
func Error(args ...interface{}) {
	log.Debug(args...)
}

// Errorf logs message with ERROR level
func Errorf(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

// Critical logs message with CRITICAL level
func Critical(args ...interface{}) {
	log.Debug(args...)
}

// Criticalf logs message with CRITICAL level
func Criticalf(message string, args ...interface{}) {
	log.Debugf(message, args...)
}
