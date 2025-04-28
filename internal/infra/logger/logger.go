package logger

import (
	"github.com/sirupsen/logrus"
)

var Log = NewLogrusAdapter()

type LogLevel int

const (
	DebugLevel LogLevel = 0
	InfoLevel  LogLevel = 1
	WarnLevel  LogLevel = 2
	ErrorLevel LogLevel = 3
	FatalLevel LogLevel = 4
)

// Logger is an interface for logging.
type Logger interface {
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
	SetLevel(level LogLevel)
}

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
func (l *LogrusAdapter) SetLevel(level LogLevel) {
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
