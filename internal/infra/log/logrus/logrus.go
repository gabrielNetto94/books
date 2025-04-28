package logrus

import (
	"books/internal/infra/log"

	"github.com/sirupsen/logrus"
)

// LogrusAdapter is an adapter that implements the Logger interface using Logrus.
type LogrusAdapter struct {
	logger *logrus.Logger
}

// NewLogrusAdapter creates a new instance of LogrusAdapter.
func NewLogrusAdapter() *LogrusAdapter {

	return &LogrusAdapter{logrus.New()}
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
func (l *LogrusAdapter) SetLevel(level log.LogLevel) {
	switch level {
	case 0:
		l.logger.SetLevel(logrus.DebugLevel)
	case 1:
		l.logger.SetLevel(logrus.InfoLevel)
	case 2:
		l.logger.SetLevel(logrus.WarnLevel)
	case 3:
		l.logger.SetLevel(logrus.ErrorLevel)
	case 4:
		l.logger.SetLevel(logrus.FatalLevel)
	default:
		l.logger.SetLevel(logrus.InfoLevel)
	}
}
