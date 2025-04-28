package log

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
