package logger

import (
	"books/internal/config/env"

	"github.com/sirupsen/logrus"
)

var Log = NewLogrusAdapter()

// Logger is an interface for logging.
type Logger interface {
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
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

func (l *LogrusAdapter) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *LogrusAdapter) Warn(args ...any) {
	l.logger.Warn(args...)
}

func (l *LogrusAdapter) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *LogrusAdapter) Fatal(args ...any) {
	l.logger.Fatal(args...)
}
