package adaptlog

import "errors"

// Level type
type Level int

// These logging levels are used to setup the logging
const (
	// Panic level
	Panic Level = iota
	// Fatal level
	Fatal
	// Error level
	Error
	// Warn level
	Warn
	// Info level
	Info
	// Debug level
	Debug
)

// ErrorLogger interface. Introduces Error logging facilities.
type ErrorLogger interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})
}

// WarnLogger interface. Introduces Warn logging facilities.
type WarnLogger interface {
	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})
}

// InfoLogger interface. Introduces Info logging facilities.
type InfoLogger interface {
	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})
}

// DebugLogger interface. Introduces Debug logging facilities.
type DebugLogger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})
}

// LvlLogger interface. Introduces Error, Warn, Info and Debug logging facilities.
type LvlLogger interface {
	FatalLogger
	PanicLogger
	ErrorLogger
	WarnLogger
	InfoLogger
	DebugLogger
}

var levelLogger LvlLogger
var defaultLevel Level

// ConfigLevelLogger configures a leveled logger
func ConfigLevelLogger(logger LvlLogger, level Level) {
	levelLogger = logger
	defaultLevel = level
}

// NewDefaultLevelLogger creates a new logger with the default level
func NewDefaultLevelLogger() (*LevelLogger, error) {

	if levelLogger == nil {
		return nil, errors.New("Level logger is not configured!")
	}

	return &LevelLogger{levelLogger, defaultLevel}, nil
}

// NewLevelLogger creates a new level logger with a specified log level
func NewLevelLogger(level Level) (*LevelLogger, error) {

	if levelLogger == nil {
		return nil, errors.New("Level logger is not configured!")
	}

	return &LevelLogger{levelLogger, level}, nil
}

// LevelLogger for logging with level support
type LevelLogger struct {
	logger LvlLogger
	level  Level
}

// Panic level logging
func (l *LevelLogger) Panic(args ...interface{}) {
	if l.level < Panic {
		return
	}
	l.logger.Panic(args)
}

// Panicf level logging with message
func (l *LevelLogger) Panicf(msg string, args ...interface{}) {
	if l.level < Panic {
		return
	}
	l.logger.Panicf(msg, args)
}

// Panicln level logging with new line
func (l *LevelLogger) Panicln(args ...interface{}) {
	if l.level < Panic {
		return
	}
	l.logger.Panicln(args)
}

// Fatal level logging
func (l *LevelLogger) Fatal(args ...interface{}) {
	if l.level < Fatal {
		return
	}
	l.logger.Fatal(args)
}

// Fatalf level logging with message
func (l *LevelLogger) Fatalf(msg string, args ...interface{}) {
	if l.level < Fatal {
		return
	}
	l.logger.Fatalf(msg, args)
}

// Fatalln level logging with new line
func (l *LevelLogger) Fatalln(args ...interface{}) {
	if l.level < Fatal {
		return
	}
	l.logger.Fatalln(args)
}

// Error level logging
func (l *LevelLogger) Error(args ...interface{}) {
	if l.level < Error {
		return
	}
	l.logger.Error(args)
}

// Errorf level logging with message
func (l *LevelLogger) Errorf(msg string, args ...interface{}) {
	if l.level < Error {
		return
	}
	l.logger.Errorf(msg, args)
}

// Errorln level logging with new line
func (l *LevelLogger) Errorln(args ...interface{}) {
	if l.level < Error {
		return
	}
	l.logger.Errorln(args)
}

// Warn level logging
func (l *LevelLogger) Warn(args ...interface{}) {
	if l.level < Warn {
		return
	}
	l.logger.Warn(args)
}

// Warnf level logging with message
func (l *LevelLogger) Warnf(msg string, args ...interface{}) {
	if l.level < Warn {
		return
	}
	l.logger.Warnf(msg, args)
}

// Warnln level logging with new line
func (l *LevelLogger) Warnln(args ...interface{}) {
	if l.level < Warn {
		return
	}
	l.logger.Warnln(args)
}

// Info level logging
func (l *LevelLogger) Info(args ...interface{}) {
	if l.level < Info {
		return
	}
	l.logger.Info(args)
}

// Infof level logging with message
func (l *LevelLogger) Infof(msg string, args ...interface{}) {
	if l.level < Info {
		return
	}
	l.logger.Infof(msg, args)
}

// Infoln level logging with new line
func (l *LevelLogger) Infoln(args ...interface{}) {
	if l.level < Info {
		return
	}
	l.logger.Infoln(args)
}

// Debug level logging
func (l *LevelLogger) Debug(args ...interface{}) {
	if l.level < Debug {
		return
	}
	l.logger.Debug(args)
}

// Debugf level logging with message
func (l *LevelLogger) Debugf(msg string, args ...interface{}) {
	if l.level < Debug {
		return
	}
	l.logger.Debugf(msg, args)
}

// Debugln level logging with new line
func (l *LevelLogger) Debugln(args ...interface{}) {
	if l.level < Debug {
		return
	}
	l.logger.Debugln(args)
}
