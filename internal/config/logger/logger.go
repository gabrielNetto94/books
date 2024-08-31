package logger

import (
	"books/internal/config/env"

	"github.com/sirupsen/logrus"
)

var Log = NewLogrusAdapter()

// Logger is an interface for logging.
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

// LogrusAdapter is an adapter that implements the Logger interface using Logrus.
type LogrusAdapter struct {
	logger *logrus.Logger
}

// NewLogrusAdapter creates a new instance of LogrusAdapter.
func NewLogrusAdapter() *LogrusAdapter {

	l := logrus.New()

	if env.IsProduction() {
		l.SetLevel(logrus.ErrorLevel)
	}
	l.SetLevel(logrus.InfoLevel)

	return &LogrusAdapter{logger: l}
}

func (l *LogrusAdapter) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogrusAdapter) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *LogrusAdapter) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogrusAdapter) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}
