package adaptlog

import "errors"

// LogLevel type
type LogLevel uint8

// Logging levels.
const (
	// Panic
	Panic LogLevel = iota
	// Fatal
	Fatal
	// Error
	Error
	// Warn
	Warn
	// Info
	Info
	// Debug
	Debug
)

// LevelLogger interface
type LevelLogger interface {
	ExtendedLogger
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})

	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})

	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})
}

var leveledLogger LevelLogger
var initialized bool

func init() {
	initialized = false
}

// InitializeLeveledLogger initializes a leveled logger. Only once needed
func InitializeLeveledLogger(logger LevelLogger) {
	leveledLogger = logger
	initialized = true
}

// NewLeveledLogger creates a new leveled logger
func NewLeveledLogger() (*LeveledLogger, error) {

	if !initialized {
		return nil, errors.New("text string")
	}

	return &LeveledLogger{
		logger: leveledLogger,
	}, nil
}

// LeveledLogger for logging with level support
type LeveledLogger struct {
	logger LevelLogger
}

// Panic logs level
func (l *LeveledLogger) Panic(args ...interface{}) {
	l.logger.Panic(args)
}

// Panicf logs level with message
func (l *LeveledLogger) Panicf(msg string, args ...interface{}) {
	l.logger.Panicf(msg, args)
}

// Panicln logs level with new line
func (l *LeveledLogger) Panicln(args ...interface{}) {
	l.logger.Panicln(args)
}

// Fatal logs level
func (l *LeveledLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args)
}

// Fatalf logs level with message
func (l *LeveledLogger) Fatalf(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args)
}

// Fatalln logs level with new line
func (l *LeveledLogger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args)
}

// Error logs level
func (l *LeveledLogger) Error(args ...interface{}) {
	l.logger.Error(args)
}

// Errorf logs level with message
func (l *LeveledLogger) Errorf(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args)
}

// Errorln logs level with new line
func (l *LeveledLogger) Errorln(args ...interface{}) {
	l.logger.Errorln(args)
}

// Warn logs level
func (l *LeveledLogger) Warn(args ...interface{}) {
	l.logger.Warn(args)
}

// Warnf logs level with message
func (l *LeveledLogger) Warnf(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args)
}

// Warnln logs level with new line
func (l *LeveledLogger) Warnln(args ...interface{}) {
	l.logger.Warnln(args)
}

// Info logs level
func (l *LeveledLogger) Info(args ...interface{}) {
	l.logger.Info(args)
}

// Infof logs level with message
func (l *LeveledLogger) Infof(msg string, args ...interface{}) {
	l.logger.Infof(msg, args)
}

// Infoln logs level with new line
func (l *LeveledLogger) Infoln(args ...interface{}) {
	l.logger.Infoln(args)
}

// Debug logs level
func (l *LeveledLogger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

// Debugf logs level with message
func (l *LeveledLogger) Debugf(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args)
}

// Debugln logs level with new line
func (l *LeveledLogger) Debugln(args ...interface{}) {
	l.logger.Debugln(args)
}
